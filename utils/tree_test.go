package utils

import (
	"cobradir/clcolor"
	"fmt"
	"io/ioutil"
	"testing"
)

var file string

func TestLevelsearch(t *testing.T) {
	file := "D:\\goproject"
	finfo, _ := ioutil.ReadDir(file)
	root := Levelsearch(finfo, 2, file)
	show(root)
}

func show(node FileNode) {
	c := node.GetChild()
	fmt.Println(c)
	if c == nil {
		return
	}
	for _, v := range c {
		s := fmt.Sprint(v.GetVal())
		cs := clcolor.Blue(s)
		fmt.Printf(cs)
		show(v)
	}
}
