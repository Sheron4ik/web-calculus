package models

type Task struct {
	Id            int64   `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime string  `json:"operation_time"`
}

type Expression struct {
	Id     int64   `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

var (
	Idle       = "idle"
	InProgress = "in_progress"
	Completed  = "completed"
	Failed     = "failed"
)
