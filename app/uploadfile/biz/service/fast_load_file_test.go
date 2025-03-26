package service

import (
	"context"
	"testing"
	uploadfile "github.com/S240/AetherCloud/app/uploadfile/kitex_gen/uploadfile"
)

func TestFastLoadFile_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFastLoadFileService(ctx)
	// init req and assert value

	req := &uploadfile.FastLoadRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
