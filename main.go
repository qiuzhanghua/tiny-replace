package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var regex *regexp.Regexp

func init() {
	if runtime.GOOS == "Windows" {
		regex = regexp.MustCompile("[%].+[%]")
	} else {
		regex = regexp.MustCompile("\\$\\{.+\\}")
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:\n   %s <filename>\n", os.Args[0])
		os.Exit(0)
	}
	filename := os.Args[1]
	// filename = ReplaceString(filename)
	f, err := ioutil.ReadFile(filename)
	AssetNil(err)
	content := string(f)
	c2 := ReplaceString(content)
	err = ioutil.WriteFile(filename, []byte(c2), 0644)
	AssetNil(err)
	os.Exit(0)
}

func ReplaceString(s string) string {
	envs := regex.FindAllString(s, -1)
	for _, e := range envs {
		e2 := e[2 : len(e)-1]
		env := os.Getenv(e2)
		s = strings.ReplaceAll(s, e, env)
	}
	return s
}

func AssetNil(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
