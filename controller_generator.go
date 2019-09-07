package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"path/filepath"
)

type ControllerGenerator struct {
	service  *proto.Service
	rpcs     []*proto.RPC
	messages []*proto.Message
}

func init() {
	controllerGenerator := &ControllerGenerator{

	}

	Register("controller_generator", controllerGenerator)
}

func (g *ControllerGenerator) Generate(opt *Option) (err error) {
	//先读取指定的proto文件，将文件交给解析器解析
	file, err := os.Open(opt.ProtoFilePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Open file %s failed: %v\n", opt.ProtoFilePath, err)
		return
	}
	parser := proto.NewParser(file)
	info, err := parser.Parse()
	if err != nil {
		fmt.Printf("Pars file %s failed: %v\n", opt.ProtoFilePath, err)
		return
	}

	proto.Walk(info,
		proto.WithMessage(g.handleMessage),
		proto.WithRPC(g.handleRpc),
		proto.WithService(g.handleService),
	)

	return
}

func (g *ControllerGenerator) handleService(service *proto.Service) {
	g.service = service
}

func (g *ControllerGenerator) handleMessage(message *proto.Message) {
	g.messages = append(g.messages, message)
}

func (g *ControllerGenerator) handleRpc(rpc *proto.RPC) {
	g.rpcs = append(g.rpcs, rpc)
}

func (g *ControllerGenerator) generateCode(opt *Option) (err error) {
	filePath := filepath.Join(opt.OutputFilePath, "controller", fmt.Sprintf("%s.go", g.service.Name))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		fmt.Printf("create controller file %s failed: %v\n", g.service.Name, err)
		return
	}


}