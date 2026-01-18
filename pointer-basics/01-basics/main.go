package main

import "fmt"

func main() {
	fmt.Println("--- 1. 変数とアドレスの関係 ---")
	// 通常の変数を宣言
	count := 10
	fmt.Printf("変数 count の値: %d\n", count)
	fmt.Printf("変数 count のアドレス (住所): %v\n", &count)

	fmt.Println("\n--- 2. ポインタ変数の作成 ---")
	// countのアドレスを格納する変数が「ポインタ変数」
	// int型の変数のアドレスなので、型は *int になります
	var ptr *int = &count

	fmt.Printf("ポインタ ptr の値 (つまりcountの住所): %v\n", ptr)
	fmt.Printf("ポインタ ptr が指している中身 (*ptr): %d\n", *ptr)

	fmt.Println("\n--- 3. ポインタ経由での書き換え ---")
	// ポインタを使って、その住所にあるデータを書き換える
	*ptr = 99
	fmt.Printf("*ptr = 99 を実行しました。\n")
	fmt.Printf("変数 count の新しい値: %d (魔法のように変わりました！)\n", count)
}
