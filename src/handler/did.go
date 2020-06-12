// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package handler

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	createDID  = "createDID"
	deleteDID  = "deleteDID"
	updateHash = "updateHash"
	updateURI  = "updateURI"
	getHash    = "getHash"
	getURI     = "getURI"
)

type DID interface {
	CreateDID(did, didHash, url string) (hash string, err error)
	DeleteDID(did string) (hash string, err error)
	UpdateHash(did, didHash string) (hash string, err error)
	UpdateUri(did, uri string) (hash string, err error)
	GetHash(did string) (hash string, err error)
	GetUri(did string) (uri string, err error)
}

type did struct {
	endpoint string
	account  account.Account
	contract address.Address
	abi      abi.ABI
	gasPrice *big.Int
	gasLimit uint64
}

func NewDID(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64) (d DID, err error) {
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return
	}
	var acc account.Account
	if privateKey != "" {
		acc, err = account.HexStringToAccount(privateKey)
		if err != nil {
			return
		}
	}

	addr, err := address.FromString(contract)
	if err != nil {
		return
	}
	d = &did{endpoint, acc, addr, abi, gasPrice, gasLimit}
	return
}

func (d *did) CreateDID(id, didHash, url string) (hash string, err error) {
	if len(didHash) != 64 {
		err = errors.New("hash should be 32 bytes")
		return
	}
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	hashBytes, err := hex.DecodeString(didHash)
	if err != nil {
		return
	}
	h, err := cli.Contract(d.contract, d.abi).Execute(createDID, id, hashBytes, url).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) DeleteDID(did string) (hash string, err error) {
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	h, err := cli.Contract(d.contract, d.abi).Execute(deleteDID, did).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) UpdateHash(did, didHash string) (hash string, err error) {
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	hashBytes, err := hex.DecodeString(didHash)
	if err != nil {
		return
	}
	h, err := cli.Contract(d.contract, d.abi).Execute(updateHash, did, hashBytes).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) UpdateUri(did, uri string) (hash string, err error) {
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	h, err := cli.Contract(d.contract, d.abi).Execute(updateURI, did, uri).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) GetHash(did string) (hash string, err error) {
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	ret, err := cli.ReadOnlyContract(d.contract, d.abi).Read(getHash, did).Call(context.Background())
	if err != nil {
		return
	}
	hashBytes := [32]uint8{}
	err = ret.Unmarshal(&hashBytes)
	if err != nil {
		return
	}
	hash = hex.EncodeToString(hashBytes[:])
	return
}

func (d *did) GetUri(did string) (uri string, err error) {
	conn, err := newDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	ret, err := cli.ReadOnlyContract(d.contract, d.abi).Read(getURI, did).Call(context.Background())
	if err != nil {
		return
	}
	err = ret.Unmarshal(&uri)
	if err != nil {
		return
	}
	return
}

func newDefaultGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(10 * time.Second)),
		grpc_retry.WithMax(1),
	}
	return grpc.Dial(endpoint,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}
