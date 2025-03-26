package apigw

import (
	"context"

	"github.com/S240/AetherCloud/app/apigw/biz/service"
	"github.com/S240/AetherCloud/app/apigw/biz/utils"
	apigw "github.com/S240/AetherCloud/app/apigw/hertz_gen/apigw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUploadPreSignUrls .
// @router /file/getuploadurl [POST]
func GetUploadPreSignUrls(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apigw.UploadFileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetUploadPreSignUrlsService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CompleteUpload .
// @router /file/completeUpload [POST]
func CompleteUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apigw.CompletedUploadRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCompleteUploadService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
