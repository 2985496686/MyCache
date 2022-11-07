package MyCache

type ByteValue struct {
	b []byte
}

func (v ByteValue) Len() uint64 {
	return uint64(len(v.b))
}

func (v ByteValue) String() string {
	return string(v.b)
}

func (v ByteValue) ByteSlice() []byte {
	return bytesClone(v.b)
}

func bytesClone(b []byte) []byte {
	bytes := make([]byte, len(b))
	copy(bytes, b)
	return bytes
}
