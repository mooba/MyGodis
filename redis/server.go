// author pengchengbai@shopee.com
// date 2021/7/18

package redis

import (
	"context"
	"net"
)

type Handler struct {

}

func MakeHandler() *Handler {
	return &Handler{}
}


func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn)  {

}