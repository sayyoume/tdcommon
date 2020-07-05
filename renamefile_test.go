package td

import "testing"

func TestRenameFiles(t *testing.T) {
	path := "F:\\BaiduYunDownload\\golang视频\\尚硅谷 Go语言核心编程课程\\视频-4（更新）"
	subString :="_Go核心编程"
	RenameFiles(path,subString)
}