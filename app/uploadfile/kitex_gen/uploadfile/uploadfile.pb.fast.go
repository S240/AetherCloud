// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package uploadfile

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *FastLoadRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FastLoadRequest[number], err)
}

func (x *FastLoadRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileMd5, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FastLoadRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FileName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FastLoadRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FileSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FastLoadRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.IsChunked, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *FastLoadRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.ChunkSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FastLoadResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FastLoadResponse[number], err)
}

func (x *FastLoadResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsFastload, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *PreSignMessage) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PreSignMessage[number], err)
}

func (x *PreSignMessage) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PresignUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PreSignMessage) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PartNumber, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PartsMessage) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PartsMessage[number], err)
}

func (x *PartsMessage) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PartNumber, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PartsMessage) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Etag, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *InituploadFileRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InituploadFileRequest[number], err)
}

func (x *InituploadFileRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileMd5, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *InituploadFileRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FileName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *InituploadFileRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FileSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *InituploadFileRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.IsChunked, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *InituploadFileRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.ChunkSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *InituploadFileResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InituploadFileResponse[number], err)
}

func (x *InituploadFileResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v PartsMessage
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Parts = append(x.Parts, &v)
	return offset, nil
}

func (x *InituploadFileResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FileMd5, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FastLoadRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *FastLoadRequest) fastWriteField1(buf []byte) (offset int) {
	if x.FileMd5 == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileMd5())
	return offset
}

func (x *FastLoadRequest) fastWriteField2(buf []byte) (offset int) {
	if x.FileName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFileName())
	return offset
}

func (x *FastLoadRequest) fastWriteField3(buf []byte) (offset int) {
	if x.FileSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFileSize())
	return offset
}

func (x *FastLoadRequest) fastWriteField4(buf []byte) (offset int) {
	if !x.IsChunked {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetIsChunked())
	return offset
}

func (x *FastLoadRequest) fastWriteField5(buf []byte) (offset int) {
	if x.ChunkSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetChunkSize())
	return offset
}

func (x *FastLoadResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *FastLoadResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.IsFastload {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsFastload())
	return offset
}

func (x *PreSignMessage) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PreSignMessage) fastWriteField1(buf []byte) (offset int) {
	if x.PresignUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPresignUrl())
	return offset
}

func (x *PreSignMessage) fastWriteField2(buf []byte) (offset int) {
	if x.PartNumber == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetPartNumber())
	return offset
}

func (x *PartsMessage) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PartsMessage) fastWriteField1(buf []byte) (offset int) {
	if x.PartNumber == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetPartNumber())
	return offset
}

func (x *PartsMessage) fastWriteField2(buf []byte) (offset int) {
	if x.Etag == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEtag())
	return offset
}

func (x *InituploadFileRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *InituploadFileRequest) fastWriteField1(buf []byte) (offset int) {
	if x.FileMd5 == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileMd5())
	return offset
}

func (x *InituploadFileRequest) fastWriteField2(buf []byte) (offset int) {
	if x.FileName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFileName())
	return offset
}

func (x *InituploadFileRequest) fastWriteField3(buf []byte) (offset int) {
	if x.FileSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFileSize())
	return offset
}

func (x *InituploadFileRequest) fastWriteField4(buf []byte) (offset int) {
	if !x.IsChunked {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetIsChunked())
	return offset
}

func (x *InituploadFileRequest) fastWriteField5(buf []byte) (offset int) {
	if x.ChunkSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetChunkSize())
	return offset
}

func (x *InituploadFileResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *InituploadFileResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Parts == nil {
		return offset
	}
	for i := range x.GetParts() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetParts()[i])
	}
	return offset
}

func (x *InituploadFileResponse) fastWriteField2(buf []byte) (offset int) {
	if x.FileMd5 == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFileMd5())
	return offset
}

func (x *FastLoadRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *FastLoadRequest) sizeField1() (n int) {
	if x.FileMd5 == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileMd5())
	return n
}

func (x *FastLoadRequest) sizeField2() (n int) {
	if x.FileName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFileName())
	return n
}

func (x *FastLoadRequest) sizeField3() (n int) {
	if x.FileSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFileSize())
	return n
}

func (x *FastLoadRequest) sizeField4() (n int) {
	if !x.IsChunked {
		return n
	}
	n += fastpb.SizeBool(4, x.GetIsChunked())
	return n
}

func (x *FastLoadRequest) sizeField5() (n int) {
	if x.ChunkSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetChunkSize())
	return n
}

func (x *FastLoadResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *FastLoadResponse) sizeField1() (n int) {
	if !x.IsFastload {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsFastload())
	return n
}

func (x *PreSignMessage) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PreSignMessage) sizeField1() (n int) {
	if x.PresignUrl == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPresignUrl())
	return n
}

func (x *PreSignMessage) sizeField2() (n int) {
	if x.PartNumber == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetPartNumber())
	return n
}

func (x *PartsMessage) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PartsMessage) sizeField1() (n int) {
	if x.PartNumber == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetPartNumber())
	return n
}

func (x *PartsMessage) sizeField2() (n int) {
	if x.Etag == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEtag())
	return n
}

func (x *InituploadFileRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *InituploadFileRequest) sizeField1() (n int) {
	if x.FileMd5 == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileMd5())
	return n
}

func (x *InituploadFileRequest) sizeField2() (n int) {
	if x.FileName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFileName())
	return n
}

func (x *InituploadFileRequest) sizeField3() (n int) {
	if x.FileSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFileSize())
	return n
}

func (x *InituploadFileRequest) sizeField4() (n int) {
	if !x.IsChunked {
		return n
	}
	n += fastpb.SizeBool(4, x.GetIsChunked())
	return n
}

func (x *InituploadFileRequest) sizeField5() (n int) {
	if x.ChunkSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetChunkSize())
	return n
}

func (x *InituploadFileResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *InituploadFileResponse) sizeField1() (n int) {
	if x.Parts == nil {
		return n
	}
	for i := range x.GetParts() {
		n += fastpb.SizeMessage(1, x.GetParts()[i])
	}
	return n
}

func (x *InituploadFileResponse) sizeField2() (n int) {
	if x.FileMd5 == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFileMd5())
	return n
}

var fieldIDToName_FastLoadRequest = map[int32]string{
	1: "FileMd5",
	2: "FileName",
	3: "FileSize",
	4: "IsChunked",
	5: "ChunkSize",
}

var fieldIDToName_FastLoadResponse = map[int32]string{
	1: "IsFastload",
}

var fieldIDToName_PreSignMessage = map[int32]string{
	1: "PresignUrl",
	2: "PartNumber",
}

var fieldIDToName_PartsMessage = map[int32]string{
	1: "PartNumber",
	2: "Etag",
}

var fieldIDToName_InituploadFileRequest = map[int32]string{
	1: "FileMd5",
	2: "FileName",
	3: "FileSize",
	4: "IsChunked",
	5: "ChunkSize",
}

var fieldIDToName_InituploadFileResponse = map[int32]string{
	1: "Parts",
	2: "FileMd5",
}
