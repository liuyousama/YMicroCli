package main


var mainTemplate = `
package main
import(
	"github.com/grpc/grpc-go"
	"context"
	""
)

func main () {
	grpcServer := grpc.NewServer()
	{{range .}}
    Register{{.Service.Name}}Server(grpcServer, new(HelloServiceImpl))
	{{end}}
	
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(listener)
}
`