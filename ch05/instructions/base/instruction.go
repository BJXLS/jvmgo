package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码中提取操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

// NoOperandsInstruction 表示没有操作数的指令
type NoOperandsInstruction struct {
}

// BranchInstruction 表示跳转指令
type BranchInstruction struct {
	Offset int
}

// Index8Instruction 存储和加载指令，通过索引读取局部变量表
type Index8Instruction struct {
	Index uint // 局部变量表中的index
}

// Index16Instruction 抽象访问运行时常量池的操作
type Index16Instruction struct {
	Index uint
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 什么也不做
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
