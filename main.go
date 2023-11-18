package main

import (
	"fmt"
	"flag"
	//"os"
)

func main(){
	// var str = flag.String("オプション名", "初期値", "説明")
	f := flag.String("flag1", "hoge", "flag 1")
	flag.Parse()

	fmt.Printf("Hello %s", *f)
}