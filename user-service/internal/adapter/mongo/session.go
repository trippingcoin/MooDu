package mongo

import (
	"context"

	"github.com/aftosmiros/moodu/user-service/internal/adapter/mongo/dao"
	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefreshToken struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionRefreshTokens = "refreshTokens"
)

func NewRefreshToken(conn *mongo.Database) *RefreshToken {
	return &RefreshToken{
		conn:       conn,
		collection: collectionRefreshTokens,
	}
}

func (r *RefreshToken) Create(ctx context.Context, session domain.Session) error {
	_, err := r.conn.Collection(r.collection).InsertOne(ctx, dao.FromSession(session))

	return err
}

func (r *RefreshToken) GetByToken(ctx context.Context, token string) (domain.Session, error) {
	var session dao.Session
	err := r.conn.Collection(r.collection).FindOne(ctx, bson.M{"refreshToken": token}).Decode(&session)
	if err != nil {
		return domain.Session{}, err
	}

	return dao.ToSession(session), nil
}

func (r *RefreshToken) DeleteByToken(ctx context.Context, token string) error {
	_, err := r.conn.Collection(r.collection).DeleteOne(ctx, bson.M{"refreshToken": token})

	return err
}
