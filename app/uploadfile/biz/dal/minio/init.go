package minio

import (
	"log"
	"github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/S240/AetherCloud/app/uploadfile/conf"
)

var MinIOClt *minio.Core
func Init() {
	var err error
	MinIOClt, err = minio.NewCore(conf.GetConf().MinioConfig.MinioEndPoint, &minio.Options{
        Creds:  credentials.NewStaticV4(conf.GetConf().MinioConfig.MinioAccessKey, conf.GetConf().MinioConfig.MinioSecretKey, ""),
        Secure: false,
    })
	if err != nil {
		log.Println("new minio Core fail: ", err)
		panic(err)
	}
	log.Printf("%#v\n", MinIOClt) 
}
