package replace

import (
	`encoding/json`
	`path/filepath`
)

const (
	// 所有支持的文件替换类型
	// 文件内容替换
	TypeFileWrite Type = "write"
	// 字符串替换
	TypeStringContent Type = "string"
	// 文件替换
	TypeFile Type = "file"
	// JSON文件替换
	TypeJSON Type = "json"
	// XML文件替换
	TypeXML Type = "xml"
)

type (
	// ReplaceType 替换类型
	Type string

	// Replace 替换
	Replace struct {
		// Filename 文件名，使用相对路径
		Filename string `json:"filename" validate:"required"`
		// Type 类型
		Type Type `default:"file" json:"type" validate:"omitempty,oneof=file string write json xml"`
		// Value 需要修改的值
		Value interface{} `json:"value" validate:"required"`
	}
)

// NewReplace 创建新的替换
func NewReplace(filename string, replaceType Type, value interface{}) Replace {
	return Replace{
		Filename: filename,
		Type:     replaceType,
		Value:    value,
	}
}

func (r *Replace) UnmarshalJSON(data []byte) (err error) {
	type cloneType Replace

	rawMsg := json.RawMessage{}
	r.Value = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(r)); nil != err {
		return
	}

	switch r.Type {
	case TypeFileWrite:
		fwr := FileWriteReplace{}
		if err = json.Unmarshal(rawMsg, &fwr); nil != err {
			return
		}
		r.Value = fwr
	case TypeStringContent:
		cr := StringContentReplace{}
		if err = json.Unmarshal(rawMsg, &cr); nil != err {
			return
		}
		r.Value = cr
	case TypeFile:
		fr := FileReplace{}
		if err = json.Unmarshal(rawMsg, &fr); nil != err {
			return
		}
		r.Value = fr
	case TypeJSON:
		jr := JSONReplace{}
		if err = json.Unmarshal(rawMsg, &jr); nil != err {
			return
		}
		r.Value = jr
	case TypeXML:
		xr := XMLReplace{}
		if err = json.Unmarshal(rawMsg, &xr); nil != err {
			return
		}
		r.Value = xr
	default:
		err = ErrorNotSupportReplace
	}

	return
}

func (r Replace) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}

func (r *Replace) Replace(dir string) (err error) {
	return r.Value.(Replacer).Replace(filepath.Join(dir, r.Filename))
}
