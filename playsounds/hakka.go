package playsounds

import (
	"strconv"
	"strings"
)

func HakkaRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)
	callNumberSlice := strings.Split(n, "")
	counterNumberSlice := strings.Split(c, "")

	all := []string{"來賓"}

	PutMiddle(&callNumberSlice)
	PutMiddle(&counterNumberSlice)

	all = append(all, callNumberSlice...)
	all = append(all, AfterNumber(counterNumberSlice)...)
	Put(HakkaPath, all)
	PutWAVExtension(all)

	return all
}
