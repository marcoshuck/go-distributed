package nodes

const (
	StatusCreated status = iota
	StatusRestarting
	StatusRunning
	StatusPaused
	StatusStopped
	StatusExited
	StatusDead
)

type status int64

type Status interface {
	ToString() string
}

func (s status) ToString() string {
	switch s {
	case StatusCreated:
		return "created"
	case StatusRestarting:
		return "restarting"
	case StatusRunning:
		return "running"
	case StatusPaused:
		return "paused"
	case StatusStopped:
		return "stopped"
	case StatusExited:
		return "exited"
	case StatusDead:
		return "dead"
	}
	panic("invalid status")
}
