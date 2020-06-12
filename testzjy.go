package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	nonce := 0
	bits :=8//256位前多少位为0
	target :=big.NewInt(1)
	target.Lsh(target,uint(256-bits+1))
	serviceStr := "block data" //服务字符串
	var hashInt big.Int
	for {
		data := serviceStr + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])
		fmt.Println(hashInt.String(),nonce)
		//fmt.Printf("%x\n", hash)
		if hashInt.Cmp(target) ==-1{
			fmt.Printf("本机挖矿成功")
			return
		}
		nonce++
	}
}
