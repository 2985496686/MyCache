package MyCache

type PeerPicker interface {
	PickPeer(key string)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
