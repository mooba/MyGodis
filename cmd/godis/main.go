// author pengchengbai@shopee.com
// date 2021/7/17

package main

import (
	"bufio"
	"context"
	"io"
	"net"
	"pengcheng.site/godis/cmd"
	"pengcheng.site/godis/conf"
	"pengcheng.site/godis/conf/log"
	RedisServer "pengcheng.site/godis/redis"
	"pengcheng.site/godis/tcp"
)




func handleConn(conn net.Conn) {
	log.Info("message received")
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

		conn.Write(echoMsg)
	}
}


var tcpConfig = conf.GetTcpConfig()

func Init()  {
	// init logger
	log.NewLoggerWithRotate()

	// init tcp config
	tcpConfig.Init("")
}

func RunTcpServer(ctx context.Context) {
	tcp.ListenAndServe(tcpConfig, RedisServer.MakeHandler(), ctx)
}



func main() {
	// 应该先注册还是先init?
	Init()
	cmd.RegisterAndLock(RunTcpServer, "godis-server")
	cmd.Run()
}

