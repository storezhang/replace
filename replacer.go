package replace

// Replacer 替换
type Replacer interface {
	// Replace 执行替换
	Replace(filename string) (err error)
}
