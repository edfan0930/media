package playsounds

import (
	"fmt"
	"strconv"
	"strings"
)

func MandarinRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)

	var callNumberSlice, counterNumberSlice []string

	all := []string{"來賓"}
	if callNumber >= 10 && callNumber < 100 && callNumber%10 == 0 {
		callNumberSlice = []string{strconv.Itoa(callNumber)}
	} else {

		callNumberSlice = strings.Split(n, "")
		if len(n) == 2 {

			LessHundred(&callNumberSlice)
		}
	}
	//如果為十位數, 並且個位數為0
	if counterNumber >= 10 && counterNumber < 100 && counterNumber%10 == 0 {
		counterNumberSlice = []string{strconv.Itoa(counterNumber)}

	} else {

		counterNumberSlice = strings.Split(c, "")
		if len(c) == 2 {

			LessHundred(&counterNumberSlice)
		}
	}

	all = append(all, callNumberSlice...)
	all = append(all, AfterNumber(counterNumberSlice)...)

	Put(MandarinPath, all)
	PutWAVExtension(all)
	return all
}

// MandarinWithSingle
func SingleWithMandarin(number int) []string {

	str := strconv.Itoa(number)

	// 使用 strings.Split 将字符串切分为字符切片
	chars := strings.Split(str, "")
	chars = append(chars, "號")
	Put(MandarinPath, chars)
	PutWAVExtension(chars)

	return chars
}

// LessHundred
func LessHundred(n *[]string) {

	number := *n
	if number[len(number)-1] == "0" {
		return
	}
	for i, v := range number {
		if i == 0 {
			number[i] = fmt.Sprintf("%s%s", v, "0")
		}
	}
}
