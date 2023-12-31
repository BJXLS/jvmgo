package rtda

type Thread struct {
	pc    int    // TODO what is pc?
	stack *Stack // 是虚拟机栈指针
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), // 创建stack结构体实例，参数表示要创建的stack最多可以容纳多少帧
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}
