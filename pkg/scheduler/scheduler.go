package scheduler

import "time"

type Scheduler interface {
	Schedule(d time.Duration, f func())
}
