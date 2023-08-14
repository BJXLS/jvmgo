package heap

import "jvmgo/ch06/classfile"

// Field 结构和ClassFile相同
type Field struct {
	ClassMember          // 这里是对内容进行了一个抽象
	slotId          uint // 在数组中的位置
	constValueIndex uint // 常量数的index，用于初始化赋值
}

// 从classFile复制fields
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField) // 常量
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}
