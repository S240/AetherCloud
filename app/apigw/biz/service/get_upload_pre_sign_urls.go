package service

import (
	"context"
	"fmt"
	uploadfile "github.com/S240/AetherCloud/client/uploadfile/kitex_gen/uploadfile"
	apigw "github.com/S240/AetherCloud/app/apigw/hertz_gen/apigw"
	UploadFileClient "github.com/S240/AetherCloud/client/uploadfile/rpc/uploadfile"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUploadPreSignUrlsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUploadPreSignUrlsService(Context context.Context, RequestContext *app.RequestContext) *GetUploadPreSignUrlsService {
	return &GetUploadPreSignUrlsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUploadPreSignUrlsService) Run(req *apigw.UploadFileRequest) (resp *apigw.StandardResponse, err error) {
	resp = &apigw.StandardResponse{
		Code:4000,
		Message:"获取失败",
		Data: nil,
	}
	// 参数验证
	if len(req.FileMd5) == 0 {
		resp.Message = "MD5为空"
		return resp, nil
	}
	if len(req.FileName) == 0 {
		resp.Message = "文件名为空"
		return resp, nil
	}
	if req.FileSize<0 || req.ChunkSize<0 {
		resp.Message = "文件大小错误"
		return resp, nil
	}
	//TODO 是否fastload
	
	//

	// 初始化传输信息
	inituploadReq := &uploadfile.InituploadFileRequest{
		FileMd5: req.FileMd5,
		FileName: req.FileName,
		FileSize: req.FileSize,
		IsChunked: req.IsChunked,
		ChunkSize: req.ChunkSize,
	}
	inituploadResp, err := UploadFileClient.InitUploadFile(h.Context, inituploadReq)
	fmt.Println(inituploadResp)

	if err != nil {
		resp.Message = "文件上传服务暂时不可用"
		return resp, err
	}
	return resp,nil
}
