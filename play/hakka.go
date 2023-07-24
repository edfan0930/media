package play

import (
	"strconv"
	"strings"
)

func HakkaRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)
	callNumberSlice := strings.Split(n, "")

	all := []string{"來賓"}

	PutMiddle(&callNumberSlice)
	all = append(all, callNumberSlice...)
	all = append(all, "號", "請到", c, "號", "櫃台")
	Put(HakkaPath, all)
	PutWAVExtension(all)

	return all

}
