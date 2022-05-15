package usecases

import "github.com/gna69/tg-bot/internal/entity"

type TrainingScheduler interface {
	AddWorkout(workout *entity.Workout) error
	UpdateWorkout(newWorkout *entity.Workout) error
	DeleteWorkout(id uint) error
	GetWorkouts() []entity.Workout
}
