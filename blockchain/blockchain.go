package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	. "time"
)

type BlockChain struct {
	lastHash Hash
	db       *leveldb.DB
}

func NewBlockchain(db *leveldb.DB) *BlockChain {
	bc := &BlockChain{
		db: db,
	}
	// 初始化 lastHash
	// 读取最后的区块哈希
	data, err := bc.db.Get([]byte("lastHash"), nil)
	if err == nil { // 读取到 lasthash
		bc.lastHash = Hash(data)
	}
	return bc
}

func (bc *BlockChain) AddGensisBlock() *BlockChain {
	if bc.lastHash != "" {
		return bc
	}

	return bc.AddBlock("The Gensis Block")
}
func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	// 构建区块
	b := NewBlock(bc.lastHash, txs)
	// 将区块加入到链的存储结构中
	if bs, err := BlockSerialize(*b); err != nil {
		log.Fatal("block can not be serialized.")
	} else if err = bc.db.Put([]byte("b_" + b.hashCurr), bs, nil); err != nil {
		log.Fatal("block can not be saved")
	}

	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.hashCurr
	// 将最后的区块哈希存储到数据库中
	if err := bc.db.Put([]byte("lastHash"), []byte(b.hashCurr), nil); err != nil {
		log.Fatal("lastHas can not be saved")
	}
	return bc

}


func (bc *BlockChain) GetBlock(hash Hash) (*Block, error) {
	data, err := bc.db.Get([]byte("b_"+hash), nil)
	if err != nil {
		return nil, err
	}
	//反序列化
	b, err := BlockUnSerialize(data)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
func (bc *BlockChain) Iterate() {
	for hash := bc.lastHash; hash != ""; {
		b, err := bc.GetBlock(hash)
		if err != nil {
			log.Fatalf("Block<%s> is not exists", hash)
			return
		}
		fmt.Println("HashCurr", b.hashCurr)
		fmt.Println("Tsx", b.txs)
		fmt.Println("Time", b.header.time.Format(UnixDate))
		fmt.Println("HashPrev", b.header.hashPrevBlock)
		fmt.Println("-----------------------------")
		hash = b.header.hashPrevBlock
	}
}
func (bc*BlockChain)Clear(){
   bc.db.Delete([]byte("lastHash"),nil)
	iter := bc.db.NewIterator(util.BytesPrefix([]byte("b_")), nil)
	for iter.Next() {
		bc.db.Delete(iter.Key(), nil)
	}
	iter.Release()
	bc.lastHash=""
}