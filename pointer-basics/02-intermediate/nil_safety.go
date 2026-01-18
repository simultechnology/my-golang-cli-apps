package main

import "fmt"

type User struct {
	Name string
}

// Goの面白い特徴: ポインタレシーバがnilでも呼び出せる
func (u *User) SayHello() {
	if u == nil {
		fmt.Println("  [メソッド内] 私はnilです。(Ghost)")
		return
	}
	fmt.Printf("  [メソッド内] Hello, I am %s\n", u.Name)
}

func main() {
	fmt.Println("--- 1. nilポインタの危険性 ---")
	var ptr *int
	fmt.Printf("初期化していないポインタの値: %v\n", ptr)

	// ▼ これを実行するとクラッシュします (Panic)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x10055ed04]
	//
	//fmt.Printf("中身を見ようとすると...: %d\n", *ptr)
	fmt.Println("注意: nilポインタをデリファレンス (*ptr) すると panic (クラッシュ) します！")

	fmt.Println("\n--- 2. nilレシーバの許容 ---")
	// 構造体のポインタがnilの場合
	var nobody *User
	fmt.Printf("nobodyの値: %v\n", nobody)

	// 多くの言語ではここでクラッシュしますが、Goは呼び出し可能です
	// (メソッド内で nil チェックをしている場合)
	nobody.SayHello()

	fmt.Println("\n--- 3. ちゃんと実体がある場合 ---")
	somebody := &User{Name: "Gopher"}
	somebody.SayHello()
}
