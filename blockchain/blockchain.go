package blockchain

import (
	"blockchain/block"
	"blockchain/pow"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

type BlockChain struct {
	lastHash block.Hash
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
		bc.lastHash = block.Hash(data)
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
	b := block.NewBlock(bc.lastHash, txs)
	p:=pow.NewPow(b)
	nonce,hash:=p.Proof()
	if hash==""{
		log.Fatal("block Hashcash proof failed")
	}
	b.SetNonce(nonce).SetHashCurr(hash) //级联调用
	// 将区块加入到链的存储结构中
	if bs, err := block.BlockSerialize(*b); err != nil {
		log.Fatal("block can not be serialized.")
	} else if err = bc.db.Put([]byte("b_" + b.GetHashCurr()), bs, nil); err != nil {
		log.Fatal("block can not be saved")
	}

	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.GetHashCurr()
	// 将最后的区块哈希存储到数据库中
	if err := bc.db.Put([]byte("lastHash"), []byte(b.GetHashCurr()), nil); err != nil {
		log.Fatal("lastHas can not be saved")
	}
	return bc

}


func (bc *BlockChain) GetBlock(hash block.Hash) (*block.Block, error) {
	data, err := bc.db.Get([]byte("b_"+hash), nil)
	if err != nil {
		return nil, err
	}
	//反序列化
	b, err := block.BlockUnSerialize(data)
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
		pow :=pow.NewPow(b)
		if !pow.Validate(){
			log.Fatalf("Block <%s> is not Valid.",hash)
			continue
		}
		fmt.Println("HashCurr", b.GetHashCurr())
		fmt.Println("Txs", b.GetTxs())
		fmt.Println("Time", b.GetTime())
		fmt.Println("HashPrev", b.GetPrevHash())
		fmt.Println("-----------------------------")
		hash = b.GetPrevHash()
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