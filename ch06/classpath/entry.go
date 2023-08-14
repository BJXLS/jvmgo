package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 负责寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// String 相当于toString
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// TODO 这个功能感觉效率很低下啊，记得优化一下
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	// 直接通过路径进行读取
	return newDirEntry(path)
}
