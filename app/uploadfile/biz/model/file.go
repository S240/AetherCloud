package model
import "time"
type FileState int

const (
    FileStateUnknown FileState = iota // 未知状态
    FileStateUploading                // 上传中
    FileStateUploaded                 // 上传完成
    FileStateFailed                   // 上传失败
)


type MetaFile struct {
	Base

    FileName     string    `gorm:"size:255;not null" json:"file_name"`               // 文件名
    FileMd5     string    `gorm:"type:char(40);uniqueIndex;not null" json:"file_sha1"` // 文件哈希，用于秒传
    FileSize     int64     `gorm:"not null" json:"file_size"`                        // 文件大小（字节）
    FileState   FileState                         
    UploadAt     time.Time `gorm:"autoCreateTime" json:"upload_at"`
}