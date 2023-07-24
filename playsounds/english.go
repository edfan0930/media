package playsounds

import (
	"strconv"
	"strings"
)

// EnglishRules
func EnglishRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)

	all := []string{"來賓", "號"}

	all = append(all, strings.Split(n, "")...)

	all = append(all, append([]string{"請到", "櫃台", "號"}, strings.Split(c, "")...)...)

	Put(EnglishPath, all)
	PutWAVExtension(all)

	return all
}
