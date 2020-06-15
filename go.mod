module blockchain

go 1.14

require (
	github.com/mr-tron/base58 v1.2.0
	github.com/syndtr/goleveldb v1.0.0
)

replace ./golang.org/x/sys => ./github.com/golang/sys

replace ./golang.org/x/crypto => ./github.com/golang/crypto
