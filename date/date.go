package datetime

import "time"

func FirstDayOfWeek(date time.Time) time.Time {
	offset := int(time.Monday - date.Weekday())
	if offset > 0 {
		offset = -6
	}
	date = date.AddDate(0, 0, offset)
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func LastDayOfWeek(date time.Time) time.Time {
	return FirstDayOfWeek(date).AddDate(0, 0, 6)
}

func FirstDayOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func LastDayOfMonth(date time.Time) time.Time {
	return FirstDayOfMonth(date).AddDate(0, 1, -1)
}

func FirstDayOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
}

func LastDayOfYear(date time.Time) time.Time {
	return FirstDayOfYear(date).AddDate(1, 0, -1)
}

func TimeModDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func ToDayUnix(date time.Time) int64 {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).Unix()
}
