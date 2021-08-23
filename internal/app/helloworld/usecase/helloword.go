/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/19 16:02
 */
package usecase

import (
	"context"
	"fmt"

	"app/api/helloworld"
	"app/ecode"
)

// server is used to implement helloworld.GreeterServer.
type HelloWorldUseCase struct {
	helloworld.UnimplementedGreeterServer
}

func NewHelloWorldUseCase() helloworld.IGreeterUseCase {
	return &HelloWorldUseCase{}
}

// SayHello implements helloworld.GreeterServer
func (s *HelloWorldUseCase) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {

	if in.Name == "error" {
		return nil, ecode.ErrInvalidArgument
	}
	if in.Name == "panic" {
		panic("server panic")
	}

	return &helloworld.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}
