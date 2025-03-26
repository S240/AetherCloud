package service

import (
	"context"
	uploadfile "github.com/S240/AetherCloud/app/uploadfile/kitex_gen/uploadfile"
)

type FastLoadFileService struct {
	ctx context.Context
} // NewFastLoadFileService new FastLoadFileService
func NewFastLoadFileService(ctx context.Context) *FastLoadFileService {
	return &FastLoadFileService{ctx: ctx}
}

// Run create note info
func (s *FastLoadFileService) Run(req *uploadfile.FastLoadRequest) (resp *uploadfile.FastLoadResponse, err error) {
	// Finish your business logic.

	return
}
