package bytecache

// ByteCache cache k/v (byte array only)
type ByteCache interface {
	Set(k, v []byte) bool
	Del(k []byte)
	Get(k []byte) (v []byte, exists bool)
	Clear()
	Save()
}
