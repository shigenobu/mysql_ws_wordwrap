package test

import (
	"crypto/md5"
	"fmt"
	"github.com/shigenobu/mysql_ws_wordwrap/func"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestMbIndex1(t *testing.T) {
	input := "𠮷野家で𠮷野がご飯をたべる"

	pos := _func.MbIndex(input, "ご飯")
	fmt.Println(pos)
	assert.Equal(t, 7, pos)
}

func TestMbIndex2(t *testing.T) {
	input := "𠮷野家で𠮷野がご飯をたべる𠮷野家で𠮷野がご飯をたべる"

	pos := _func.MbIndex(input, "ご飯")
	fmt.Println(pos)
	assert.Equal(t, 7, pos)
}

func TestSimple1(t *testing.T) {
	input := "aaabbbccc"

	output1 := _func.Wordwrap(input, 3, "<br>")
	fmt.Println(output1)

	output2 := _func.Wordwrap(output1, 3, "<br>")
	fmt.Println(output2)

	assert.Equal(t, output1, output2)
}

func TestSimple2(t *testing.T) {
	input := "田中は本日会社にいった。会社にいくと、そこはまるで地獄絵図であった。田中は会社から逃げ出した。"

	output1 := _func.Wordwrap(input, 5, "会社")
	fmt.Println(output1)

	output2 := _func.Wordwrap(output1, 5, "会社")
	fmt.Println(output2)

	assert.Equal(t, output1, output2)
}


func TestSimple3(t *testing.T) {
	//php > $input = 'aa<br>abbbccc<br>dddee<br>e';
	//php > echo wordwrap($input, 3, '<br>', true);
	//aa<br>abb<br>bcc<br>c<br>ddd<br>ee<br>e

	input := "aa<br>abbbccc<br>dddee<br>e"

	output1 := _func.Wordwrap(input, 3, "<br>")
	fmt.Println(output1)

	output2 := _func.Wordwrap(output1, 3, "<br>")
	fmt.Println(output2)

	assert.Equal(t, output1, output2)
}

func TestHugeString(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./test1.txt")
	input := string(bytes)

	output1 := _func.Wordwrap(input, 100, "\n")
	fmt.Println(md5.Sum([]byte(output1)))

	output2 := _func.Wordwrap(output1, 100, "\n")
	fmt.Println(md5.Sum([]byte(output2)))

	assert.Equal(t, output1, output2)

	fmt.Println(output2)
}

func TestSurrogatePair(t *testing.T) {
	input := "𠮷野家で𠮷野がご飯をたべる"

	output := _func.Wordwrap(input, 3, ",")
	fmt.Println(output)

	assert.Equal(t, "𠮷野家,で𠮷野,がご飯,をたべ,る", output)
}

func TestDefault1(t *testing.T) {
	input := "𠮷野家で𠮷野がご飯をたべる"

	output := _func.Wordwrap(input, -1, "")
	fmt.Println(output)

	assert.Equal(t, "𠮷野家で𠮷野がご飯をたべる", output)
}

