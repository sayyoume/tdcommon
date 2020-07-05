package td

import (
	"fmt"
	"strings"
	"testing"
)

func TestMonthDays(t *testing.T) {
	var  s1 ="k1_goo.mp3"
	 s1 = strings.Replace(s1,"_goo","",1)
	 days := MonthDays(1992,10)
	 fmt.Println(days)
}
