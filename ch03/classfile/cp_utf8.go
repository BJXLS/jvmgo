package classfile

import (
	"bytes"
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	str string
}

// 先读取出[]byte，再解码成Go字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
