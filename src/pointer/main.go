package main

import "fmt"

func main() {
	var num int = 10

	// &を使うと変数のアドレスを取得できる
	// ex.numのアドレス: 0xc0000a4008
	fmt.Println("numのアドレス:", &num)
	// ex.numの値: 10
	fmt.Println("numの値:", num)

	// ポインタ型の変数を宣言 numのアドレスを格納
	// & = アドレス演算子 * = ポインタ演算子
	// ポインタはアドレスを格納するための変数
	// ポインタ型の変数を宣言するには、型の前に*をつける
	var ptr *int = &num

	// ex.ptrが指すアドレス: 0xc0000a4008
	fmt.Println("ptrが指すアドレス:", ptr)
	// 型定義の*とは別で*をつけるとポインタが指すアドレスの値を取得できる デリファレンス
	// ex.ptrが指す値: 10
	fmt.Println("ptrが指す値:", *ptr)
	// これはnumと同じ
	fmt.Println("ptrが指す値:", *&num)
	// ポインタ型自体のアドレスを取得することもできる ポインタも変数でメモリに保存されているのでアドレスを持つ
	// ex.ptrのアドレス: 0xc000006028
	fmt.Println("ptr自体のアドレス", &ptr)

	// ポインタはnumのアドレスを格納しているので、numの値を変更することができる
	*ptr = 20

	// ex.numの新しい値: 20
	fmt.Println("numの新しい値:", num)
}
