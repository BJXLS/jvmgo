package extended

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IFNULL struct{ base.BranchInstruction } // 根据引用是否是null进行跳转
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
