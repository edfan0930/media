package playsounds

import "fmt"

func Put(path string, filesName []string) {

	for i, v := range filesName {

		filesName[i] = fmt.Sprintf("%s/%s", path, v)
	}
}

// PutWAVExtension
func PutWAVExtension(fileName []string) {

	for i, v := range fileName {
		fileName[i] = v + ".wav"
	}
}
