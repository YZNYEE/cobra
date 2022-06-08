package utils

import (
	"fmt"
	"os"
)
import "io/ioutil"

type FileNode struct {
	val   os.FileInfo
	child []FileNode
	level int
}

func ToString(file os.FileInfo) string {
	name := file.Name()
	size := file.Size()
	filemode := file.Mode()
	time := file.ModTime()
	isdir := file.IsDir()
	s := fmt.Sprint(name, "||", size, "||", filemode, "||", time, "||", isdir, "\n")
	return s
}

func ShortString(file os.FileInfo) string {
	name := file.Name()
	size := file.Size()
	s := fmt.Sprint(name, "||", size)
	return s
}

func (receiver FileNode) GetChild() []FileNode {
	return receiver.child
}
func (receiver FileNode) GetVal() os.FileInfo {
	return receiver.val
}

func Levelsearch(file []os.FileInfo, upper int, dir string) FileNode {
	root := FileNode{nil, nil, 0}
	root.child = search(file, 1, upper, dir)
	return root
}

func search(file []os.FileInfo, level int, upper int, predir string) []FileNode {
	if level > upper {
		return nil
	}
	res := make([]FileNode, 0)
	for i := 0; i < len(file); i++ {
		f := file[i]
		d := f.IsDir()
		a := FileNode{f, nil, level}
		if d {
			nf := predir + "\\" + f.Name()
			nextfile, _ := ioutil.ReadDir(nf)
			//fmt.Println(nf)
			//fmt.Println(nextfile)
			a.child = search(nextfile, level+1, upper, nf)
		}
		res = append(res, a)
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
