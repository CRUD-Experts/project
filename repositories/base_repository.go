package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository[T any] struct {
	collection *mongo.Collection
}

func NewBaseRepository[T any](collection *mongo.Collection) *BaseRepository[T] {
	return &BaseRepository[T]{collection: collection}
}

func (r *BaseRepository[T]) GetCollection() *mongo.Collection {
	return r.collection
}

// Create a new document
func (r *BaseRepository[T]) Create(ctx context.Context, document T) (interface{}, error) {
	result, err := r.collection.InsertOne(ctx, document)
	
	return result.InsertedID, err
}

// Find a document by ID
func (r *BaseRepository[T]) FindByID(ctx context.Context, id interface{}, result *T) (interface{}, error) {
	err := r.collection.FindOne(ctx, id).Decode(result)
	
	return result, err
}


// Find all documents
func (r *BaseRepository[T]) FindAll(ctx context.Context, filter interface{}, result *[]T) (interface{}, error) {
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, result)
	
	return result, err
}

// Find all documents with pagination
func (r *BaseRepository[T]) FindWithPagination(ctx context.Context, filter interface{}, page, pageSize int64) ([]T, error) {
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * pageSize)
	findOptions.SetLimit(pageSize)

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []T
	err = cur.All(ctx, &results)
	
	return results, err
}

// Update a document
func (r *BaseRepository[T]) Update(ctx context.Context, filter interface{}, update interface{}) (interface{}, error) {
	result, err := r.collection.UpdateOne(ctx, filter, update)
	
	return result, err
}

// Update many documents
func (r *BaseRepository[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (interface{}, error) {
	result, err := r.collection.UpdateMany(ctx, filter, update)
	
	return result, err
}

// Delete a document
func (r *BaseRepository[T]) Delete(ctx context.Context, filter interface{}) (interface{}, error) {
	result, err := r.collection.DeleteOne(ctx, filter)
	
	return result, err
}

// Delete many documents
func (r *BaseRepository[T]) DeleteMany(ctx context.Context, filter interface{}) (interface{}, error) {
	result, err := r.collection.DeleteMany(ctx, filter)
	
	return result, err
}
