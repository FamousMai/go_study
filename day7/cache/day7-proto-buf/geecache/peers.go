package geecache

import pb "day7/cache/day7-proto-buf/geecache/geecachepb"

// PeerPicker 的 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter
// 接口 PeerGetter 的 Get() 方法用于从对应 group 查找缓存值。PeerGetter 就对应于上述流程中的 HTTP 客户端。
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
