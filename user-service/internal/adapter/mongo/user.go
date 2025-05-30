package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aftosmiros/moodu/user-service/internal/adapter/mongo/dao"
	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db         *mongo.Database
	collection string
}

const userCollection = "users"

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db:         db,
		collection: userCollection,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, dao.FromUser(user))
	if mongo.IsDuplicateKeyError(err) {
		return domain.ErrEmailAlreadyExists
	}
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	_, err := domain.ParseID(id)
	if err != nil {
		return nil, domain.ErrInvalidID
	}

	var u dao.User
	err = r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id, "is_deleted": false}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	domainUser := dao.ToUser(&u)
	return domainUser, nil
}

func (r *UserRepository) GetByBarcode(ctx context.Context, barcode string) (*domain.User, error) {
	var u dao.User
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"barcode": barcode, "is_deleted": false}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return dao.ToUser(&u), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u dao.User
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"email": email, "is_deleted": false}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return dao.ToUser(&u), nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": dao.FromUser(user)}

	_, err := r.db.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) SoftDelete(ctx context.Context, id string) error {
	objID, err := domain.ParseID(id)
	if err != nil {
		return domain.ErrInvalidID
	}

	_, err = r.db.Collection(r.collection).UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"is_deleted": true}},
	)
	return err
}

func (r *UserRepository) EnsureIndexes(ctx context.Context) error {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{"email": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("unique_email"),
		},
		{
			Keys: bson.M{"barcode": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("unique_barcode"),
		},
	}

	_, err := r.db.Collection(r.collection).Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}
	return nil
}
