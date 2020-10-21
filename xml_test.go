package replace

import (
	`testing`
)

func TestXML(t *testing.T) {
	xml := NewXMLReplace("", XMLReplaceElement{
		XPath: "",
		Type:  XMLReplaceTypeAttr,
		Key:   "",
		Value: "",
	})

	if err := xml.Replace(""); nil != err {
		t.Error(err)
	}
}
