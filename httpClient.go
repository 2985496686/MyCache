package MyCache

import (
	"fmt"
	"io"
	"net/http"
)

type HttpGetter struct {
	baseUrl string
}

func (h *HttpGetter) Get(group string, key string) ([]byte, error) {
	var url = fmt.Sprintf("%v%v/%v",
		h.baseUrl, group, key,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("server return error code:%s", resp.Status)
	}
	value, _ := io.ReadAll(resp.Body)
	return value, nil
}
