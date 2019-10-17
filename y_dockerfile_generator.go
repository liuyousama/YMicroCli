package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type DockerfileGenerator struct {

}

func init()  {
	Register("dockerfile_generator", new(DockerfileGenerator))
}

func (*DockerfileGenerator)Generate(opt *Option, service *ServiceInfo) (err error) {
	var file *os.File
	defer func() {_ = file.Close()}()
	file, err = os.OpenFile(filepath.Join(opt.OutputFilePath, "Dockerfile"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		err = fmt.Errorf("Create Dockerfile failed: %v. ", err)
		return
	}

	dockerfileCode := fmt.Sprintf(dockerfileCodeFmt, 8080, service.Module)

	_, err = fmt.Fprint(file, dockerfileCode)
	if err != nil {
		err = fmt.Errorf("Write code to Dockerfile failed: %v. ", err)
		return
	}

	return
}

var dockerfileCodeFmt = `
#源镜像
FROM golang:latest
#作者
MAINTAINER liuyousama "634308664@qq.com"
#设置go mod以及proxy相关
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
ENV GO111MODULE on

#将服务器的go工程代码加入到docker容器中
ADD . /home/project
#设置工作目录
WORKDIR /home/project
#go mod依赖
RUN go mod tidy
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE %d
#最终运行docker的命令
ENTRYPOINT  ["./%s"]
`