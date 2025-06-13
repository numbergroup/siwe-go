package siwe

import (
	"fmt"
	"strings"
	"time"

	"github.com/dchest/uniuri"
)

const ISO8601Layout = "2006-01-02T15:04:05Z0700"

func parseTimestamp(fields map[string]any, key string) (*string, error) {
	var value string

	if val, ok := fields[key]; ok {
		switch parsedTime := val.(type) {
		case time.Time:
			value = parsedTime.UTC().Format(time.RFC3339)
		case string:
			_, err := time.Parse(time.RFC3339, parsedTime)
			if err != nil {
				return nil, &InvalidMessage{fmt.Sprintf("Invalid format for field `%s`: %s", key, err.Error())}
			}
			value = parsedTime
		default:
			return nil, &InvalidMessage{fmt.Sprintf("`%s` must be either an ISO8601 formatted string or time.Time", key)}
		}
	}

	if value == "" {
		return nil, nil
	}

	return &value, nil
}

func GenerateNonce() string {
	return uniuri.NewLen(16)
}

func isEmpty(str *string) bool {
	return str == nil || len(strings.TrimSpace(*str)) == 0
}

func isStringAndNotEmpty(m map[string]any, k string) (*string, bool) {
	if v, ok := m[k]; ok {
		switch s := v.(type) {
		case string:
			if s != "" {
				return &s, true
			}
		}
	}
	return nil, false
}
