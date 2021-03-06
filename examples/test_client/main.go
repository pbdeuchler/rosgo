package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ppg/rosgo/examples/srv"
	"github.com/ppg/rosgo/ros"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("USAGE: test_client <int> <int>")
		os.Exit(1)
	}

	node := ros.NewNode("client")
	defer node.Shutdown()
	logger := node.Logger()
	logger.SetSeverity(ros.LogLevelDebug)
	cli := node.NewServiceClient("/add_two_ints", srvs.SrvAddTwoInts)
	defer cli.Shutdown()
	var err error
	var a, b int64
	a, err = strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Print(err)
		fmt.Println()
		os.Exit(1)
	}
	b, err = strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Print(err)
		fmt.Println()
		os.Exit(1)
	}
	var srv srvs.AddTwoInts
	srv.Request.A = int32(a)
	srv.Request.B = int32(b)
	if err := cli.Call(&srv); err != nil {
		fmt.Print(err)
		fmt.Println()
	} else {
		fmt.Printf("%d + %d = %d",
			srv.Request.A, srv.Request.B, srv.Response.Sum)
		fmt.Println()
	}
}
