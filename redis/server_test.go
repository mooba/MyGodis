// author pengchengbai@shopee.com
// date 2021/7/20

package redis

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}
