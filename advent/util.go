package advent

import (
	"io/ioutil"
	"log"
)

func ReadInput(directory string) string {
	content, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
