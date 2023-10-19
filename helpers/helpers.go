package helpers

import (
	"time"
)

const charsetAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "123456789"

// GetNow generate time local zone Asia Jakarta
func GetNow() (time.Time, string) {
	TimeZoneJKT, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(TimeZoneJKT)
	return now, now.Format("2006-01-02 15:04:05")
}
