package pow

import (
	"blockchain/blockchain"
	"math"
	"math/big"
)

type ProofofWork struct {
	block  *blockchain.Block
	target *big.Int
}

func NewPow(b *blockchain.Block) *ProofofWork {
	p := &ProofofWork{
		block: b,
	}
	p.target.Lsh(p.target, uint(blockchain.Hashlen-b.GetBits()+1))
	return p
}
func (p *ProofofWork) Proof() (int, string) {
	var hashInt big.Int
	serviceStr := p.block.GenServiceStr()
	nonce := 0
	//贴袋计算，防止移除
	for nonce <= math.MaxInt64 {

	}
	return 0, ""
}
