package person_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindPersonByPhone(ctx context.Context, phone string) ([]byte, error)
	FindPersonByID(ctx context.Context, accountId string) ([]byte, error)
	FindPersonProfile(ctx context.Context, _id string) ([]byte, error)
	FindPersonCrendential(ctx context.Context, _id string) ([]byte, error)
	Update(ctx context.Context, _id string, person any) (*mongo.UpdateResult, error)
	Save(ctx context.Context, person any) (*mongo.InsertOneResult, error)
}
