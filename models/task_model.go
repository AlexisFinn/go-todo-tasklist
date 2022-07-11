package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	TASK_TODO = "Todo"
	TASK_DONE = "Done"
)

type Task struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Status   string             `json:"status,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" validate:"required"`
}
