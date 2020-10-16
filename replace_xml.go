package replace

import (
	`encoding/json`

	`github.com/beevik/etree`
)

const (
	// 替换属性
	XMLReplaceTypeAttr XMLReplaceType = "attr"
	// 替换文本
	XMLReplaceTypeText XMLReplaceType = "text"
)

type (
	// XMLReplaceType XML的操作类型
	XMLReplaceType string

	// XMLReplaceElement XML文件替换元素
	XMLReplaceElement struct {
		// XPath 节点路径
		XPath string `json:"xpath" validate:"required"`
		// Type XML替换类型
		Type XMLReplaceType `json:"type" validate:"required,oneof=attr text"`
		// Key 操作的键
		Key string `json:"key" validate:"omitempty"`
		// Value 要设置的值
		Value string `json:"value" validate:"required"`
	}

	// XMLReplace XML文件替换
	XMLReplace struct {
		// Elements 需要替换的元素列表
		Elements []XMLReplaceElement `json:"elements" validate:"required,dive"`
	}
)

func (xr XMLReplace) String() string {
	jsonBytes, _ := json.MarshalIndent(xr, "", "    ")

	return string(jsonBytes)
}

func (xre XMLReplaceElement) String() string {
	jsonBytes, _ := json.MarshalIndent(xre, "", "    ")

	return string(jsonBytes)
}

func (xr *XMLReplace) Replace(filename string) (err error) {
	xml := etree.NewDocument()
	if err = xml.ReadFromFile(filename); nil != err {
		return
	}

	for _, element := range xr.Elements {
		xmlElements := xml.FindElements(element.XPath)

		switch element.Type {
		case XMLReplaceTypeText:
			for _, e := range xmlElements {
				e.SetText(element.Value)
			}
		case XMLReplaceTypeAttr:
			for _, e := range xmlElements {
				e.RemoveAttr(element.Key)
				e.CreateAttr(element.Key, element.Value)
			}
		}
	}

	err = xml.WriteToFile(filename)

	return
}
