// author pengchengbai@shopee.com
// date 2021/7/18

package tcp

import (
	"context"
	"net"
)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
