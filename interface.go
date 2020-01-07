package bytecache

// ByteCache  cache key and value as []byte
type ByteCache interface {
	Set(k, v []byte) bool
	Del(k []byte)
	Get(k []byte) (v []byte, exists bool)
	Clear()
	Save()
}
