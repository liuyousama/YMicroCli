package main

import (
	"fmt"
	"os"
	"path"
)

var AllDirectoryList []string = []string{
	"controller",
	"proto",
	"main",
	"scripts",
	"model",
	"config",
	"app/router",
	"app/config",
	"grpc",
}

type DirectoryGenerator struct {
	fileList []string
}

func init()  {
	directoryGenerator := &DirectoryGenerator{
		fileList:AllDirectoryList,
	}

	Register("directory_generator", directoryGenerator)
}

func (g *DirectoryGenerator) Generate(opt *Option) (err error) {
	for _, v := range g.fileList {
		err =  os.MkdirAll(path.Join(opt.OutputFilePath,v), 0755)
		if err != nil {
			fmt.Printf("make dir %s error:%v", v, err)
			return
		}
	}

	return
}

