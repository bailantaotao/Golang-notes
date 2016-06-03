package pool

/*
用來保存和檢索特定的物件，可以減少記憶體的分配和ＧＣ的壓力
GET會返回Pool中任意的物件．
如果Pool為空，則在呼叫Get時，會取得New function的內容

資料Put進去時，資料是有可能消失的，也就是被ＧＣ了
好處就是Pool不會一直增長，因為都被GO GC了
每次清理都會將裡面所有資料都清空
*/

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	type book struct {
		Name   string
		Author string
	}

	var pipe = &sync.Pool{New: func() interface{} { return &book{Name: "Hello world", Author: "Edwin"} }}

	b := &book{
		Name:   "Sync pool pattern example",
		Author: "Edwin",
	}
	pipe.Put(b)

	// output: &{Sync pool pattern example Edwin}
	t.Log(pipe.Get())
	// output: &{Hello world Edwin}
	t.Log(pipe.Get())
}
