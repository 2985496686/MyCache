package MyCache

import (
	"MyCache/protobuf/proto"
	"fmt"
	proto2 "github.com/golang/protobuf/proto"
	"io"
	"net/http"
)

type HttpGetter struct {
	baseUrl string
}

func (h *HttpGetter) Get(req *proto.Request, res *proto.Response) error {
	var url = fmt.Sprintf("%v%v/%v",
		h.baseUrl, req.Group, req.Key,
	)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.Status != "200 OK" {
		return fmt.Errorf("server return error code:%s", resp.Status)
	}
	value, _ := io.ReadAll(resp.Body)
	if err = proto2.Unmarshal(value, res); err != nil {
		return fmt.Errorf("decoding response body:%v", err)
	}
	return nil
}
