package block

import (
	//"crypto/sha256"
	"fmt"
	time2 "time"
)

const nodeVersion = 0
const Hashlen =256
const blockBits = 8

type Hash = string
//区块的构造方法
func NewBlock(prevHash Hash, tsx string) *Block { //当前的交易信息
	b := &Block{
		header: BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: prevHash, // 设置前面的区块哈希
			time:          time2.Now(),
			bits:          blockBits,
		},
		txs:       tsx, // 设置数据
		txCounter: 1,   // 计数交易
	} // 计算设置当前区块的哈希
	return b
}

//区块主体
type Block struct {
	header    BlockHeader //结构体的嵌套
	txs       string      //交易列表
	txCounter int         //交易计数器
	hashCurr  Hash        //当前区块链哈希
}

//去块头
type BlockHeader struct {
	version        int
	hashPrevBlock  Hash
	hashMerkleRoot Hash
	time           time2.Time
	bits           int
	nonce          int
} //这里会涉及到一个结构体的封装性 ，外面的包不能随便的进行调用，所以会设置set 和个体
//方法
//func (bh *BlockHeader) stringify() string {
//	return fmt.Sprintf("%d%s%s%d%d%d",
//		bh.version,
//		bh.hashPrevBlock,
//		bh.hashMerkleRoot,
//		bh.time.UnixNano(), // 得到时间戳，nano 级别
//		bh.bits,
//		bh.nonce,
//	)
//}
//
////设置当前区块hash
//func (b *Block) setHashCurr() *Block {
//	headerStr := b.header.stringify()
//	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte (headerStr)))
//	return b
//}
func (b*Block)GetBits()int{
	return b.header.bits
}
func (b*Block) GetHashCurr() Hash {
	return b.hashCurr
}

func (b*Block) GetTxs() string {
	return b.txs
}

func (b*Block) GetTime() time2.Time {
	return b.header.time
}

func (b*Block) GetPrevHash() Hash {
	return b.header.hashPrevBlock
}

func (b*Block) GetNonce() int {
	return b.header.nonce
}

//生成用于Pow的服务字符串
func (b *Block) GenServiceStr() string {
	return fmt.Sprintf("%d%s%s%s%d",
		b.header.version,
		b.header.hashPrevBlock,
		b.header.hashMerkleRoot,
		b.header.time.Format("2006-01-02 15:04:05.999999999 -0700 MST"),
		b.header.bits,
	)
}
func (b *Block) SetHashCurr(hash Hash) *Block {

	b.hashCurr = hash
	return b
}
func (b*Block)SetNonce(nonce int )*Block {
	b.header.nonce=nonce
	return  b
}