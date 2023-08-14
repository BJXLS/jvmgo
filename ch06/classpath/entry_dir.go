package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

// 统一使用new开头的函数构造结构体，并且称此为构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// 组合path，并且读入内容
func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, e, err
}

func (e *DirEntry) String() string {
	return e.absDir
}
