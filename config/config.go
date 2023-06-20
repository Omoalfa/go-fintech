package config

import (
	"os"
	"strings"
)

var Config = make(map[string]any)
var Test string

func SetUpConfig() {
	for _, value := range os.Environ() {
		data := strings.SplitN(value, "=", 2)
		Config[data[0]] = data[1]
	}
	// fmt.Println(os.Environ())
}
