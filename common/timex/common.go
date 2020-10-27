package timex

import "time"

func NowMs() int64 {
	return time.Now().UnixNano() / 1e6
}
