package replace

import (
	`encoding/json`
	`errors`
	`io/ioutil`
	`os`
	`testing`
)

const (
	userJSONFilename   = ".testdata/user.json"
	expectedUsername   = "storezhang@gmail.com"
	expectedSchoolName = "中国电子科技大学"
	userJSON           = `
{
  "username": "storezhang",
  "password": "123456",
  "school": {
    "name": "电子科技大学",
    "address": "中国 四川 成都"
  }
}
`
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	School struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}
}

func TestJSON(t *testing.T) {
	if err := ioutil.WriteFile(userJSONFilename, []byte(userJSON), os.ModePerm); nil != err {
		t.Error(err)
	}

	replace := NewJSONReplace(userJSONFilename, JSONReplaceElement{
		Path:  "username",
		Value: expectedUsername,
	}, JSONReplaceElement{
		Path:  "school.name",
		Value: expectedSchoolName,
	})

	if err := replace.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(userJSONFilename); nil != err {
		t.Error(err)
	} else {
		user := &User{}
		if err = json.Unmarshal(fileData, user); nil != err {
			t.Error(errors.New("文件修改失败"))
		}

		if expectedUsername != user.Username || expectedSchoolName != user.School.Name {
			t.Error(errors.New("文件修改失败"))
		}
	}
}
