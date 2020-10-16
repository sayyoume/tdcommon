package td

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
func GenerateRangeNum(min, max int) int {
	randNum := r.Intn(max - min) + min
	return randNum
}

func TdEncode(str string){
	//字符转换16进制
	src := []byte(str)
	encodedStr := hex.EncodeToString(src)
	// [72 101 108 108 111]
	//fmt.Println(src)
	fmt.Println(encodedStr)
}

func TdDecode(str string){
	//16进制转换字符串
	decodestr ,_:= hex.DecodeString(str)
	fmt.Println(string(decodestr))
}
