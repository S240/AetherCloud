package service

import (
	"context"

	apigw "github.com/S240/AetherCloud/app/apigw/hertz_gen/apigw"
	"github.com/cloudwego/hertz/pkg/app"
)

type CompleteUploadService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCompleteUploadService(Context context.Context, RequestContext *app.RequestContext) *CompleteUploadService {
	return &CompleteUploadService{RequestContext: RequestContext, Context: Context}
}

func (h *CompleteUploadService) Run(req *apigw.CompletedUploadRequest) (resp *apigw.StandardResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
