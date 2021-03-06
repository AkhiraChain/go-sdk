
## Multiwrap

You can access the Multiwrap interface from the SDK as follows:

```
import (
	"github.com/akhirachain/go-sdk/akhira"
)

privateKey = "..."

sdk, err := akhira.NewAkhiraSDK("mumbai", &akhira.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetMultiwrap("{{contract_address}}")
```

```go
type Multiwrap struct {
    *ERC721
}
```

### func \(\*Multiwrap\) [GetWrappedContents](<https://github.com/akhirachain/go-sdk/blob/main/akhira/multiwrap.go#L69>)

```go
func (multiwrap *Multiwrap) GetWrappedContents(wrappedTokenId int) (*MultiwrapBundle, error)
```

Get the contents of a wrapped token bundle\.

wrappedTokenId: the ID of the wrapped token bundle

returns: the contents of the wrapped token bundle

#### Example

```
tokenId := 0
contents, err := contract.GetWrappedContents(tokenId)
erc20Tokens := contents.Erc20Tokens
erc721Tokens := contents.Erc721Tokens
erc1155Tokens := contents.Erc1155Tokens
```

### func \(\*Multiwrap\) [Unwrap](<https://github.com/akhirachain/go-sdk/blob/main/akhira/multiwrap.go#L205>)

```go
func (multiwrap *Multiwrap) Unwrap(wrappedTokenId int, recipientAddress string) (*types.Transaction, error)
```

#### Unwrap a wrapped token bundle into its contents

wrappedTokenId: the ID of the wrapped token bundle

recipientAddress: the optional address to send the wrapped token to

returns: the contents of the wrapped token bundle

#### Example

```
tokenId := 0
tx, err := contract.Unwrap(tokenId, "")
```

### func \(\*Multiwrap\) [Wrap](<https://github.com/akhirachain/go-sdk/blob/main/akhira/multiwrap.go#L156>)

```go
func (multiwrap *Multiwrap) Wrap(contents *MultiwrapBundle, wrappedTokenMetadata interface{}, recipientAddress string) (*types.Transaction, error)
```

Wrap any number of ERC20, ERC721, or ERC1155 tokens into a single wrapped token

contents: the tokens to wrap into a single wrapped token

wrappedTokenMetadata: the NFT Metadata or URI to as the metadata for the wrapped token

recipientAddress: the optional address to send the wrapped token to

returns: the transaction receipt of the wrapping

#### Example

```
contents := &akhira.MultiwrapBundle{
	ERC20Tokens: []*akhira.MultiwrapERC20{
		&akhira.MultiwrapERC20{
			ContractAddress: "0x...",
			Quantity:        1,
		},
	},
	ERC721Tokens: []*akhira.MultiwrapERC721{
		&akhira.MultiwrapERC721{
			ContractAddress: "0x...",
			TokenId:         1,
		},
	},
	ERC1155Tokens: []*akhira.MultiwrapERC1155{
		&akhira.MultiwrapERC1155{
			ContractAddress: "0x...",
			TokenId:         1,
			Quantity:        1,
		},
	},
}

wrappedTokenMetadata := &akhira.NFTMetadataInput{
	Name: "Wrapped Token"
}

// This will mint the wrapped token to the connected wallet
tx, err := contract.Wrap(contents, wrappedTokenMetadata, "")
```

## type [MultiwrapBundle](<https://github.com/akhirachain/go-sdk/blob/main/akhira/types.go#L215-L219>)

```go
type MultiwrapBundle struct {
    ERC20Tokens   []*MultiwrapERC20
    ERC721Tokens  []*MultiwrapERC721
    ERC1155Tokens []*MultiwrapERC1155
}
```

## type [MultiwrapERC1155](<https://github.com/akhirachain/go-sdk/blob/main/akhira/types.go#L209-L213>)

```go
type MultiwrapERC1155 struct {
    ContractAddress string
    TokenId         int
    Quantity        int
}
```

## type [MultiwrapERC20](<https://github.com/akhirachain/go-sdk/blob/main/akhira/types.go#L199-L202>)

```go
type MultiwrapERC20 struct {
    ContractAddress string
    Quantity        float64
}
```

## type [MultiwrapERC721](<https://github.com/akhirachain/go-sdk/blob/main/akhira/types.go#L204-L207>)

```go
type MultiwrapERC721 struct {
    ContractAddress string
    TokenId         int
}
```
