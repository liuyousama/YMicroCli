package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"io/ioutil"
	"os"
)

type Generator interface {
	Generate(opt *Option, services []*ServiceInfo) error
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
	services, err := parseProtoFile(opt)
	if err != nil {
		return err
	}

	//生成控制器文件
	err = generators.generator["controller_generator"].Generate(opt, services)
	if err != nil {
		return err
	}

	//

	return
}

func parseProtoFile(opt *Option) (services []*ServiceInfo, err error) {
	services = make([]*ServiceInfo, 0)
	fileInfo, err := os.Stat(opt.ProtoFilePath)
	if err != nil {
		fmt.Printf("Check file %s failed: %v\n", opt.ProtoFilePath, err)
	}
	if fileInfo.IsDir() {
		fileList, err := ioutil.ReadDir(opt.ProtoFilePath)
		if err != nil {
			fmt.Printf("Open directory %s failed: %v\n", opt.ProtoFilePath, err)
			return
		}
		for _, fileInfo := range fileList {
			file, err := os.Open(fileInfo.Name())
			if err != nil {
				fmt.Printf("Open file %s failed: %v", fileInfo.Name(), err)
				return
			}

			parser := proto.NewParser(file)
			p, err := parser.Parse()
			if err != nil {
				fmt.Printf("Parse file %s failed: %v", fileInfo.Name(), err)
				return
			}

			service := new(ServiceInfo)

			proto.Walk(p,
				proto.WithService(service.handleService),
				proto.WithRPC(service.handleRpc),
				proto.WithMessage(service.handlerMessage),
			)

			services = append(services, service)
		}
	}else {
		file, err := os.Open(opt.ProtoFilePath)
		if err != nil {
			fmt.Printf("Open file %s failed: %v", opt.ProtoFilePath, err)
			return
		}

		parser := proto.NewParser(file)
		p, err := parser.Parse()
		if err != nil {
			fmt.Printf("Parse file %s failed: %v", fileInfo.Name(), err)
			return
		}

		service := new(ServiceInfo)

		proto.Walk(p,
			proto.WithService(service.handleService),
			proto.WithRPC(service.handleRpc),
			proto.WithMessage(service.handlerMessage),
		)

		services = append(services, service)
	}

	return
}

type ServiceInfo struct {
	Service  *proto.Service
	Rpcs     []*proto.RPC
	Messages []*proto.Message
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
