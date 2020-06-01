package nodes

const (
	StatusCreated status = iota
	StatusConnecting
	StatusErrConnecting
	StatusConnected
	StatusRestarting
	StatusRunning
	StatusPausing
	StatusPaused
	StatusStopping
	StatusStopped
	StatusExiting
	StatusExited
	StatusKilling
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
	case StatusConnecting:
		return "connecting"
	case StatusErrConnecting:
		return "error connecting"
	case StatusConnected:
		return "error connecting"
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
