package rtda

type Frame struct {
	lower        *Frame        // 用来实现链表数据结构（代表更低一级的栈，可以类比链表的next）
	localVars    LocalVars     // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
