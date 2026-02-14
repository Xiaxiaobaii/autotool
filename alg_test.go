package autotool

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewArrayStack[int](3)
	stack.Push(5)
	if stack.MaxSize() != 3 {
		t.Error("stack max size != 3!")
		t.FailNow()
	} else if stack.Size() != 1 {
		t.Error("stack size != 1!")
		t.FailNow()
	}
	stack.Push(4)
	if stack.Size() != 2 {
		t.Error("stack size != 2!")
		t.FailNow()
	} else if stack.Peek() != 4 {
		t.Error("stack peek error!")
		t.FailNow()
	} else if fmt.Sprintf("%v", stack.ToSlice()) != "[5 4 0]" {
		t.Error("stack data error!")
		t.FailNow()
	}
	stack.Pop()
	if stack.Size() != 1 {
		t.Error("stack size != 1!")
		t.FailNow()
	} else if stack.Peek() != 5 {
		t.Error("stack peek error!")
		t.FailNow()
	} else if fmt.Sprintf("%v", stack.ToSlice()) != "[5 0 0]" {
		t.Error("stack data error!")
		t.FailNow()
	}
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	if stack.Size() != 3 {
		t.Error("stack size != 1!")
		t.FailNow()
	} else if stack.Peek() != 5 {
		t.Error("stack peek error!")
		t.FailNow()
	} else if fmt.Sprintf("%v", stack.ToSlice()) != "[5 4 5]" {
		t.Error("stack data error!")
		t.FailNow()
	}
}

func TestQuene(t *testing.T) {
	data := NewArrayQuene[int](3)
	if data.MaxSize() != 3 {
		t.Error("quene max size != 3")
	} else if data.Size() != 0 {
		t.Error("quene size != 0")
	}
	data.Push(4)
	data.Push(5)
	if data.Size() != 2 {
		t.Error("quene size != 2")
	} else if fmt.Sprintf("%v", data.ToSlice()) != "[4 5 0]" {
		t.Errorf("quene data error: %v", data.data)
	}
	data.Pop()
	if data.Size() != 1 {
		t.Error("quene size != 1")
	} else if fmt.Sprintf("%v", data.ToSlice()) != "[5 0 0]" {
		t.Error("quene data error")
	}else if data.Peek() != 5 {
		t.Error("quene peek error")
	}
}
