.PHONY: abi test docs publish

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out internal/abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out internal/abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out internal/abi/token_erc1155.go --type TokenERC1155
	# If you want to generate drop contracts, you'll have to delete a struct
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out internal/abi/drop_erc721.go --type DropERC721
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC1155.json --out internal/abi/drop_erc1155.go --type DropERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Multiwrap.json --out internal/abi/multiwrap.go --type Multiwrap


	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out internal/abi/ierc20.go --type IERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC721.json --out internal/abi/ierc721.go --type IERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC1155.json --out internal/abi/ierc1155.go --type IERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC165.json --out internal/abi/ierc165.go --type IERC165

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md --repository.default-branch main ./akhira
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/finish.md
	node ./scripts/generate-snippets.mjs

cmd: FORCE
	cd cmd/akhira && go build -o ../../bin/akhira && cd -

test-nft-read:
	./bin/akhira nft getAll -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/akhira nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira nft mintLink -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-sigmint:
	./bin/akhira nft sigmint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/akhira edition getAll -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/akhira edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-sigmint:
	./bin/akhira edition sigmint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira edition sigmint-tokenid -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-read:
	./bin/akhira token get -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-write:
	./bin/akhira token mint -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira token mintBatch -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-read:
	./bin/akhira multiwrap getAll -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	# ./bin/akhira multiwrap getContents -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-write:
	./bin/akhira multiwrap wrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC} -n ${GO_NFT_COLLECTION} -e ${GO_EDITION} -t ${GO_TOKEN}
	./bin/akhira multiwrap unwrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/akhira nftdrop getAll -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-write:
	./bin/akhira nftdrop createBatch -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira nftdrop claim -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-read:
	./bin/akhira editiondrop getAll -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-write:
	./bin/akhira editiondrop createBatch -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/akhira editiondrop claim -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-storage:
	./bin/akhira storage upload
	./bin/akhira storage uploadBatch
	./bin/akhira storage uploadImage
	./bin/akhira storage uploadImageLink

test-custom:
	./bin/akhira custom set -a ${GO_CUSTOM} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test:
	make cmd
	make test-nft-read
	make test-nft-write
	make test-edition-read
	make test-edition-write
	make test-token-read
	make test-token-write
	make test-drop-read
	make test-drop-write
	make test-edition-drop-read
	make test-edition-drop-write
	make test-multiwrap-read
	make test-multiwrap-write
	make test-storage

publish:
	# Make sure to pass the TAG variable to this command ex: `make publish TAG=v2.0.0`
	# Publish following https://go.dev/doc/modules/publishing
	go mod tidy
	git tag $(TAG)
	git push origin $(TAG)
	GOPROXY=proxy.golang.org go list -m github.com/akhira-dev/go-sdk@$(TAG)

FORCE: ;