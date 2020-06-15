package main

import (
	"blockchain/blockchain"
	"blockchain/wallet"
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
)

//命令行工具
	func main() {
		//fmt.Println(os.Args[2])
		//数据库连接
		dbpath := "testdb"
		db, err := leveldb.OpenFile(dbpath, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

	//初始化区块链

	bc := blockchain.NewBlockchain(db)
	bc.AddGensisBlock()

	arg1 := ""
	if len(os.Args) > 1 {
		arg1 = os.Args[1]
	}

	switch strings.ToLower(arg1) { //这里的命令是通过switch进行展示的
		case "createblock":
			fs := flag.NewFlagSet("createblock", flag.ExitOnError)
			txs := fs.String("txs", "xxxxx", "")
			fs.Parse(os.Args[2:]) //命令行工具，通过os.args[]来通过下标来获取
			bc.AddBlock(*txs)
			//if !fs.Parsed(){
			//	log.Fatal("createblock args parsed error.")
			//}
			//fmt.Println(txs,*txs)

		case "show":
			bc.Iterate()
		//初始化操作，删除已有区块，重新增加一个创世区块
		case "init":
			bc.Clear()
			bc.AddGensisBlock()
		case "create wallet":
		w:=wallet.NewWallet()
		fmt.Printf("your address %s \n",w.Address)
		case "help":
			fallthrough

		default:
			Usage()
		}
	}

	func Usage() {
		fmt.Println("bcli is a tool for Blockchain.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("\t%s\t\t%s\n", "bcli createblock <txs>", "create block on blockchain")
		fmt.Printf("\t%s\t\t\t%s\n", "bcli init", "initial blockchain")
		fmt.Printf("\t%s\t\t\t%s\n", "bcli help", "help info for bcli")
		fmt.Printf("\t%s\t\t\t%s\n", "bcli show", "show blocks in chain.")
	} //命令还要根据字母顺序进行排序
