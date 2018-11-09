package main

import (
	"flag"
	"fmt"
)

func main() {
	// パラメータ名は,パラメータ名、デフォルト値、説明文
	//　戻り値はポインタになる
	i := flag.Int("n", 1000, "-n で整数を指定する")
	s := flag.String("s", "status", "-s で文字列を入力")

	// パラメータの解析
	flag.Parse()

	// 取得したパラメータの出力
	fmt.Println(*i, *s)
}
