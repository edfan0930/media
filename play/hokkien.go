package play

import (
	"fmt"
	"strconv"
	"strings"
)

func HokkienRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)
	callNumberSlice := strings.Split(n, "")

	all := []string{"來賓"}

	PutMiddle(&callNumberSlice)
	all = append(all, callNumberSlice...)
	all = append(all, "號", "請到", c, "號", "櫃台")
	Put(HokkienPath, all)
	PutWAVExtension(all)

	return all
}

func PutMiddle(n *[]string) {

	number := *n
	length := len(number)
	for i := range number {

		if (length - i) == 4 {
			number = append(number[:i+1], append([]string{"仟"}, number[i+1:]...)...)
		}
		if (length - i) == 3 {
			number = append(number[:i+1], append([]string{"佰"}, number[i+1:]...)...)
		}
		if (length - i) == 2 {
			fmt.Print("in the ten")
			number = append(number[:i+1], append([]string{"拾"}, number[i+1:]...)...)
			fmt.Println("over append", number)
		}
	}
	*n = number
}
