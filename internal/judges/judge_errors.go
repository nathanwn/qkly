package judges

import (
	"fmt"
)

type UnsupportedJudge struct {
	url string
}

func (err UnsupportedJudge) Error() string {
	return fmt.Sprintf("Unsupported judge with problem url %s", err.url)
}
