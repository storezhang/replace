package replace

import (
	`encoding/json`
	`io/ioutil`
	`os`
)

// FileWriteReplace 文件替换
type FileWriteReplace struct {
	// 内容
	Content string `json:"content" validate:"required"`
}

// NewFileWriteReplace 文件内容替换
func NewFileWriteReplace(fileName string, content string) Replace {
	return NewReplace(fileName, TypeFileWrite, FileWriteReplace{Content: content})
}

func (fwr *FileWriteReplace) Replace(filename string) (err error) {
	return ioutil.WriteFile(filename, []byte(fwr.Content), os.ModePerm)
}

func (fwr FileWriteReplace) String() string {
	jsonBytes, _ := json.MarshalIndent(fwr, "", "    ")

	return string(jsonBytes)
}
