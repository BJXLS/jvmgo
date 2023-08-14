package heap

import "strings"

// SymRef symbolic reference
// 引用符号的父类
type SymRef struct {
	cp        *ConstantPool // 符号引用，所在运行时常量池指针
	className string        // 类的完全限定名
	class     *Class        // 缓存解析后的类结构体指针，这样类符号引用只需要解析一次
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil { // 如果没有被解析，则开始解析
		self.resolveClassRef()
	}
	return self.class // 直接返回类指针，缓存
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName() // 想要访问要满足其一：被访问者为public，或者两者在同一包下
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
