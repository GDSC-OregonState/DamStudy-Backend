package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	Close()
	GetDB() *mongo.Client
	NewRoomService(dbName, collectionName string) *RoomServiceImpl
}

type service struct {
	db *mongo.Client
}

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_ROOT_PASSWORD")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", username, password, host, port, database)))
	log.Printf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, database)

	if err != nil {
		log.Fatal(err)
	}

	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) Close() {
	if err := s.db.Disconnect(context.Background()); err != nil {
		log.Fatalf("error closing db connection: %v", err)
	}
}

func (s *service) GetDB() *mongo.Client {
	return s.db
}

func (s *service) NewRoomService(dbName, collectionName string) *RoomServiceImpl {
	return NewRoomServiceImpl(s.db, dbName, collectionName)
}
