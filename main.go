package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var regex *regexp.Regexp

func init() {
	regex = regexp.MustCompile("[$]?[{%].+[}%]")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:\n   %s <filename>\n", os.Args[0])
		os.Exit(0)
	}
	filename := os.Args[1]
	// filename = ReplaceString(filename)  //perhaps needed
	f, err := ioutil.ReadFile(filename)
	AssetNil(err)
	content := string(f)
	//fmt.Println(content)
	c2 := ReplaceString(content)
	err = ioutil.WriteFile(filename, []byte(c2), 0644)
	AssetNil(err)
	//fmt.Println(c2)
	os.Exit(0)
}

func ReplaceString(s string) string {
	envs := regex.FindAllString(s, -1)
	if len(envs) == 0 {
		return s
	}
	for _, e := range envs {
		if len(e) < 3 {
			continue
		}
		var e2 string
		if strings.HasPrefix(e, "$") {
			e2 = e[2 : len(e)-1]
		} else {
			e2 = e[1 : len(e)-1]
		}
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
