package replace

import (
	`github.com/storezhang/gox`
)

var (
	// ErrorNotSupportReplace 不支持的文件替换
	ErrorNotSupportReplace = &gox.CodeError{ErrorCode: 103, Message: "不支持的文件替换"}
)
