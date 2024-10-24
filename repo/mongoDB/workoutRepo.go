package mongodb

import (
	"context"

	"github.com/google/uuid"
	"github.com/nilspolek/Workout-Tracker/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoWorkoutRepository struct {
	workoutCollection *mongo.Collection
}

func NewWorkoutRepository(uri string) (repo.WorkoutRepoInterface, error) {
	var (
		err    error
		temp   = MongoWorkoutRepository{}
		client *mongo.Client
		ctx    = context.Background()
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	temp.workoutCollection = client.Database("workout-tracker").Collection("workouts")
	return temp, err
}

func (r MongoWorkoutRepository) GetWorkoutsByUserId(userId uuid.UUID) ([]repo.Workout, error) {
	var (
		workouts = []repo.Workout{}
		ctx      = context.Background()
	)
	cur, err := r.workoutCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var workout repo.Workout
		err = cur.Decode(&workout)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

func (r MongoWorkoutRepository) GetWorkoutsByCathegory(cathegory string) ([]repo.Workout, error) {
	var (
		ctx      = context.Background()
		workouts = []repo.Workout{}
		err      error
	)
	cur, err := r.workoutCollection.Find(ctx, bson.M{"cathegory": cathegory})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var workout repo.Workout
		err = cur.Decode(&workout)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

func (r MongoWorkoutRepository) GetWorkouts() ([]repo.Workout, error) {
	var (
		ctx      = context.Background()
		workouts = []repo.Workout{}
		err      error
	)
	cur, err := r.workoutCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var workout repo.Workout
		err = cur.Decode(&workout)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

func (r MongoWorkoutRepository) GetWorkout(id uuid.UUID) (repo.Workout, error) {
	var (
		ctx     = context.Background()
		workout repo.Workout
		err     error
	)
	err = r.workoutCollection.FindOne(ctx, bson.M{"id": id}).Decode(&workout)
	if err != nil {
		return workout, err
	}
	return workout, nil
}

func (r MongoWorkoutRepository) CreateWorkout(workout repo.Workout) error {
	var (
		ctx = context.Background()
	)
	_, err := r.workoutCollection.InsertOne(ctx, workout)
	if err != nil {
		return err
	}
	return nil
}

func (r MongoWorkoutRepository) UpdateWorkout(workout repo.Workout) error {
	var (
		ctx = context.Background()
	)
	_, err := r.workoutCollection.ReplaceOne(ctx, bson.M{"id": workout.Id}, workout)
	if err != nil {
		return err
	}
	return nil
}

func (r MongoWorkoutRepository) DeleteWorkout(id uuid.UUID) error {
	var (
		ctx = context.Background()
	)
	_, err := r.workoutCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r MongoWorkoutRepository) Close() error {
	return r.workoutCollection.Database().Client().Disconnect(context.Background())
}
