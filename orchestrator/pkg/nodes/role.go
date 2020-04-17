package nodes

const (
	RoleManager role = "manager"
	RoleWorker role = "worker"
)

type Role interface {
	IsValid() bool
}

type role string

func (r role) IsValid() bool {
	return r == "manager" || r == "worker"
}