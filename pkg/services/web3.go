package services

import (
	"errors"
	httperrors "web3-with-go/pkg/http/errors"

	"github.com/chenzhijie/go-web3"
)

func getNode(protocol string) (*web3.Web3, error) {
	rpcs := map[string]string{
		"ETHEREUM": "",
		"BSC":      "",
	}
	rpcProviderUrl := rpcs[protocol]
	if rpcProviderUrl == "" {
		return nil, errors.New("Protocol not supported")
	}
	web3, err := web3.NewWeb3(rpcProviderUrl)
	if err != nil {
		return nil, err
	}
	return web3, nil
}

func GetCurrentBlock(protocol string) (uint64, *httperrors.HttpError) {
	node, err := getNode(protocol)
	if err != nil {
		return 0, httperrors.NewBadRequestHttpError(err.Error())
	}
	blockNumber, err2 := node.Eth.GetBlockNumber()
	if err2 != nil {
		return 0, httperrors.NewInternalErrorHttpError(err2.Error())
	}

	if err != nil {
		err := httperrors.NewInternalErrorHttpError(err.Error())
		return 0, err
	}
	return blockNumber, nil
}
