package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository struct {
	collection *mongo.Collection
}

func NewBaseRepository(collection *mongo.Collection) *BaseRepository {
	return &BaseRepository{collection: collection}
}

func (r *BaseRepository) GetCollection() *mongo.Collection {
	return r.collection
}

func (r *BaseRepository) Create(ctx context.Context, document interface{}) error {
    _, err := r.collection.InsertOne(ctx, document)
    if err != nil {
        return err
    }
    return nil
}

// FindByID finds a document by its ID
func (r *BaseRepository) FindByID(ctx context.Context, id interface{}, result interface{}) error {
    err := r.collection.FindOne(ctx, id).Decode(result)
    if err != nil {
        return err
    }
    return nil
}

func (r *BaseRepository) FindAll(ctx context.Context, filter interface{}, result interface{}) error {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

// FindWithPagination finds documents with pagination
func (r *BaseRepository) FindWithPagination(ctx context.Context, filter interface{}, page, pageSize int64) ([]interface{}, error) {
    findOptions := options.Find()
    findOptions.SetSkip((page - 1) * pageSize)
    findOptions.SetLimit(pageSize)

    cursor, err := r.collection.Find(ctx, filter, findOptions)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []interface{}
    for cursor.Next(ctx) {
        var result interface{}
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// FindWhere finds documents that match the specified filter and returns them
func (r *BaseRepository) FindWhere(ctx context.Context, filter interface{}) ([]interface{}, error) {
    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []interface{}
    for cursor.Next(ctx) {
        var result interface{}
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}



// Update updates an existing document in the collection
func (r *BaseRepository) Update(ctx context.Context, id interface{}, update interface{}) (interface{}, error) {
	result := r.collection.FindOneAndUpdate(ctx, id, update)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

// UpdateMany updates multiple documents in the collection
func (r *BaseRepository) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (interface{}, error) {
	result, err := r.collection.UpdateMany(ctx, filter, update)
	return result, err;
}

// Delete deletes a document from the collection
func (r *BaseRepository) Delete(ctx context.Context, id interface{}) error {
    _, err := r.collection.DeleteOne(ctx, id)
    
    return err
}

// DeleteMany deletes multiple documents from the collection
func (r *BaseRepository) DeleteMany(ctx context.Context, filter interface{}) error {
	_, err := r.collection.DeleteMany(ctx, filter)
	return err
}