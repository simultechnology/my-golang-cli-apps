package main

import "fmt"

func main() {
	fmt.Println("--- 1. 配列の代入 (Full Copy) ---")
	// 配列 (長さ固定)
	original := [3]string{"A", "B", "C"}

	// 代入すると、配列全体がコピーされます
	copied := original

	fmt.Printf("originalのアドレス: %p\n", &original)
	fmt.Printf("copied  のアドレス: %p (全然違う場所です！)\n", &copied)

	// copiedを変更してみる
	copied[0] = "Z"

	fmt.Printf("original[0]: %s (元のデータは安全保たれます)\n", original[0])
	fmt.Printf("copied[0]  : %s\n", copied[0])

	fmt.Println("\n--- 2. 関数への受け渡し (Copy) ---")
	testArrayArg(original)
	fmt.Printf("関数呼び出し後の original[0]: %s (変わっていません)\n", original[0])
}

// 配列を受け取る関数
func testArrayArg(arr [3]string) {
	fmt.Printf("  [関数内] 受け取った配列のアドレス: %p (これもコピーです)\n", &arr)
	arr[0] = "Modified"
}
