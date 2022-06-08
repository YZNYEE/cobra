package show

import (
	"cobradir/sort"
	"os"
)
import "fmt"
import "cobradir/utils"
import "cobradir/clcolor"

func Showlist(filelist []os.FileInfo, sortmod string) {
	//n := len(filelist)
	fmt.Println("文件名", "||", "大小", "||", "模式", "||", "修改时间", "||", "目录？")
	filelist = sort.SortFile(filelist, sortmod)
	for i := range filelist {
		file := filelist[i]
		name := file.Name()
		size := file.Size()
		filemode := file.Mode()
		time := file.ModTime()
		isdir := file.IsDir()
		fmt.Println(i, "-->:", name, "||", size, "||", filemode, "||", time, "||", isdir)
	}
}
func Showtree(filelist []os.FileInfo, upper int, dir string, hide bool, sortmod string) {
	root := utils.Levelsearch(filelist, upper, dir)
	printprefix(root, "", hide, sortmod)

}
func printprefix(node utils.FileNode, prefix string, hide bool, sortmod string) {
	//fmt.Println("prefix:", prefix)
	//fmt.Println(node.GetVal())
	child := node.GetChild()
	//fmt.Println("sortmod:", sortmod)
	childs := sort.SortFileNode(child, sortmod)
	fmt.Println()
	if child != nil {
		fmt.Println("==================================================================")
		for i := 0; i < len(childs); i++ {
			cur := childs[i]
			name := cur.GetVal().Name()
			//fmt.Println(name)
			if name[0] == byte('.') && !hide {
				continue
			}
			s := utils.ShortString(cur.GetVal())
			outs := prefix + s
			outs = clcolor.Blue(outs)
			fmt.Println(outs)
			printprefix(cur, prefix+"->->", hide, sortmod)
		}
		fmt.Println("==================================================================")
	}
}
