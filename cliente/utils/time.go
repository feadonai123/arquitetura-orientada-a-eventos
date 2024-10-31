package utils

import (
	"fmt"
	"time"
)

func DelayMs(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func Timestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}