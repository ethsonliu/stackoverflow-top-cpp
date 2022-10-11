package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	readme, err := os.OpenFile("README.md", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	readmeB, err := io.ReadAll(readme)
	if err != nil {
		panic(err)
	}
	readmeS := string(readmeB)
	flag := "<!-- catalog -->"
	splits := strings.Split(readmeS, flag)
	ctl := strings.Builder{}
	ctl.WriteString("\n\n")
	err = filepath.Walk("question", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || !strings.Contains(info.Name(), ".md") {
			return nil
		}
		ctl.WriteString(fmt.Sprintf("- [%s](%s)\n", strings.Trim(info.Name(), "[]"), strings.ReplaceAll(path, " ", "%20")))
		return nil
	})
	readmeS = splits[0] + flag + ctl.String() + "\n" + flag + splits[2]
	if err != nil {
		panic(err)
	}
	err = readme.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = readme.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	_, err = readme.WriteString(readmeS)
	if err != nil {
		panic(err)
	}
}
