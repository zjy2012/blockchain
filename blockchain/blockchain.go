package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	. "time"
)

type BlockChian struct {
	lastHash Hash
	db *leveldb.DB
}

func NewBlockchain(db *leveldb.DB) *BlockChian {
	bc := &BlockChian{
		db: db,
	}
	return bc
}
func (bc *BlockChian) AddGensisBlock() *BlockChian {
	if bc.lastHash!=""{
return  bc
	}

	return bc.AddBlock("The Gensis Block")
}
func (bc *BlockChian) AddBlock(txs string) *BlockChian {
b:=NewBlock(bc.lastHash,txs)
if bs,err:=BlockSerialize(*b);err!=nil{
	bc.db.Put([]byte(b.hashCurr),bs,nil)
}

bc.lastHash=b.hashCurr

	return bc
}
func (bc*BlockChian)Iterate(){
	for hash :=bc.lastHash;hash!="";{
		b:=bc.blocks[hash]
		fmt.Println("HashCurr",b.hashCurr)
		fmt.Println("Tsx",b.txs)
		fmt.Println("Time",b.header.time.Format(UnixDate))
		fmt.Println("HashPrev",b.header.hashPrevBlock)
		fmt.Println("-----------------------------")
		hash=b.header.hashPrevBlock
	}
}