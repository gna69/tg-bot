package managers

import (
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

type WorkoutManager struct {
	Conn *pgx.Conn
}

func NewWorkoutManager(conn *pgx.Conn) *WorkoutManager {
	return &WorkoutManager{Conn: conn}
}

func (w *WorkoutManager) AddWorkout(workout *entity.Workout) error {
	return nil
}

func (w *WorkoutManager) UpdateWorkout(newWorkout *entity.Workout) error {
	return nil
}

func (w *WorkoutManager) DeleteWorkout(id uint) error {
	return nil
}

func (w *WorkoutManager) GetWorkouts() []entity.Workout {
	return nil
}
