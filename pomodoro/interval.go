package pomodoro

import(
	"context" // Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"errors"
	"fmt"
	"time"
)

// declaration of various time constants for pomo
const(
	CategoryPomodoro = "pomodoro"
	CategoryShortBreak = "ShortBreak"
	CategoryLongBreak= "LongBreak"
	
)

// state constants
const (
	StateNotStarted = iota
	StateRunning
	StatePaused
	StateDone
	StateCancelled
)
// defining struct to represent pomo interval
type Interval struct{
	ID int64
	StartTime time.Time
	PlannedDuration time.Duration
	ActualDuration time.Duration
	Category string
	state int
}
type Repository interface {
	Create(i Interval) (int64,error)
	Update(i Interval) error
	ByID(id int64) (Interval,error)
	Last () (Interval,error)
	Breaks(n int) ([]Interval,error)
 }
