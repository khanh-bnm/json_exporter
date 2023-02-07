package exporter

import (
	"strconv"
	"strings"
	"time"
)

func ParseRatesFreshness(rt string) (string, error) {
	s := strings.Split(rt, ".")
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, s[0])
	freshness := time.Now().Unix() - t.Unix()
	return strconv.FormatInt(freshness, 10), err
}
