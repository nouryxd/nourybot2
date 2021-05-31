package humanize

import (
	"time"

	"github.com/dustin/go-humanize"
)

func HumanizeTime(t time.Time) string {
	return humanize.Time(t)
}
