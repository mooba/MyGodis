// author pengchengbai@shopee.com
// date 2021/7/18

package redis

import (
	"bufio"
	"context"
	"io"
	"net"
	"pengcheng.site/godis/conf/log"
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
	reader := bufio.NewReader(conn)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Info("connection close")
			} else {
				log.Errorf("error reading: %v", err)
			}
			return
		}
		echoMsg := []byte(readString)
		log.Info("message received: ", string(echoMsg))

		conn.Write(echoMsg)
	}
}
