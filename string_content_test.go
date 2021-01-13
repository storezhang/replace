package replace

import (
	"errors"
	"io/ioutil"
	"os"
	`strings`
	"testing"
)

const (
	stringContentFilename = ".testdata/string_content.xml"
)

func TestStringContent(t *testing.T) {
	if err := ioutil.WriteFile(androidXMLFilename, []byte(androidXML), os.ModePerm); nil != err {
		t.Error(err)
	}

	replace := NewStringContentReplace(stringContentFilename, "com.class100.yunke.dev", "com.class100.yunke")

	if err := replace.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(androidXMLFilename); nil != err {
		t.Error(err)
	} else {
		if strings.Contains(string(fileData), "com.class100.yunke.dev") {
			t.Error(errors.New("文件修改失败"))
		}
	}
}
