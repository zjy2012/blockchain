package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type  Block struct {
	CurrHash string
	Txs string
}
func main(){
	b:=&Block{
		"1456151215622",
		"lalalal",
	}
	var bb bytes.Buffer
	enc :=gob.NewEncoder(&bb)//构造编码器
	enc.Encode(b)//编码到bb里
	//fmt.Println(bb.Bytes(),bb.String())
	result :=bb.Bytes()
	//解码
	var bbr bytes.Buffer
	//将之前编码的数据，放入到缓冲中
	bbr.Write(result)
	dec :=gob.NewDecoder(&bb)
	b1:=Block{}//解码时，需要提供解码的数据类型
	dec.Decode(&b1)
	fmt.Println(b1)
}