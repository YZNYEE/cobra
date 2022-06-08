package sort

import "os"
import "sort"

func SortFile(files []os.FileInfo, mod string) []os.FileInfo {
	switch mod {
	case "":
		return files
	case "n":
		nsort := NameSort(files)
		sort.Sort(nsort)
		return []os.FileInfo(nsort)
	case "s":
		nsort := SizeSort(files)
		sort.Sort(nsort)
		return []os.FileInfo(nsort)
	case "t":
		nsort := timeSort(files)
		sort.Sort(nsort)
		return []os.FileInfo(nsort)
	default:
		return files

	}
}
