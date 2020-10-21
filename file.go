package replace

import (
	`encoding/json`

	`github.com/storezhang/transfer`
)

// 文件替换
type FileReplace struct {
	// 文件
	File transfer.File `json:"file" validate:"required"`
}

// NewFileReplace 文件替换
func NewFileReplace(filename string, file transfer.File) Replace {
	return NewReplace(filename, TypeFile, FileReplace{File: file})
}

func (fr FileReplace) Replace(filename string) (err error) {
	if "" == filename {
		return
	}

	return fr.File.Download(filename, true)
}

func (fr FileReplace) String() string {
	jsonBytes, _ := json.MarshalIndent(fr, "", "    ")

	return string(jsonBytes)
}
