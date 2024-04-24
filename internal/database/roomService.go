package database

import (
	"context"
	"time"

	"damstudy-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomServiceImpl struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func NewRoomServiceImpl(client *mongo.Client, dbName, collectionName string) *RoomServiceImpl {
	return &RoomServiceImpl{
		db:             client,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (rs *RoomServiceImpl) GetAll() ([]models.Room, error) {
	collection := rs.db.Database(rs.dbName).Collection(rs.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []models.Room
	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rs *RoomServiceImpl) Create(room models.Room) (*mongo.InsertOneResult, error) {
	collection := rs.db.Database(rs.dbName).Collection(rs.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rs *RoomServiceImpl) Update(room models.Room) (*mongo.UpdateResult, error) {
	collection := rs.db.Database(rs.dbName).Collection(rs.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": room.ID}, bson.M{"$set": room})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rs *RoomServiceImpl) Delete(id string) (*mongo.DeleteResult, error) {
	collection := rs.db.Database(rs.dbName).Collection(rs.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rs *RoomServiceImpl) GetByID(id string) (models.Room, error) {
	collection := rs.db.Database(rs.dbName).Collection(rs.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var room models.Room
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&room)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}
