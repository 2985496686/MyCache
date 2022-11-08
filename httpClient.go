package MyCache

import (
	"fmt"
	"net/http"
)

type HttpGetter struct {
	baseUrl string
}

func (h *HttpGetter) Get(group string, key string) ([]byte, error) {
	var url = fmt.Sprintf("%v%v/%v",
		defaultBasePath, group, key,
	)
	var value []byte
	resp, err := http.Get(url)
	if err != nil {
		return value, err
	}
	status, _ := resp.Body.Read(value)
	if status != 200 {
		return nil, fmt.Errorf("server return error code:%d", status)
	}
	return value, nil
}
