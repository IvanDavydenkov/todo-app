package entities

import "time"

type TodoStatusEnum string

const (
	TodoStatusCreated    TodoStatusEnum = "created"
	TodoStatusDone       TodoStatusEnum = "done"
	TodoStatusInProgress TodoStatusEnum = "in-progress"
	TodoStatusCanceled   TodoStatusEnum = "canceled"
)

type Todo struct {
	Title       string         `json:"title"`
	Subtitle    string         `json:"subtitle"`
	Id          string         `json:"id"`
	Status      TodoStatusEnum `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	LastChanged time.Time      `json:"lastChanged"`
}
