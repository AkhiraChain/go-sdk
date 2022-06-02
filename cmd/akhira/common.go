package main

import (
	"log"

	"github.com/akhirachain/go-sdk/akhira"
)

var (
	akhiraSDK *akhira.AkhiraSDK
)

func initSdk() {
	if sdk, err := akhira.NewAkhiraSDK(
		chainRpcUrl,
		&akhira.SDKOptions{
			PrivateKey: privateKey,
		},
	); err != nil {
		panic(err)
	} else {
		akhiraSDK = sdk
	}
}

func getNftCollection() (*akhira.NFTCollection, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	if contract, err := akhiraSDK.GetNFTCollection(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT Collection object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEdition() (*akhira.Edition, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition on chain %v, contract %v\n", chainRpcUrl, editionAddress)

	if contract, err := akhiraSDK.GetEdition(editionAddress); err != nil {
		log.Println("Failed to create an Edition object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getToken() (*akhira.Token, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Token on chain %v, contract %v\n", chainRpcUrl, tokenAddress)

	if contract, err := akhiraSDK.GetToken(tokenAddress); err != nil {
		log.Println("Failed to create an Token object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getNftDrop() (*akhira.NFTDrop, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Drop on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := akhiraSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEditionDrop() (*akhira.EditionDrop, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition Drop on chain %v, contract %v\n", chainRpcUrl, editionDropContractAddress)

	if contract, err := akhiraSDK.GetEditionDrop(editionDropContractAddress); err != nil {
		log.Println("Failed to create an Edition Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMultiwrap() (*akhira.Multiwrap, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Multiwrap on chain %v, contract %v\n", chainRpcUrl, multiwrapContractAddress)

	if contract, err := akhiraSDK.GetMultiwrap(multiwrapContractAddress); err != nil {
		log.Println("Failed to create a Multiwrap object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getCustom() (*akhira.SmartContract, error) {
	if akhiraSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Custom on chain %v, contract %v\n", chainRpcUrl, customContractAddress)

	if contract, err := akhiraSDK.GetContract(customContractAddress); err != nil {
		log.Println("Failed to create an Custom object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() akhira.IpfsStorage {
	if akhiraSDK == nil {
		initSdk()
	}

	return akhiraSDK.Storage
}
