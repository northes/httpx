package httpx

import (
	"fmt"
)

func ToBearerToken(token string) string {
	return fmt.Sprintf("%s %s", BearerTokenPrefix, token)
}
