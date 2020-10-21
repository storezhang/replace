package replace

import (
	`encoding/json`
	`io/ioutil`
	`os`
	`strings`
)

// StringContentReplace 文本内容替换
type StringContentReplace struct {
	// Old 旧内容
	Old string `json:"old" validate:"required"`
	// New 新内容
	New string `json:"new" validate:"required"`
}

// NewStringContentReplace 字符串替换
func NewStringContentReplace(filename string, old string, new string) Replace {
	return NewReplace(filename, TypeStringContent, StringContentReplace{
		Old: old,
		New: new,
	})
}

func (scr StringContentReplace) Replace(filename string) (err error) {
	if "" == filename {
		return
	}

	var fileContent []byte
	if fileContent, err = ioutil.ReadFile(filename); nil != err {
		return
	}

	newContent := strings.ReplaceAll(string(fileContent), scr.Old, scr.New)
	err = ioutil.WriteFile(filename, []byte(newContent), os.ModePerm)

	return
}

func (scr StringContentReplace) String() string {
	jsonBytes, _ := json.MarshalIndent(scr, "", "    ")

	return string(jsonBytes)
}
