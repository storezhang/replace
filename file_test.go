package replace

import (
	`testing`

	`github.com/storezhang/transfer`
)

func TestFile(t *testing.T) {
	file := NewFileReplace("", transfer.NewLocalFile(""))

	if err := file.Replace(""); nil != err {
		t.Error(err)
	}
}
