package main

import (
	"context"
	uploadfile "github.com/S240/AetherCloud/app/uploadfile/kitex_gen/uploadfile"
	"github.com/S240/AetherCloud/app/uploadfile/biz/service"
)

// UploadFileServiceImpl implements the last service interface defined in the IDL.
type UploadFileServiceImpl struct{}

// FastLoadFile implements the UploadFileServiceImpl interface.
func (s *UploadFileServiceImpl) FastLoadFile(ctx context.Context, req *uploadfile.FastLoadRequest) (resp *uploadfile.FastLoadResponse, err error) {
	resp, err = service.NewFastLoadFileService(ctx).Run(req)

	return resp, err
}

// InitUploadFile implements the UploadFileServiceImpl interface.
func (s *UploadFileServiceImpl) InitUploadFile(ctx context.Context, req *uploadfile.InituploadFileRequest) (resp *uploadfile.InituploadFileResponse, err error) {
	resp, err = service.NewInitUploadFileService(ctx).Run(req)

	return resp, err
}
