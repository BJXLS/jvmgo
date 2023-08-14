package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()                    // tableswitch指令操作码后面有0~3个字节的padding，以保证defaultOffset字节码的地址是4的倍数
	self.defaultOffset = reader.ReadInt32() // 默认情况下执行跳转所需的字节码便宜量
	self.low = reader.ReadInt32()           // low和high记录case的取值范围
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount) // 这是一个索引表，里面对应各种case情况下，执行跳转所需的字节码偏移量
}

// Execute 先判断变量是否在范围内，如果在，就进行分情况跳转。
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
