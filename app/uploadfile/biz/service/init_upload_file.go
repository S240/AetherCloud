package service

import (
	"context"
	uploadfile "github.com/S240/AetherCloud/app/uploadfile/kitex_gen/uploadfile"
)

type InitUploadFileService struct {
	ctx context.Context
} // NewInitUploadFileService new InitUploadFileService
func NewInitUploadFileService(ctx context.Context) *InitUploadFileService {
	return &InitUploadFileService{ctx: ctx}
}

// Run create note info
func (s *InitUploadFileService) Run(req *uploadfile.InituploadFileRequest) (resp *uploadfile.InituploadFileResponse, err error) {
	// Finish your business logic.

	return
}
