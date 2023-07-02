package rtda

type Frame struct {
	lower        *Frame    // 用来实现链表数据结构（代表更低一级的栈，可以类比链表的next）
	localVars    LocalVars // 保存局部变量表指针
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
