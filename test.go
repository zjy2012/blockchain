package main

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	//"github.com/tyler-smith/go-bip32"
	"fmt"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(mnemonic)
	//生成密钥对
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
//构建种子，生成主密钥
	masterKey, _ := bip32.NewMasterKey(seed)
	//然后再生成公钥
	publicKey := masterKey.PublicKey()
	fmt.Println("Private key:",masterKey.String())
	fmt.Println("Public key:",publicKey.String())
	userMnemonic :=mnemonic
	//生成熵
	userMnemonic, _ = bip39.EntropyFromMnemonic(userMnemonic)
	userSeed :=bip39.NewSeed(userMnemonic,"Secret Passphrase")
}
