package main

import (
	"fmt"
	"os"
	"strings"
)

//命令行工具
func main() {
	fmt.Println(os.Args[2])
	arg1 := ""
	if len(os.Args) > 1 {
		arg1 = os.Args[1]
	}
	switch strings.ToLower(arg1) {
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
	}//命令还要根据字母顺序进行排序

