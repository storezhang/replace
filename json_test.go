package replace

import (
	`errors`
	`io/ioutil`
	`os`
	`testing`
)

const (
	userJSONFilename = ".testdata/user.json"
	userJSON         = `
{
  "username": "storezhang",
  "password": "123456",
  "school": {
    "name": "电子科技大学",
    "address": "中国 四川 成都"
  }
}
`
	expectedUserJSON = `
{
  "username": "storezhang@gmail.com",
  "password": "123456",
  "school": {
    "name": "中国电子科技大学",
    "address": "中国 四川 成都"
  }
}
`
)

func TestJSON(t *testing.T) {
	if err := ioutil.WriteFile(userJSONFilename, []byte(userJSON), os.ModePerm); nil != err {
		t.Error(err)
	}

	json := NewJSONReplace(userJSONFilename, JSONReplaceElement{
		Path:  "username",
		Value: "storezhang@gmail.com",
	}, JSONReplaceElement{
		Path:  "school.name",
		Value: "中国电子科技大学",
	})

	if err := json.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(userJSONFilename); nil != err {
		t.Error(err)
	} else if expectedUserJSON != string(fileData) {
		t.Error(errors.New("文件修改失败"))
	}
}
