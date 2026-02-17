package autotool

type CryptType string

type LogType string

type ColorType string

type errorString struct {
	s string
}

type Kvalue[T any] struct {
	key   T
	value T
}

type TokenBucket struct {
	Maxtoken   int
	TokenChan  chan struct{}
	LimitSpeed int
	stop       bool
}

type ServiceLess struct {
	Run func()
}
