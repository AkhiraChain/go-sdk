package akhira

import (
  pkg "github.com/akhirachain/go-sdk/pkg/akhira"
)

func NewAkhiraSDK(rpcUrlOrChainName string, options *pkg.SDKOptions) (*pkg.AkhiraSDK, error) {
  return pkg.NewAkhiraSDK(rpcUrlOrChainName, options)
}

