package worker

type Worker interface {
	Do()
}

type worker struct {
}

func (w *worker) Do() {
	panic("implement me")
}

func NewWorker() Worker {
	return &worker{}
}
