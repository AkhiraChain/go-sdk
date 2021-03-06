package akhira

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AkhiraSDK struct {
	*ProviderHandler
	Storage IpfsStorage
}

// NewAkhiraSDK
//
// Create a new instance of the Thirdweb SDK
//
// rpcUrlOrName: the name of the chain to connection to (e.g. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche") or the RPC URL to connect to
//
// options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL
func NewAkhiraSDK(rpcUrlOrChainName string, options *SDKOptions) (*AkhiraSDK, error) {
	rpc, err := getDefaultRpcUrl(rpcUrlOrChainName)
	if err != nil {
		return nil, err
	}

	provider, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	if options == nil {
		handler, err := NewProviderHandler(provider, "")
		if err != nil {
			return nil, err
		}

		storage := newIpfsStorage(defaultIpfsGatewayUrl)
		sdk := &AkhiraSDK{
			ProviderHandler: handler,
			Storage:         *storage,
		}
		return sdk, nil
	} else {
		gatewayUrl := options.GatewayUrl
		if gatewayUrl == "" {
			gatewayUrl = defaultIpfsGatewayUrl
		}

		handler, err := NewProviderHandler(provider, options.PrivateKey)
		if err != nil {
			return nil, err
		}

		storage := newIpfsStorage(gatewayUrl)

		sdk := &AkhiraSDK{
			ProviderHandler: handler,
			Storage:         *storage,
		}
		return sdk, nil
	}
}

// GetNFTCollection
//
// Get an NFT Collection contract SDK instance
//
// address: the address of the NFT Collection contract
func (sdk *AkhiraSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	return newNFTCollection(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEdition
//
// Get an Edition contract SDK instance
//
// address: the address of the Edition contract
func (sdk *AkhiraSDK) GetEdition(address string) (*Edition, error) {
	return newEdition(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetToken
//
// Returns a Token contract SDK instance
//
// address: address of the token contract
//
// Returns a Token contract SDK instance
func (sdk *AkhiraSDK) GetToken(address string) (*Token, error) {
	return newToken(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetNFTDrop
//
// Get an NFT Drop contract SDK instance
//
// address: the address of the NFT Drop contract
func (sdk *AkhiraSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	return newNFTDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEditionDrop
//
// Get an Edition Drop contract SDK instance
//
// address: the address of the Edition Drop contract
func (sdk *AkhiraSDK) GetEditionDrop(address string) (*EditionDrop, error) {
	return newEditionDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMultiwrap
//
// Get a Multiwrap contract SDK instance
//
// address: the address of the Multiwrap contract
func (sdk *AkhiraSDK) GetMultiwrap(address string) (*Multiwrap, error) {
	return newMultiwrap(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetContract
//
// Get an instance of a custom contract deployed with akhira deploy
//
// address: the address of the contract
func (sdk *AkhiraSDK) GetContract(address string) (*SmartContract, error) {
	abi, err := fetchContractMetadataFromAddress(address, sdk.GetProvider(), &sdk.Storage)
	if err != nil {
		return nil, err
	}

	return sdk.GetContractFromAbi(address, abi)
}

// GetContractFromABI
//
// Get an instance of ant custom contract from its ABI
//
// address: the address of the contract
//
// abi: the ABI of the contract
func (sdk *AkhiraSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error) {
	return newSmartContract(
		sdk.GetProvider(),
		common.HexToAddress(address),
		abi,
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

func getDefaultRpcUrl(rpcUrlorName string) (string, error) {
	switch rpcUrlorName {
	case "mumbai":
		return "https://polygon-mumbai.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "rinkeby":
		return "https://eth-rinkeby.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "goerli":
		return "https://eth-goerli.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "polygon":
		return "https://polygon-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "mainnet", "ethereum":
		return "https://eth-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "fantom":
		return "https://rpc.ftm.tools", nil
	case "avalanche":
		return "https://rpc.ankr.com/avalanche", nil
	case "celo":
		return "https://forno.celo.org", nil
	default:
		if strings.HasPrefix(rpcUrlorName, "http") {
			return rpcUrlorName, nil
		} else {
			return "", fmt.Errorf("invalid rpc url or chain name: %s", rpcUrlorName)
		}
	}
}
