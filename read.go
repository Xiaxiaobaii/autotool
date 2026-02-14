package autotool

type Reader struct {
	data    []byte
	counter int
}

func NewReader(buff []byte) *Reader {
	return &Reader{
		data:    buff,
		counter: 0,
	}
}

func (read *Reader) ReadByte() (byte, error) {
	if len(read.data) <= read.counter {
		return 0, Error("Buff Len Max")
	}
	read.counter += 1
	return read.data[read.counter-1], nil
}

func (read *Reader) ReadBuf(Buff *[]byte) error {
	Len := len(*Buff)
	var err error
	for i := 0; i < Len; i++ {
		(*Buff)[i], err = read.ReadByte()
	}
	return err
}