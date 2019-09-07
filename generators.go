package main

type Generator interface {
	Generate(opt *Option) error
}

type Generators struct {
	generator map[string]Generator
}

var generators = &Generators{
	generator:make(map[string]Generator),
}

func Register(name string, generator Generator)  {
	generators.generator[name] = generator
}

func GenerateAll(opt *Option) (err error) {
	err = generators.generator["directory_generator"].Generate(opt)
	if err != nil {
		return
	}

	//解析proto文件


	return
}