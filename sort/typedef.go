package sort

import "os"
import "cobradir/utils"

type NameSort []os.FileInfo
type SizeSort []os.FileInfo
type timeSort []os.FileInfo

type FSortnSlice []utils.FileNode
type FSortsSlice []utils.FileNode
type FSorttSlice []utils.FileNode

func (receiver NameSort) Len() int {
	return len(receiver)
}
func (receiver NameSort) Less(i, j int) bool {
	return receiver[i].Name() < receiver[j].Name()
}
func (receiver NameSort) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver SizeSort) Len() int {
	return len(receiver)
}
func (receiver SizeSort) Less(i, j int) bool {
	return receiver[i].Size() < receiver[j].Size()
}
func (receiver SizeSort) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func (receiver timeSort) Len() int {
	return len(receiver)
}
func (receiver timeSort) Less(i, j int) bool {
	return receiver[i].ModTime().Unix() < receiver[j].ModTime().Unix()
}
func (receiver timeSort) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func (receiver FSortnSlice) Len() int {
	return len(receiver)
}
func (receiver FSortnSlice) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver FSortnSlice) Less(i, j int) bool {
	return receiver[i].GetVal().Name() < receiver[j].GetVal().Name()
}

func (receiver FSorttSlice) Len() int {
	return len(receiver)
}
func (receiver FSorttSlice) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver FSorttSlice) Less(i, j int) bool {
	return receiver[i].GetVal().ModTime().Unix() < receiver[j].GetVal().ModTime().Unix()
}
func (receiver FSortsSlice) Len() int {
	return len(receiver)
}
func (receiver FSortsSlice) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}
func (receiver FSortsSlice) Less(i, j int) bool {
	return receiver[i].GetVal().Size() < receiver[j].GetVal().Size()
}
