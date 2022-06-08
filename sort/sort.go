package sort

import "os"
import "sort"
import "cobradir/utils"

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

func SortFileNode(files []utils.FileNode, mod string) []utils.FileNode {
	switch mod {
	case "":
		return files
	case "n":
		nsort := FSortnSlice(files)
		sort.Sort(nsort)
		return []utils.FileNode(nsort)
	case "s":
		nsort := FSortsSlice(files)
		sort.Sort(nsort)
		return []utils.FileNode(nsort)
	case "t":
		nsort := FSorttSlice(files)
		sort.Sort(nsort)
		return []utils.FileNode(nsort)
	default:
		return files

	}
}
