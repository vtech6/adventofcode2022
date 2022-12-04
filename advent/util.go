package advent

import (
	"io/ioutil"
	"log"
	"strings"
)

type _string struct {
	value string
}

func ReadInput(directory string) _string {
	content, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Fatal(err)
	}
	return _string{value: string(content)}
}

func (str *_string) Split(separator string) []string {
	splitArray := strings.Split(str.value, separator)
	return splitArray[:len(splitArray)-1]
}
