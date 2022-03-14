package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// 02-01-2006T15:04:05Z : Gün-Ay-Yıl
// 2006-01-02T15:04:05Z
// 01-02-2006T15:04:05Z : Ay-Gün-Yıl
// 2006-01-02T15:04:05Z : Yıl-Ay-Gün

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
