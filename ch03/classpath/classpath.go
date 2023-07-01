package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// Parse 使用jreOption解析启动类路径和扩展类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// 这里用启动和扩展类去加载jre（也就是这一步包含了前面的两个路径的读取）
	cp.parseBootAndExtClasspath(jreOption)
	// 用用户启动
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	fmt.Printf("Jre Dir: %v\n", jreDir)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 如果用户没有提供-cp选项，就用当前目录作为用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	// 优先使用用户输入的-Xjre选项作为jre目录。
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果没有输入该选项，则在当前目录寻找jre目录
	if exists("./jre") {
		return "./jre"
	}
	// 如果还没有就尝试读取JAVA_HOME
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// TODO 对三个部分依次读取（所以这里代表的是三类类加载器，但是这三类怎么双亲委派机制，怎么进行分开读取的？）
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
