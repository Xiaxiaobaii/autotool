package autotool

import (
	"github.com/kardianos/service"
)

func NewCap(maxToken, limitSpeed int) *TokenBucket {
	return &TokenBucket{
		Maxtoken:   maxToken,
		TokenChan:  make(chan struct{}, maxToken),
		LimitSpeed: limitSpeed,
		stop:       true,
	}
}

func (cap *TokenBucket) Start() {
	cap.stop = false
	go cap.OutPut()
}

func (cap *TokenBucket) OutPut() {
	for {
		if cap.stop {
			close(cap.TokenChan)
			return
		}
		cap.TokenChan <- struct{}{}
		Sleep(cap.LimitSpeed)
	}
}

func (cap *TokenBucket) Stop() {
	cap.stop = true
}

func (cap *TokenBucket) Store(number int) {
	for number > 0 {
		<-cap.TokenChan
		number -= 1
	}
}

func (arg *ServiceLess) Stop(s service.Service) error {
	return nil
}

func (arg *ServiceLess) Start(s service.Service) error {
	arg.Run()
	return nil
}

/*基于数组实现的栈*/
type ArrayStack[T interface{}] struct {
	Data       []T //数据
	maxPolicy  StackPolicy
	sizeCursor int
	maxSize    int
}

type StackPolicy string

const (
	Drop StackPolicy = "Drop"
)

/*初始化栈*/
func NewArrayStack[T any](size int) *ArrayStack[T] {
	return &ArrayStack[T]{
		Data:       make([]T, size),
		maxPolicy:  Drop,
		maxSize:    size,
		sizeCursor: 0,
	}
}

/*栈的长度*/
func (s *ArrayStack[T]) Size() int {
	return len(s.Data)
}

/*栈是否为空*/
func (s *ArrayStack[T]) IsEmpty() bool {
	return s.sizeCursor == 0
}

/*入栈*/
func (s *ArrayStack[T]) Push(v T) {
	if s.Size() == s.sizeCursor {
		s.Pop()
		s.sizeCursor -= 1
	}
	s.Data[s.sizeCursor] = v
	s.sizeCursor += 1
}

/*出栈*/
func (s *ArrayStack[T]) Pop() any {
	if s.Size() == 0 {
		return 0
	}
	val := s.Peek()
	temp := s.Data[:len(s.Data)-1]
	s.Data = make([]T, s.maxSize)
	copy(s.Data, temp)
	s.sizeCursor -= 1
	return val
}

/* 获取栈顶元素 */
func (s *ArrayStack[T]) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	val := s.Data[s.sizeCursor-1]
	return val
}

/* 获取 Slice 用于打印 */
func (s *ArrayStack[T]) ToSlice() []T {
	return s.Data
}
