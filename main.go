package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	opt := Option{}

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"f",
			Value:"./hello.proto",
			Usage:"指定单个proto文件或者多个proto文件夹的位置，默认为当前目录下的test.proto",
			Destination:&opt.ProtoFilePath,
		},
		cli.StringFlag{
			Name:"o",
			Value:"./micro_output",
			Usage:"指定框架代码生成的位置，默认为当前目录下的micro_output文件夹",
			Destination:&opt.OutputFilePath,
		},
		cli.StringFlag{
			Name:"m",
			Value:"micro",
			Usage:"指定生成的工程的模块名，默认为micro",
			Destination:&opt.ProjectModule,
		},
	}

	app.Action = func(c *cli.Context) error {
		err := GenerateAll(&opt)
		if err != nil {
			fmt.Printf("Cli work failed: %v", err)
			return err
		}

		return nil
	}
	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
}
