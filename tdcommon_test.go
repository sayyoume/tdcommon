package td

import (
	"fmt"
	"testing"
)

func TestGenerateRangeNum(t *testing.T) {

	for i := 1;i<10;i++{
		res := GenerateRangeNum(100,10000)
		fmt.Println(res)
	}



}
