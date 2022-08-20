package timer

import (
	"github.com/go-co-op/gocron"
	"time"
)

func Run(call func(), expr string) error {
	s := gocron.NewScheduler(time.Local)
	if _, err := s.CronWithSeconds(expr).SingletonMode().Do(call); err != nil {
		return err
	}
	s.StartAsync()
	return nil
}
