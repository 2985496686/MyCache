package MyCache

import (
	"MyCache/consistent"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type HTTPPool struct {
	basePath    string
	self        string
	mu          sync.Mutex
	peers       consistent.Map
	httpGetters map[string]*HttpGetter
}

const defaultBasePath = "/my-cache/"

const defaultReplicas = 50

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		basePath: defaultBasePath,
		self:     self,
	}
}

func (h *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", h.self, fmt.Sprintf(format, v...))
}

func (h *HTTPPool) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if hasPrefix := strings.HasPrefix(request.URL.Path, h.basePath); !hasPrefix {
		panic(fmt.Sprintf("HTTPPool server is no expects url path:%s", request.URL.Path))
	}
	split := strings.Split(request.URL.Path, "/")
	if len(split) < 4 {
		panic("url path must have groupName and key")
	}
	groupName := split[2]
	key := split[3]
	group := GetGroup(groupName)
	if group == nil {
		panic(fmt.Sprintf("%s is not exist", groupName))
	}
	v, err := group.Get(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Write(v.ByteSlice())
}

func (h *HTTPPool) Set(peers ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.peers.Add(peers...)
	h.httpGetters = make(map[string]*HttpGetter, len(peers))
	for _, peer := range peers {
		h.httpGetters[peer] = &HttpGetter{peer + h.basePath}
	}
}

func (h *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.peers.Get(key); peer != "" && peer != h.self {
		log.Printf("pick peer %s\n", key)
		httpGetter := h.httpGetters[peer]
		return httpGetter, true
	}
	return nil, false
}
