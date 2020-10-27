package cache_keys

import "fmt"

func AccessTokenKey(accessToken string) string {
	return fmt.Sprintf("ac:%s", accessToken)
}
