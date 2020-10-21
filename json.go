package replace

import (
	`encoding/json`
	`io/ioutil`
	`os`

	`github.com/tidwall/sjson`
)

type (
	// JSONReplaceElement JSON替换元素
	JSONReplaceElement struct {
		// 节点路径
		Path string `json:"path" validate:"required"`
		// 要设置的值
		Value string `json:"value" validate:"required"`
	}

	// JSONReplace JSON替换
	JSONReplace struct {
		// 修改元素列表
		Elements []JSONReplaceElement `json:"elements" validate:"required,dive"`
	}
)

// NewFileReplace JSON文件修改
func NewJSONReplace(filename string, elements ...JSONReplaceElement) Replace {
	return NewReplace(filename, TypeJSON, JSONReplace{Elements: elements})
}

func (jr JSONReplace) Replace(filename string) (err error) {
	if "" == filename {
		return
	}

	var fileContent []byte
	if fileContent, err = ioutil.ReadFile(filename); nil != err {
		return
	}

	for _, element := range jr.Elements {
		if fileContent, err = sjson.SetBytes(fileContent, element.Path, element.Value); nil != err {
			return
		}
	}

	err = ioutil.WriteFile(filename, fileContent, os.ModePerm)

	return
}

func (jr JSONReplace) String() string {
	jsonBytes, _ := json.MarshalIndent(jr, "", "    ")

	return string(jsonBytes)
}
