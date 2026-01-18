package main

import "fmt"

// スライスの「ポインタ」を受け取る関数
// *[]int は「スライス(3点セット)がある場所の住所」です
func appendPtr(s *[]int) {
	// *s は呼び出し元のスライスそのものです
	// これを上書きすると、呼び出し元も変わります
	*s = append(*s, 999)
	fmt.Printf("  [関数内] appendしました。長さ: %d, 容量: %d\n", len(*s), cap(*s))
}

func main() {
	slice := []int{10, 20}
	fmt.Printf("呼び出し前: %v\n", slice)

	// スライスの「住所」を渡す (&slice)
	appendPtr(&slice)

	fmt.Printf("呼び出し後: %v (変わりました！)\n", slice)

	fmt.Println("\n★ 解説:")
	fmt.Println("  スライスのポインタ (*[]int) を渡せば、戻り値なしで append を反映できます。")
	fmt.Println("  ただし、毎回 `*s` と書くのが少し面倒なのと、")
	fmt.Println("  Go言語では `s = append(s, val)` と書くのが慣習(Idiom)になっているため、")
	fmt.Println("  あまり頻繁には使われないテクニックです。")
}
