package hakka

import (
	"fmt"
	"strings"

	"github.com/edfan0930/playaudio/wav"
)

func Play(s, c string) {

	str := strings.Split(s, "")
	fileExtension(str)
	fmt.Println("the str", str)
	//	digits := AddWords(str, c)
	//	fmt.Println("the digits", digits)
	wav.PlayAudio(afterDigit(str, "3"))

}

func fileExtension(s []string) {
	for k, v := range s {
		s[k] = v + ".wav"
	}
}

func AddWords(s []string, c string) []string {

	digit := []string{}

	for k, v := range s {
		digit = append(digit, PathAndWAV(v))
		if k == 4 {
			digit = append(digit, PathAndWAV("仟"))
			break
		}

		if k == 3 {
			digit = append(digit, PathAndWAV("佰"))
			break
		}

		if k == 2 {
			digit = append(digit, PathAndWAV("拾"))
			break
		}
	}

	return digit
}

func afterDigit(digit []string, c string) []string {

	digit = append(digit, PathAndWAV("來賓"), PathAndWAV("請到"), PathAndWAV(c), PathAndWAV("號"), PathAndWAV("櫃台"))
	return digit
}

func PathAndWAV(name string) string {
	return name + ".wav"
}
