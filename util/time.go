package util

import "time"

func GetToday0THTime() int64 {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return addTime.Unix()
}

func GetTomorrow0THTime() int64 {
	ts := time.Now().AddDate(0, 0, 1)
	timeTomorrow := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
	return timeTomorrow.Unix()
}

func GetYesterday0THTime() int64 {
	ts := time.Now().AddDate(0, 0, -1)
	timeYesterday := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
	return timeYesterday.Unix()
}
