package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

// ClassLoader 类加载器
type ClassLoader struct {
	cp       *classpath.Classpath // 依赖cp来读取class文件，co保存Classpath指针
	classMap map[string]*Class    // 记录已经加载的类数据，key是类的完全限定名；classMap可以看做是方法区的具体实现
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class // 类已经加载
	}
	return self.loadNonArrayClass(name) // 数组类要区分开
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 1. 先找到class文件，读取数据到内存
	// 2. 解析class文件生成类数据，放入方法区
	// 3. 最后进行链接
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// 寻找class文件，并且读入内存
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

// 解析超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		// 可以实现双亲委派机制，递归调用父类的加载接口
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 类的链接分为验证和准备两个必要的阶段
func link(class *Class) {
	/*
		验证是连接阶段的第一步
		这一阶段的目的是确保 Class 文件的字节流中包含的信息符合《Java 虚拟机规范》的全部约束要求
		保证这些信息被当作代码运行后不会危害虚拟机自身的安全。
	*/
	verify(class)
	/*
		准备阶段是正式为类变量分配内存并设置类变量初始值的阶段
		这些内存都将在方法区中分配。
	*/
	prepare(class)
}

func verify(class *Class) {
	// TODO
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 计算实例字段的个数，同时给他们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0) // 从0开始
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount // 从超类的后面开始
	}
	for _, field := range class.fields {
		if !field.IsStatic() { // 不是类字段就编号++
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() { // long和double占两个字符
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 计算静态字段个数
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

// 给列变量分配空间
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 将static值提前写入
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
