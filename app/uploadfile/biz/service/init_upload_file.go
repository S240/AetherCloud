package service

import (
	"context"
	"time"
	"fmt"
	"log"
	"net/url"
	"sync"
    "strconv"
	uploadfile "github.com/S240/AetherCloud/app/uploadfile/kitex_gen/uploadfile"
	// "github.com/S240/AetherCloud/app/uploadfile/biz/model"
	miniocl "github.com/S240/AetherCloud/app/uploadfile/biz/dal/minio"
	"github.com/S240/AetherCloud/app/uploadfile/conf"
	"github.com/minio/minio-go/v7"
)

type InitUploadFileService struct {
	ctx context.Context
} // NewInitUploadFileService new InitUploadFileService
func NewInitUploadFileService(ctx context.Context) *InitUploadFileService {
	return &InitUploadFileService{ctx: ctx}
}


var wg sync.WaitGroup
// Run create note info
func (s *InitUploadFileService) Run(req *uploadfile.InituploadFileRequest) (resp *uploadfile.InituploadFileResponse, err error) {
	// fileMeta := model.MetaFile{
	// 	FileName: req.FileName,
	// 	FileMd5: req.FileMd5, 
	// 	FileSize: req.FileSize,
	// 	FileState: model.FileStateUploading,
	// 	UploadAt: time.Now(),
	// }
	// 文件写入MinIO存储,进行初始化
	if miniocl.MinIOClt == nil {
		fmt.Println("错误: MinIO Core 未初始化")
		return &uploadfile.InituploadFileResponse{}, fmt.Errorf(" MinIO Core 未初始化")
	}
	opts := minio.PutObjectOptions{
        ContentType: "application/octet-stream",
    }
	exists, errBucketExists := miniocl.MinIOClt.BucketExists(context.Background(),conf.GetConf().MinioConfig.MinioBucketName)
	if errBucketExists == nil && exists {
		log.Printf("We already own %s\n", conf.GetConf().MinioConfig.MinioBucketName)
	} else {
		log.Println(errBucketExists) 
	}
	
	// 创建 multipart 上传任务，获取 uploadID
    uploadID, err := miniocl.MinIOClt.NewMultipartUpload(context.Background(), conf.GetConf().MinioConfig.MinioBucketName, req.FileMd5, opts)
    if err != nil {
        log.Fatalln("NewMultipartUpload error:", err)
		return &uploadfile.InituploadFileResponse{}, fmt.Errorf("NewMultipartUpload error:", err)
    }
    log.Println("uploadID:", uploadID)
	if req.ChunkSize == 0{
		log.Fatalln("ChunkSize为0")
		return &uploadfile.InituploadFileResponse{}, fmt.Errorf("ChunkSize为0")
	}
	totalParts := int((req.FileSize + req.ChunkSize - 1) / req.ChunkSize)
	presignPairs := make([]*uploadfile.PreSignMessage, totalParts)
	// 定义要上传的分片编号,从1开始
	bucketName:=conf.GetConf().MinioConfig.MinioBucketName
	objectName:=req.FileMd5
	for partNum:=1 ; partNum<=totalParts; partNum++{
		wg.Add(1)
		go func(pn int){
			defer wg.Done()
			params := url.Values{
				"uploadId":   []string{uploadID},
				"partNumber": []string{strconv.Itoa(pn)},
			}
			// 设置签名 URL 的有效期为 60 秒
			expiry := 60 * time.Second

			// 调用 Presign 方法生成签名 URL，HTTP 方法指定为 PUT
			presignedURL, err := miniocl.MinIOClt.Presign(context.Background(), "PUT", bucketName, objectName, expiry, params)
			if err != nil {
				log.Printf("生成分片%d预签名 URL 失败: %v\n", pn, err)
				return
			}
			presignPairs[pn-1] = &uploadfile.PreSignMessage{
				PartNumber: int32(pn),
				PresignUrl: presignedURL,
			}
		}(partNum)
	}
	wg.Wait()

	fmt.Println(presignPairs)
	//更新文件表记录
	// tx := mysql.DB.Begin()
	// res := tx.Save(&fileMeta)
	// if res.Error  != nil {
	// 	tx.Rollback()
	// 	return &uploadfile.InituploadFileResponse{}, res.Error 
	// }	
	// tx.Commit()
	return &uploadfile.InituploadFileResponse{}, nil
}
