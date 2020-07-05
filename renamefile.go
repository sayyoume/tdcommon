package td

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//批量修改文件名字
/*
   k1_ggg.mp3   k1.mp3
   k2_ggg.mp3   k2.mp3
   k3_ggg.mp3   k3.mp3
   k4_ggg.mp3   k4.mp3
*/

func RenameFiles(path, subString string){

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		sDst := strings.Replace(f.Name(),subString,"",1)
		os.Rename(path +"\\" + f.Name(), path +"\\" + sDst)
	}
}
