// author pengchengbai@shopee.com
// date 2021/7/18

package tcp

import (
	"context"
	"net"
	"pengcheng.site/godis/conf"
	"pengcheng.site/godis/conf/log"
	"sync"
)

func ListenAndServe(cfg *conf.Config, handler Handler, ctx context.Context) {
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Panicf("cannot listen with tcp config=%v ", cfg)
	}

	//cfg.Address = listener.Addr().String()
	log.Info("bind: %s, start listening...", cfg.Address)
	// listen signal
	go func() {
		<-ctx.Done()
		log.Warn("shutting down...")
		_ = listener.Close() // listener.Accept() will return err immediately
		_ = handler.Close()  // close connections
	}()

	// listen port
	defer func() {
		// close during unexpected error
		_ = listener.Close()
		_ = handler.Close()
	}()

	wg := sync.WaitGroup{}
	for  {
		conn, err := listener.Accept()
		if err != nil {
			// 通常是由于listener被关闭无法继续监听导致的错误
			log.Errorf("accept err: %v", err)
			break
		}

		log.Info("accept link")
		wg.Add(1)
		go func() {
			defer wg.Done()
			handler.Handle(ctx, conn)
		}()
	}
	wg.Wait()
}
