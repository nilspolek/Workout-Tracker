package repo

import (
	"time"

	"github.com/google/uuid"
)

var (
	JwtKey = []byte("supersecretkey")
)

type User struct {
	Id       uuid.UUID `bson:"id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
	Email    string    `bson:"email"`
}

type Token struct {
	UserID    string    `bson:"user_id"`
	Token     string    `bson:"token"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type UserRepositoryInterface interface {
	GetUsers() ([]User, error)
	GetUser(username string) (User, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	GenerateJWT(userID uuid.UUID) (string, error)
	ValidateJWT(token string) (bool, error)
	Close() error
}

type WorkoutRepoInterface interface {
	GetWorkoutsByUserId(userId uuid.UUID) ([]Workout, error)
	GetWorkoutsByCathegory(cathegory string) ([]Workout, error)
	GetWorkouts() ([]Workout, error)
	GetWorkout(id uuid.UUID) (Workout, error)
	CreateWorkout(workout Workout) error
	UpdateWorkout(workout Workout) error
	DeleteWorkout(id uuid.UUID) error
	Close() error
}

type Workout struct {
	Id          uuid.UUID   `bson:"id"`
	UserId      uuid.UUID   `bson:"userId"`
	Name        string      `bson:"name"`
	Cathegory   string      `bson:"cathegory"`
	Description string      `bson:"description"`
	Excersises  []Excersise `bson:"excersises"`
}

type ExcersiseRepoInterface interface {
	GetExcersises() ([]Excersise, error)
	GetExcersise(id uuid.UUID) (Excersise, error)
	CreateExcersise(excersise Excersise) error
	UpdateExcersise(excersise Excersise) error
	DeleteExcersise(id uuid.UUID) error
	Close() error
}

type Excersise struct {
	Id          uuid.UUID `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	MuscleGroup string    `bson:"muscleGroup"`
	Sets        int       `bson:"sets"`
	Reps        int       `bson:"reps"`
}
