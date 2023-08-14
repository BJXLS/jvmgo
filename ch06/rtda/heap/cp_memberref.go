package heap

import "jvmgo/ch06/classfile"

// MemberRef 存放字段和方法符号引用的共有信息
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

// 从class文件内 存储的字段或方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
