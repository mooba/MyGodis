// author pengchengbai@shopee.com
// date 2021/7/18

package utils

import (
	"io"
	"pengcheng.site/godis/conf/log"
)

func SafeClose(closable io.Closer)  {
	err := closable.Close()
	if err != nil {
		log.Fatalf("error closing listener: %v", closable)
	}
}
