package MyCache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HTTPPool struct {
	basePath string
	self     string
}

const defaultBasePath = "/my-cache/"

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
