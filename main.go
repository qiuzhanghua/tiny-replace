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
	regex = regexp.MustCompile("[$][{].+?[}]|[%].+?[%]")
	// "[$]?[{%].+[}%]"
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
	c2 := ReplaceString(content, os.Args[1:]...)
	err = ioutil.WriteFile(filename, []byte(c2), 0644)
	AssetNil(err)
	os.Exit(0)
}

func ReplaceString(s string, rep ...string) string {
	m := make(map[string]string, len(rep))
	for _, r := range rep {
		arr := strings.Split(r, "=")
		if len(arr) == 2 {
			m[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		}
	}
	envs := regex.FindAllString(s, -1)
	if len(envs) == 0 && len(m) == 0 {
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
		if len(m) > 0 {
			val, ok := m[e2]
			if ok {
				s = strings.ReplaceAll(s, e, val)
			}
		} else {
			env, ok := os.LookupEnv(e2)
			if !ok {
				continue
			}
			if runtime.GOOS == "windows" {
				env = strings.ReplaceAll(env, "\\", "\\\\")

			}
			s = strings.ReplaceAll(s, e, env)
		}
	}
	return s
}

func AssetNil(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
