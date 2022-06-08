package show

import "os"
import "fmt"
import "cobradir/utils"
import "cobradir/clcolor"

func Showlist(filelist []os.FileInfo) {
	//n := len(filelist)
	fmt.Println("文件名", "||", "大小", "||", "模式", "||", "修改时间", "||", "目录？")
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
func Showtree(filelist []os.FileInfo, upper int, dir string, hide bool) {
	root := utils.Levelsearch(filelist, upper, dir)
	printprefix(root, "", hide)

}
func printprefix(node utils.FileNode, prefix string, hide bool) {
	//fmt.Println("prefix:", prefix)
	//fmt.Println(node.GetVal())
	child := node.GetChild()
	if child != nil {
		fmt.Println("==================================================================")
		for i := 0; i < len(child); i++ {
			cur := child[i]
			name := cur.GetVal().Name()
			//fmt.Println(name)
			if name[0] == byte('.') && !hide {
				fmt.Println("comtinue")
				continue
			}
			s := utils.ShortString(cur.GetVal())
			outs := prefix + s
			outs = clcolor.Blue(outs)
			fmt.Println(outs)
			printprefix(cur, prefix+"->->", hide)
		}
		fmt.Println("==================================================================")
	}
}
