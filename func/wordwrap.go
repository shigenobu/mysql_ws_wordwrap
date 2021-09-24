/*
 wordwrap
*/
package _func

import (
	"strings"
	"unicode/utf8"
)

// 指定位置で区切り文字を入れる
func Wordwrap(input string, limit int, split string) string {
	// 確認
	if input == "" {
		return ""
	}

	// デフォルト値
	if limit <= 0 || limit > 8192 {
		limit = 100
	}
	if split == "" {
		split = "\n"
	}

	// 分割文字の文字数
	sLen := utf8.RuneCountInString(split)

	// 初期化
	var builder strings.Builder

	// 代入
	newInput := input
	for {
		// 残っている文字数が分割文字数以下である場合
		if utf8.RuneCountInString(newInput) <= limit {
			// 残っている文字をそのままつけて終了
			builder.WriteString(newInput)
			break
		}

		// 文字列に変更する
		newRune := []rune(newInput)

		// 最初に出現する文字列のインデックスを取得する
		s := MbIndex(newInput, split)

		//fmt.Printf("newInput - %d - %s\n", s, newInput)

		// 先頭に出現した場合、区切り文字列までを追加して、区切り文字列以降を新しく代入する
		if s == 0 {
			builder.WriteString(string(newRune[:sLen]))
			newInput = string(newRune[sLen:])
			continue
		}

		// 1文字目から分割文字数の間の場合、区切り文字列の手前まで追加して、それ以降を新しく代入する
		if s > 0 && s <= limit {
			builder.WriteString(string(newRune[:s]))
			newInput = string(newRune[s:])
			continue
		}

		// 区切り文字列が分割文字数までに見つからなかった場合、分割文字数まで追加して、分割文字以降を新しく代入し、分割文字を追加する
		builder.WriteString(string(newRune[:limit]))
		newInput = string(newRune[limit:])
		builder.WriteString(split)
	}

	return builder.String()


	//// replace
	//replaced := strings.ReplaceAll(input, split, "")
	//
	//// execute
	//len := len(replaced)
	//var sb strings.Builder
	//idx := 1
	//for pos, rv := range replaced {
	//	sb.WriteRune(rv)
	//	rl := utf8.RuneLen(rv)
	//	if pos + rl >= len {
	//		break
	//	}
	//
	//	if idx % limit == 0 {
	//		sb.WriteString(split)
	//	}
	//	idx++
	//}
	//return sb.String(
	//
}

// 文字列が見つかった位置を返す（見つからなければ-1）
func MbIndex(input string, search string) int {
	s := strings.Index(input, search)
	//fmt.Printf("s:%d\n", s)

	idx := 0
	for pos, _ := range input {
		//fmt.Printf("pos:%d\n", pos)
		if pos == s {
			return idx
		}
		idx++
	}
	return -1
}