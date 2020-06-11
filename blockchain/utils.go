package blockchain

import (
	"bytes"
	"encoding/gob"
	time2 "time"
)

type BlockData struct {
	Version        int
	HashPrevBlock  Hash
	HashMerkleRoot Hash
	Time           time2.Time
	Bits           int
	Nonce          int
	Txs            string //交易列表
	TxCounter      int    //交易计数器
	HashCurr       Hash   //当前区块链哈希
}

func BlockSerialize(b Block) ([]byte, error) {
	bd := BlockData{
		Version:        b.header.version,
		HashPrevBlock:  b.header.hashPrevBlock,
		HashMerkleRoot: b.header.hashMerkleRoot,
		Time:           b.header.time,
		Bits:           b.header.bits,
		Nonce:          b.header.nonce,
		Txs:            b.txs,
		TxCounter:      b.txCounter,
		HashCurr:       b.hashCurr,
	}
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(bd); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
func BlockUnSerialize(data []byte) (Block, error) {
	buffer := bytes.Buffer{}
	buffer.Write(data)
	// 解码器
	dec := gob.NewDecoder(&buffer)
	// 解码，反序列化
	bd := BlockData{}
	if err := dec.Decode(&bd); err != nil {
		return Block{}, err
	}
	// 反序列化成功
	return Block{
		header: BlockHeader{
			version:        bd.Version,
			hashPrevBlock:  bd.HashPrevBlock,
			hashMerkleRoot: bd.HashMerkleRoot,
			time:           bd.Time,
			bits:           bd.Bits,
			nonce:          bd.Nonce,
		},
		txs:       bd.Txs,
		txCounter: bd.TxCounter,
		hashCurr:  bd.HashCurr,
	}, nil
}
