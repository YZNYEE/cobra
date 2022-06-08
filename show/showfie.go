package show

import "os"
import "fmt"

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
