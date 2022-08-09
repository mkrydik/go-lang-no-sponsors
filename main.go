// Go Lang のスポンサーの提供でお送りします
package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ターミナルウィンドウの幅・高さ情報
type TerminalSize struct {
	// 幅
	width int
	// 高さ
	height int
}

// ターミナルウィンドウの幅・高さを取得する
func getTerminalSize() *TerminalSize {
	var terminalSize TerminalSize
	var err error
	terminalSize.width, terminalSize.height, err = terminal.GetSize(syscall.Stdin)

	if err != nil {
		fmt.Printf("Error has occurred : %+v", err)
		os.Exit(1)
	}

	return &terminalSize
}

// Main
func main() {
	// 定数
	const title string = "提   供"
	const sponsor string = "mkrydik"

	// ターミナルの幅・高さを取得する
	terminalSize := getTerminalSize()

	// ターミナルの幅を基に、出力桁数を決める
	const minWidth = len(sponsor) // 出力する幅の最小値 : `len()` は厳密にはバイト数
	var width = map[bool]int{true: terminalSize.width, false: minWidth}[terminalSize.width >= minWidth]
	var halfWidth = width / 2 // 整数に切り捨てられるので奇数の場合は `width` の半分より 1 少ない値になる
	// テキストを中央揃えにするための余白を求める
	var leftPadding = halfWidth - (minWidth / 2) - 1
	var rightPadding = leftPadding
	if width%2 == 0 && minWidth%2 != 0 {
		leftPadding -= 1
	}
	if width%2 != 0 && minWidth%2 == 0 {
		rightPadding += 1
	}
	// 負数は避ける (レイアウトは崩れる)
	if leftPadding < 0 {
		leftPadding = 0
	}
	if rightPadding < 0 {
		rightPadding = 0
	}

	// ターミナルの高さを基に、出力行数を決める
	const minHeight = 5     // 最小でも5行出力する
	const bottomSpacer = 10 // 下に開ける隙間の行数
	var height = map[bool]int{true: terminalSize.height - bottomSpacer, false: minHeight}[terminalSize.height-bottomSpacer >= minHeight]
	// 行間の数を求める
	var topBottomLines = height / 3
	var centerLines = topBottomLines
	if height%3 == 1 {
		centerLines += 1
	}
	if height%3 == 2 {
		topBottomLines += 1
	}

	// 十分な表示領域がある場合はターミナルをクリアする
	if height > minHeight {
		fmt.Print("\033[H\033[2J")
	}

	// 出力開始
	var topBottomLine = "+" + strings.Repeat("-", width-2) + "+"
	var emptyLine = "|" + strings.Repeat(" ", width-2) + "|"

	// Top
	fmt.Println(topBottomLine)
	// Top Empty Lines
	for i := 0; i < topBottomLines; i++ {
		fmt.Println(emptyLine)
	}
	// Title
	fmt.Println("|" + strings.Repeat(" ", leftPadding) + title + strings.Repeat(" ", rightPadding) + "|")
	// Center Empty Lines
	for i := 0; i < centerLines; i++ {
		fmt.Println(emptyLine)
	}
	// Sponsor
	fmt.Println("|" + strings.Repeat(" ", leftPadding) + sponsor + strings.Repeat(" ", rightPadding) + "|")
	// Bottom Empty Lines
	for i := 0; i < topBottomLines; i++ {
		fmt.Println(emptyLine)
	}
	// Bottom
	fmt.Println(topBottomLine)
}
