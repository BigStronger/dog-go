package datetime

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Logf("\tFirstDayOfWeek:\t\t%s", FirstDayOfWeek(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tLastDayOfWeek:\t\t%s", LastDayOfWeek(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tFirstDayOfMonth:\t%s", FirstDayOfMonth(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tLastDayOfMonth:\t\t%s", LastDayOfMonth(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tFirstDayOfYear:\t\t%s", FirstDayOfYear(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tLastDayOfYear:\t\t%s", LastDayOfYear(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tTimeModDay:\t\t\t%s", TimeModDay(time.Now().In(time.Local)).Format("2006-01-02 15:04:05"))
	t.Logf("\tToDayUnix:\t\t\t%d", ToDayUnix(time.Now().In(time.Local)))
}
