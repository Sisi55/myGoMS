package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/building-microservices-with-go/chapter1/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld) // 핸들러 등록 ?

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
//func Register(rcvr interface{}) error { return DefaultServer.Register(rcvr) }

//func Listen(network, address string) (Listener, error) {
//	var lc ListenConfig
//	return lc.Listen(context.Background(), network, address)
//}

//type Listener interface {
//	Accept() (Conn, error)
//	Close() error
//	Addr() Addr
//}

//func ServeConn(conn io.ReadWriteCloser) {
//	DefaultServer.ServeConn(conn)
//}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

// D:\GO_WORKSPACE\MSGo\ch1\rpc\server\server.go
