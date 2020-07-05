package td

import (
	"fmt"
	"testing"
)

func TestMonthDays(t *testing.T) {
	 days := MonthDays(1992,10)
	 fmt.Println(days)
}
