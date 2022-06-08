package sort

import "os"

type NameSort []os.FileInfo
type SizeSort []os.FileInfo
type timeSort []os.FileInfo

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
