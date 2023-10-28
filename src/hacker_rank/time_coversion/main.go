package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	aux := strings.Contains(s, "PM")
	var strF string

	if aux {
		str := strings.Split(s, ":")
		fmt.Println(str[0])
		num, _ := strconv.Atoi(str[0])
		numF := ""
		if num == 12 {
			numF = strconv.Itoa(num)
		} else {
			numF = strconv.Itoa(num + 12)
		}
		strF = numF + ":" + str[1] + ":" + str[2]
	} else {
		str := strings.Split(s, ":")
		fmt.Println(str[0])
		num, _ := strconv.Atoi(str[0])
		numF := ""
		if (num - 12) == 0 {
			numF = "0" + strconv.Itoa(num-12)
		} else if (num - 12) < 0 {
			numF = "0" + strconv.Itoa(num)
		} else {
			numF = strconv.Itoa(num - 12)
		}

		strF = numF + ":" + str[1] + ":" + str[2]
	}

	return strF[0:8]
}

func main() {
	//str := "07:05:45:PM"
	str2 := "06:05:45AM"
	strt := timeConversion(str2)

	fmt.Println(strt)
}
