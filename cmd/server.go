// author pengchengbai@shopee.com
// date 2021/7/18

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pengcheng.site/godis/conf/log"
	"pengcheng.site/godis/utils"
	"sync/atomic"
	"syscall"
)

var (
	//ServerName servername.lock, default name is server
	serverName = "server"

	// IsRunning indicate whether still running, 1 means true
	IsRunning = int32(1)

	RunningServerMethod func(ctx context.Context)
)

var ctx context.Context
var cf, release func()


// RegisterAndLock register with running-server method and lock with server-name-based file-lock
func RegisterAndLock(run func(context.Context), name string) {
	serverName = name
	RunningServerMethod = run
	var err error
	release, err = lockServer(name)
	if err != nil {
		log.Infof("error locking server with name=%s, which means another instance is running", name)
		os.Exit(1)
	}

	ctx, cf = context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		log.Infof("received system signal:%d, waiting to stop...", sig)
		cf()
		atomic.StoreInt32(&IsRunning, int32(0))
		fmt.Println("system exit....")
		if release != nil {
			release()
		}
	}()
}

// Run run the registered method
func Run()  {
	RunningServerMethod(ctx)
}

// lock server with file
func lockServer(serverName string) (func(), error) {
	if len(serverName) <= 3 {
		return nil, fmt.Errorf("invalid server name")
	}
	fp, err := os.OpenFile(serverName+".lock", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	if err = syscall.Flock(int(fp.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		return nil, err
	}
	if _, err = fmt.Fprint(fp, os.Getpid()); err != nil {
		return nil, err
	}

	releaseServer := func() {
		fmt.Printf("release file lock....")
		utils.SafeClose(fp)
		os.Remove(serverName + ".lock")
	}
	return releaseServer, nil
}
