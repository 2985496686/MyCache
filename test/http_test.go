package test

import (
	"MyCache"
	"fmt"
	"log"
	"net/http"
	"testing"
)

var db3 = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *MyCache.Group {
	return MyCache.NewGroup("scores", 2<<10, MyCache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db3[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string) {
	peers := MyCache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, gee *MyCache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func TestServer1(t *testing.T) {
	createGroup()
	MyCache.NewHTTPPool("localhost:8001")
	startCacheServer("http://localhost:8001")
}

func TestServer2(t *testing.T) {
	createGroup()
	MyCache.NewHTTPPool("localhost:8002")
	startCacheServer("http://localhost:8002")
}

func TestServer3(t *testing.T) {
	createGroup()
	MyCache.NewHTTPPool("localhost:8003")
	startCacheServer("http://localhost:8003")
}

func TestClient(t *testing.T) {
	group := createGroup()
	httpPool := MyCache.NewHTTPPool("localhost:9999")
	httpPool.Set("http://127.0.0.1:8001", "http://127.0.0.1:8002", "http://127.0.0.1:8003")
	group.RegisterPeers(httpPool)
	startAPIServer("http://localhost:9999", group)
}
