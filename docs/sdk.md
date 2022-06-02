
## AkhiraSDK

```go
type AkhiraSDK struct {
    *ProviderHandler
    Storage IpfsStorage
}
```

### func [NewAkhiraSDK](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L23>)

```go
func NewAkhiraSDK(rpcUrlOrChainName string, options *SDKOptions) (*AkhiraSDK, error)
```

#### NewAkhiraSDK

Create a new instance of the Thirdweb SDK

rpcUrlOrName: the name of the chain to connection to \(e\.g\. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func \(\*AkhiraSDK\) [GetContract](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L158>)

```go
func (sdk *AkhiraSDK) GetContract(address string) (*SmartContract, error)
```

#### GetContract

Get an instance of a custom contract deployed with akhira deploy

address: the address of the contract

### func \(\*AkhiraSDK\) [GetContractFromAbi](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L174>)

```go
func (sdk *AkhiraSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error)
```

#### GetContractFromABI

Get an instance of ant custom contract from its ABI

address: the address of the contract

abi: the ABI of the contract

### func \(\*AkhiraSDK\) [GetEdition](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L86>)

```go
func (sdk *AkhiraSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*AkhiraSDK\) [GetEditionDrop](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L130>)

```go
func (sdk *AkhiraSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*AkhiraSDK\) [GetMultiwrap](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L144>)

```go
func (sdk *AkhiraSDK) GetMultiwrap(address string) (*Multiwrap, error)
```

#### GetMultiwrap

Get a Multiwrap contract SDK instance

address: the address of the Multiwrap contract

### func \(\*AkhiraSDK\) [GetNFTCollection](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L72>)

```go
func (sdk *AkhiraSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*AkhiraSDK\) [GetNFTDrop](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L116>)

```go
func (sdk *AkhiraSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*AkhiraSDK\) [GetToken](<https://github.com/akhirachain/go-sdk/blob/main/akhira/sdk.go#L102>)

```go
func (sdk *AkhiraSDK) GetToken(address string) (*Token, error)
```

#### GetToken

Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance
