package main

import (
	"blockchain/blockchain"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	//b := blockchain.NewBlock("", "Gensis Block.")
	//fmt.Println(b)
	dbpath := "testdb"
	db,err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//区块链测试
	bc:=blockchain.NewBlockchain(db)
	bc.AddGensisBlock()
	bc.
		AddBlock("first block").
		AddBlock("second block")
	bc.Iterate()
}
