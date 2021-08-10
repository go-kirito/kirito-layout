// Code generated by protoc-gen-go-kirito. DO NOT EDIT.
// versions:
// - protoc-gen-go-kirito v0.0.1
// source: api/heloworld/v1/helloworld.proto

package v1

import (
	context "context"
	application "github.com/go-kirito/pkg/application"
)

// IHelloworldUseCase is the server API for Helloworld service.
// for forward compatibility
type IHelloworldUseCase interface {
	CreateHelloworld(context.Context, *CreateHelloworldRequest) (*CreateHelloworldReply, error)
	UpdateHelloworld(context.Context, *UpdateHelloworldRequest) (*UpdateHelloworldReply, error)
	DeleteHelloworld(context.Context, *DeleteHelloworldRequest) (*DeleteHelloworldReply, error)
	GetHelloworld(context.Context, *GetHelloworldRequest) (*GetHelloworldReply, error)
	ListHelloworld(context.Context, *ListHelloworldRequest) (*ListHelloworldReply, error)
}

func RegisterHelloworldServer(app application.application, srv IHelloworldUseCase) {
	if app.HttpServer() != nil {
		RegisterGreeterHTTPServer(app.HttpServer(), srv)
	}
	if app.GrpServer() != nil {
		RegisterGreeterServer(app.HttpServer(), srv)
	}
}
