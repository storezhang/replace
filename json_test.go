package replace

import (
	`testing`
)

func TestJSON(t *testing.T) {
	json := NewJSONReplace("", JSONReplaceElement{
		Path:  "",
		Value: "",
	})

	if err := json.Replace(""); nil != err {
		t.Error(err)
	}
}
