package playsounds

import (
	"fmt"
	"strconv"
	"strings"
)

func MandarinRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)
	callNumberSlice := strings.Split(n, "")
	counterNumberSlice := strings.Split(c, "")

	all := []string{"來賓"}

	if len(n) <= 2 {
		LessHundred(&callNumberSlice)
		LessHundred(&counterNumberSlice)
	}

	all = append(all, callNumberSlice...)
	all = append(all, AfterNumber(counterNumberSlice)...)

	Put(MandarinPath, all)
	PutWAVExtension(all)
	return all
}

// LessHundred
func LessHundred(n *[]string) {

	number := *n

	for i, v := range number {
		if i == 0 {
			number[i] = fmt.Sprintf("%s%s", v, "0")
		}
	}
}
