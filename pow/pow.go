package pow

import (
	"blockchain/block"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
)

type ProofofWork struct {
	block  *block.Block
	target *big.Int
}

func NewPow(b *block.Block) *ProofofWork {
	p := &ProofofWork{
		block:  b,
		target: big.NewInt(1),
	}
	p.target.Lsh(p.target, uint(block.Hashlen-b.GetBits()+1))
	return p
}
func (p *ProofofWork) Proof() (int, block.Hash) { //用于HASH cash验证的方法，返回使用的nonce和区块的hash
	var hashInt big.Int
	serviceStr := p.block.GenServiceStr()
	nonce := 0
	fmt.Printf("Target:\t%d\n", p.target)
	//迭代计算，防止溢出
	for nonce <= math.MaxInt64 {
		hash := sha256.Sum256([]byte(serviceStr + strconv.Itoa(nonce)))
		//得到一个大整形的数据
		hashInt.SetBytes(hash[:])
		fmt.Printf("Hash :\t%s\t%d\n", hashInt.String(), nonce)
		if hashInt.Cmp(p.target) == -1 {
			return nonce, block.Hash(fmt.Sprintf("%x", hash))
		}
		nonce++
	}
	return 0, ""
}
func (p *ProofofWork) Validate() bool{
	serviceStr := p.block.GenServiceStr()
	data :=serviceStr+strconv.Itoa(p.block.GetNonce())
	hash :=sha256.Sum256([]byte(data))
	if p.block.GetHashCurr()!=fmt.Sprintf("%x",hash){
		log.Println("not equal")
		return false
	}
	//比较是否满足难题
	target := big.NewInt(1)
	target.Lsh(target, uint(block.Hashlen - p.block.GetBits() +1)) // left shift
	hashInt := new(big.Int)
	hashInt.SetBytes(hash[:])
	// 不小于
	if hashInt.Cmp(target) != -1 {
		log.Println("not less then ")
		return false
	}

	return true
}
