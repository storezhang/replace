package replace

import (
	libJSON `encoding/json`
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

// User 测试用户
type User struct {
	// Username 用户名
	Username string `json:"username"`
	// Password 密码
	Password string `json:"password"`
	// School 学校
	School struct {
		// Name 名称
		Name string `json:"name"`
		// Address 地址
		Address string `json:"address"`
	}
}

func TestJSON(t *testing.T) {
	if err := ioutil.WriteFile(userJSONFilename, []byte(userJSON), os.ModePerm); nil != err {
		t.Error(err)
	}

	json := NewJSONReplace(userJSONFilename, JSONReplaceElement{
		Path:  "username",
		Value: expectedUsername,
	}, JSONReplaceElement{
		Path:  "school.name",
		Value: expectedSchoolName,
	})

	if err := json.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(userJSONFilename); nil != err {
		t.Error(err)
	} else {
		user := &User{}
		if err = libJSON.Unmarshal(fileData, user); nil != err {
			t.Error(errors.New("文件修改失败"))
		}

		if expectedUsername != user.Username || expectedSchoolName != user.School.Name {
			t.Error(errors.New("文件修改失败"))
		}
	}
}
