package td

func IsLeapYear(nYear int) bool{
	if (nYear % 4 == 0 && nYear % 100 != 0) || nYear % 400 == 0 {
		return true
	}
	return false
}

//获取每年每个月得具体天数
func MonthDays( iYear ,iMonth int) int{
	switch iMonth {
	case 1, 3, 5, 7, 8, 10, 12:{
		return 31
	}
	case 4, 6, 9, 11:{
		return 30
	}
	case 2:{
		if IsLeapYear(iYear){
			return 29
		}
		return 28
	}
	}
	return 0
}

