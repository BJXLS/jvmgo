package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// GOTO_W goto_w和goto的唯一区别就是索引从2字节变成4字节了
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
