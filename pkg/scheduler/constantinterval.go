package scheduler

import "time"

type ConstantIntervalScheduledEvent struct {
	duration time.Duration
	function func()
}

type ConstantIntervalScheduler struct {
	events []ConstantIntervalScheduledEvent

	Scheduler
}

func NewConstantIntervalScheduler() *ConstantIntervalScheduler {
	return &ConstantIntervalScheduler{}
}

func (s *ConstantIntervalScheduler) Schedule(d time.Duration, f func()) {
	s.events = append(s.events, ConstantIntervalScheduledEvent{d, f})
}
