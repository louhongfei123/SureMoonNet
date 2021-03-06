package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ProtossGenius/SureMoonNet/basis/smn_net"
	"github.com/ProtossGenius/SureMoonNet/rpc_nitf/cltrpc/clt_rpc_rpc_itf"
	"github.com/ProtossGenius/SureMoonNet/rpc_nitf/svrrpc/svr_rpc_rpc_itf"
	"github.com/ProtossGenius/SureMoonNet/smn/net_libs/muti_service"
	"github.com/ProtossGenius/SureMoonNet/smn/net_libs/smn_rpc"
	"github.com/ProtossGenius/SureMoonNet/test/rpc_itfs/rpc_itf"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type login struct {
	rpc_itf.Login
}

func (this *login) DoLogin(user, pswd string, code int) (bool, int) {
	fmt.Println(user, pswd, code)
	return false, 0
}

func (this *login) Test1(a []string, b []int, c []uint, d []uint64, e []int32) []int {
	panic("implement me")
}
func (this *login) Test2(key string, c net.Conn) bool {
	if key == "hello" {
		smn_net.WriteString("world", c)
		return true
	} else {
		smn_net.WriteString("where is hello?", c)
		return false
	}
}

func accept(conn net.Conn) {
	sm := smn_rpc.NewRPCServiceManager(conn)
	smn_rpc.ServiceManagerRegister(sm, 404, "login", svr_rpc_rpc_itf.NewSvrRpcLogin(&login{}))
	go sm.Work()
}

func RunSvr() {
	svr, err := smn_net.NewTcpServer(1000, accept)
	check(err)
	svr.Run()
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	go RunSvr()
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:1000")
	check(err)
	sm := muti_service.NewServiceManager(conn)
	fc, _ := sm.Regitster(404, "login")
	go sm.Work()
	client := clt_rpc_rpc_itf.NewCltRpcLogin(smn_rpc.NewMessageAdapter(fc))
	b, i := client.DoLogin("user---", "pswd_____", -1)
	fmt.Println("cccccc   ", b, i)
	t2f := func(c net.Conn) {
		str, _ := smn_net.ReadString(c)
		fmt.Printf("login.Test2 cFunc, stream val: %s", str)
	}
	fmt.Println(client.Test2("hello", t2f))
	fmt.Println(client.Test2("helle", t2f))
	client.Test1(nil, nil, nil, nil, nil)
}
