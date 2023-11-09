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
	filename := os.Args[1]
	if filename == "version" {
		fmt.Printf("%s %s (%s %s)\n", "tiny-replace", AppVersion, AppRevision, AppBuildDate)
		os.Exit(0)
	}
	// filename = ReplaceString(filename)  //perhaps needed
	f, err := os.ReadFile(filename)
	AssetNil(err)
	content := string(f)
	c2 := util.ReplaceString(content, os.Args[1:]...)
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
