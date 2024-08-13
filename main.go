package main

import (
	"fmt"
	"github.com/qiuzhanghua/common/util"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:\n   %s <filename>\n", os.Args[0])
		os.Exit(0)
	}
	hasMode := false
	var replaceMode util.ReplaceMode = util.Keep
	if os.Args[1] == "-b" {
		replaceMode = util.BackSlash
		hasMode = true
	} else if os.Args[1] == "-s" {
		replaceMode = util.Slash
		hasMode = true
	} else if os.Args[1] == "-d" {
		replaceMode = util.DoubleBackSlash
		hasMode = true
	} else if os.Args[1] == "-v" || os.Args[1] == "version" {
		fmt.Printf("%s %s (%s %s)\n", "tiny-replace", AppVersion, AppRevision, AppBuildDate)
		os.Exit(0)
	}
	args := os.Args[1:]
	if hasMode {
		args = os.Args[2:]
	}
	filename := args[0]
	fmt.Println(filename)
	// filename = ReplaceString(filename)  //perhaps needed
	f, err := os.ReadFile(filename)
	AssetNil(err)
	content := string(f)
	c2 := util.ReplaceStringWithMode(content, replaceMode, args...)
	err = os.WriteFile(filename, []byte(c2), 0644)
	AssetNil(err)
	os.Exit(0)
}

func AssetNil(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
