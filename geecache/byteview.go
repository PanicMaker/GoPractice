package geecache

type ByteView struct {
	b []byte
}

// Len return the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice return a copy of the data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// return the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}

// b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
