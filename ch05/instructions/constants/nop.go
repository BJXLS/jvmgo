package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
