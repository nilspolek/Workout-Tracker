package mongodb

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nilspolek/Workout-Tracker/repo"
)

func TestNewWorkoutRepository(t *testing.T) {
	r, err := NewWorkoutRepository("mongodb://admin:admin123@localhost:27017")
	if err != nil {
		t.Errorf("Expected NewWorkoutRepository() to return a nil error, but got %v", err)
	}
	if r == nil {
		t.Error("Expected NewWorkoutRepository() to return a non-nil value")
	}
	defer r.Close()
	id := uuid.New()
	err = r.CreateWorkout(repo.Workout{
		Name:        "test",
		Id:          id,
		UserId:      uuid.New(),
		Cathegory:   "Upper Body",
		Description: "Test ur upper body strength",
		Excersises: []repo.Excersise{
			{
				Id:   uuid.New(),
				Name: "Bench Press",
				Sets: 3,
				Reps: 10,
			},
			{
				Id:   uuid.New(),
				Name: "Flys",
				Sets: 3,
				Reps: 11,
			},
		},
	})
	if err != nil {
		t.Errorf("Expected CreateWorkout() to return a nil error, but got %v", err)
	}
	w, err := r.GetWorkout(id)
	if err != nil {
		t.Errorf("Expected GetWorkout() to return a nil error, but got %v", err)
	}
	if w.Name != "test" {
		t.Errorf("Expected GetWorkout() to return a workout with name test, but got %s", w.Name)
	}
	if w.Excersises[0].Name != "Bench Press" {
		t.Errorf("Expected GetWorkout() to return a workout with excersise Bench Press, but got %s", w.Excersises[0].Name)
	}
	if w.Excersises[1].Name != "Flys" {
		t.Errorf("Expected GetWorkout() to return a workout with excersise Flys, but got %s", w.Excersises[1].Name)
	}
}
