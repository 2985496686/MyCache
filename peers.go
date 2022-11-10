package MyCache

import "MyCache/protobuf/proto"

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(req *proto.Request, res *proto.Response) error
}
