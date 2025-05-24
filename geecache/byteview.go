package geecache

type ByteView struct {
	b []byte
}

func NewByteView() *ByteView {
	return &ByteView{b:make([]byte,0)}
}

func (v ByteView)Len() int {
	return len(v.b)
}

func (v ByteView)ByteSlice()[]byte{
	return v.b
}

func (v ByteView)String()string{
	return string(v.b)
}

func cloneBytes(b []byte)[]byte {
	clone := make([]byte, len(b))
	copy(clone, b)
	return clone
}


