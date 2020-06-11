package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	dbpath := "testdb"
	db,err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	key :="zjy"
	//if err:=db.Put([]byte(key),[]byte("Blockchain Demo"),nil);err!=nil{
	//	log.Fatal(err)
	//}
	//log.Println("put success")  对这里进行了持久化的存贮，所以后面读取的时候就不会报错
	data,err :=db.Get([]byte(key),nil)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(data,string(data))

}
