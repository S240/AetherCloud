package uploadfile

import (
	"context"
	uploadfile "github.com/S240/AetherCloud/client/uploadfile/kitex_gen/uploadfile"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FastLoadFile(ctx context.Context, req *uploadfile.FastLoadRequest, callOptions ...callopt.Option) (resp *uploadfile.FastLoadResponse, err error) {
	resp, err = defaultClient.FastLoadFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FastLoadFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func InitUploadFile(ctx context.Context, req *uploadfile.InituploadFileRequest, callOptions ...callopt.Option) (resp *uploadfile.InituploadFileResponse, err error) {
	resp, err = defaultClient.InitUploadFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "InitUploadFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
