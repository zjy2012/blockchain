package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"github.com/mr-tron/base58"
	//"github.com/btcsuite/btcd "
)

type Address = string

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  ecdsa.PublicKey
	Address    Address
}

func NewWallet() *Wallet {
	w := &Wallet{}
	w.Genkey()
	w.GenAddress()
	return w
}

func (w *Wallet) Genkey() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	w.privateKey = privateKey
	return w
}
func (w *Wallet) GenAddress() *Wallet {
	//利用私钥形成公钥
	pubKey := w.genPubKey()
	// pubHash: ripemd160(sha256(pubkey))
	shaHash := sha256.Sum256(pubKey)
	rpmd := ripemd160.New()
	rpmd.Write(shaHash[:])
	pubHash := rpmd.Sum(nil)
	// 计算checkSum 校验值
	h1 := sha256.Sum256(pubKey)
	checkSum := sha256.Sum256(h1[:])
	data := append(append([]byte{0}, pubHash...), checkSum[:2]...)
	w.Address = base58.Encode(data)
	return w
}
func (w *Wallet) genPubKey() []byte {
	pubkey := append(
		w.privateKey.PublicKey.X.Bytes(),
		w.privateKey.PublicKey.Y.Bytes()...)
	return pubkey
}
