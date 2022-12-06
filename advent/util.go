package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type _string struct {
	Value string
}

func ReadInput(directory string) _string {
	content, err := ioutil.ReadFile(fmt.Sprintf("advent/%s/input.txt", directory))
	if err != nil {
		log.Fatal(err)
	}
	return _string{Value: string(content)}
}

func (str *_string) Split(separator string) []string {
	splitArray := strings.Split(str.Value, separator)
	return splitArray[:len(splitArray)-1]
}

func Trim(input []string) []string {
	trimmedInput := []string{}
	for _, element := range input {
		trimmedInput = append(trimmedInput, strings.Replace(string(element), " ", "", -1))
	}
	return trimmedInput
}
