package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"strings"
)

type Generator interface {
	Generate(opt *Option, services *ServiceInfo) error
}

type Generators struct {
	generator map[string]Generator
}

var generators = &Generators{
	generator: make(map[string]Generator),
}

func Register(name string, generator Generator) {
	generators.generator[name] = generator
}

func GenerateAll(opt *Option) (err error) {

	//创建项目目录
	err = generators.generator["directory_generator"].Generate(opt, nil)
	if err != nil {
		return
	}

	//解析proto文件
	service, err := parseProtoFile(opt)
	if err != nil {
		return err
	}

	//生成控制器文件
	err = generators.generator["controller_generator"].Generate(opt, service)
	if err != nil {
		return err
	}

	//生成路由文件
	err = generators.generator["router_generator"].Generate(opt, service)
	if err != nil {
		return err
	}

	//生成中间件文件
	err = generators.generator["middleware_generator"].Generate(opt, service)
	if err != nil {
		return err
	}

	//

	return
}

func parseProtoFile(opt *Option) (service *ServiceInfo, err error) {
	service = new(ServiceInfo)
	fileInfo, err := os.Stat(opt.ProtoFilePath)
	if err != nil {
		err = fmt.Errorf("Check file %s failed: %v\n", opt.ProtoFilePath, err)
		return
	}
	if fileInfo.IsDir() {
		err = fmt.Errorf("File %s is not a proto file, but a directory. ", opt.ProtoFilePath)
		return
	}else {
		if fileInfo.Name()[strings.LastIndex(fileInfo.Name(), "."):] != "proto" {
			err = fmt.Errorf("File %s is not a proto file. ", opt.ProtoFilePath)
		}
		file, err := os.Open(opt.ProtoFilePath)
		if err != nil {
			err = fmt.Errorf("Open file %s failed: %v. ", opt.ProtoFilePath, err)
			return
		}

		parser := proto.NewParser(file)
		p, err := parser.Parse()
		if err != nil {
			fmt.Printf("Parse file %s failed: %v", fileInfo.Name(), err)
			return
		}

		proto.Walk(p,
			proto.WithService(service.handleService),
			proto.WithRPC(service.handleRpc),
			proto.WithMessage(service.handlerMessage),
		)
	}

	service.Module = opt.ProjectModule

	return
}

type ServiceInfo struct {
	Service  *proto.Service
	Rpcs     []*proto.RPC
	Messages []*proto.Message
	Module   string
}

func (s *ServiceInfo) handleService(service *proto.Service)  {
	s.Service = service
}

func (s *ServiceInfo) handlerMessage(message *proto.Message)  {
	s.Messages = append(s.Messages, message)
}

func (s *ServiceInfo) handleRpc(rpc *proto.RPC)  {
	s.Rpcs = append(s.Rpcs, rpc)
}