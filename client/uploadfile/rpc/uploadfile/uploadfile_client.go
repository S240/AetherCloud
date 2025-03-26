package uploadfile

import (
	"context"
	uploadfile "github.com/S240/AetherCloud/client/uploadfile/kitex_gen/uploadfile"

	"github.com/S240/AetherCloud/client/uploadfile/kitex_gen/uploadfile/uploadfileservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() uploadfileservice.Client
	Service() string
	FastLoadFile(ctx context.Context, Req *uploadfile.FastLoadRequest, callOptions ...callopt.Option) (r *uploadfile.FastLoadResponse, err error)
	InitUploadFile(ctx context.Context, Req *uploadfile.InituploadFileRequest, callOptions ...callopt.Option) (r *uploadfile.InituploadFileResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := uploadfileservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient uploadfileservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() uploadfileservice.Client {
	return c.kitexClient
}

func (c *clientImpl) FastLoadFile(ctx context.Context, Req *uploadfile.FastLoadRequest, callOptions ...callopt.Option) (r *uploadfile.FastLoadResponse, err error) {
	return c.kitexClient.FastLoadFile(ctx, Req, callOptions...)
}

func (c *clientImpl) InitUploadFile(ctx context.Context, Req *uploadfile.InituploadFileRequest, callOptions ...callopt.Option) (r *uploadfile.InituploadFileResponse, err error) {
	return c.kitexClient.InitUploadFile(ctx, Req, callOptions...)
}
