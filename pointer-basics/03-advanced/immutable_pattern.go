package main

import "fmt"

// Point は2次元座標を表す小さな構造体です。
// データサイズが小さい (int 2個分) ため、コピーコストは非常に低いです。
type Point struct {
	X, Y int
}

// Move (値レシーバ)
// 元の座標は変更せず、「移動後の新しい座標」を返します。
// これが Immutability (不変性) のパターンです。
func (p Point) Move(dx, dy int) Point {
	// ここで p.X += dx としても、書き換わるのはコピーされた p だけです。
	// 新しい Point を作って返します。
	return Point{X: p.X + dx, Y: p.Y + dy}
}

// Add (参考: ポインタレシーバの場合)
// 元の座標を直接書き換えてしまいます。
// 副作用 (Side Effect) があるため、予期せぬバグの原因になることがあります。
func (p *Point) Add(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	start := Point{X: 10, Y: 20}
	fmt.Printf("1. 初期位置: %+v\n", start)

	// --- Immutable な操作 ---
	// start は変更されず、next が新しく作られる
	next := start.Move(5, 5)

	fmt.Printf("2. Move実行後:\n")
	fmt.Printf("   start (元データ): %+v (変わっていません！安全です)\n", start)
	fmt.Printf("   next  (新データ): %+v\n", next)

	// --- Mutable な操作 (ポインタ) ---
	// next そのものが書き換わってしまう
	(&next).Add(100, 100) // ポインタレシーバ

	fmt.Printf("3. Add実行後:\n")
	fmt.Printf("   next が直接書き換わりました: %+v\n", next)

	fmt.Println("\n★ 解説:")
	fmt.Println("  - Move のような「値レシーバ」で新しい値を返す設計は、")
	fmt.Println("    「いつの間にかデータが変わっていた」というバグを防げます。")
	fmt.Println("  - Pointのような小さな構造体なら、コピーしても性能への影響は無視できます。")
}
