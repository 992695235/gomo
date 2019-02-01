package dstr

import(
	"testing"
)

func TestSingleList_Init(t *testing.T)  {
	list := new(SingleList)
	list.Init()
	t.Log("single list init success")
}