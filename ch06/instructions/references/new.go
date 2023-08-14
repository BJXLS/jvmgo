package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

/*
用来创建类实例的
*/
type NEW struct {
	base.Index16Instruction // new指令的操作数是一个uint16索引，来自字节码
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 拿到类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类符号引用
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	// 创建类对象
	ref := class.NewObject()
	// 放入操作数栈
	frame.OperandStack().PushRef(ref)
}
