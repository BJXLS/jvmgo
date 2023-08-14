package rtda

import "jvmgo/ch06/rtda/heap"

type Slot struct {
	num int32        // 存放数组
	ref *heap.Object // 存放引用
}
