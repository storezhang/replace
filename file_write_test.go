package replace

import (
	`testing`
)

func TestFileWrite(t *testing.T) {
	file := NewFileWriteReplace("", "")

	if err := file.Replace(""); nil != err {
		t.Error(err)
	}
}
