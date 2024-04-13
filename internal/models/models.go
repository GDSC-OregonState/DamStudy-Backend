package models

/*
 * Room represents a room in the database
 */
type Room struct {
	ID         string   `bson:"_id,omitempty"` // MongoDB document ID
	Name       string   `bson:"name"`
	Image      string   `bson:"image"`
	NoiseLevel string   `bson:"noiseLevel"`
	Seats      int      `bson:"seats"`
	Technology []string `bson:"technology"`
	Seating    string   `bson:"seating"`
	Location   string   `bson:"location"`
}

/*
 * RoomService is the interface for the Room service
 */
type RoomService interface {
	GetAll() ([]Room, error)
	GetByID(id string) (Room, error)
	Create(room Room) error
	Update(id string, room Room) error
	Delete(id string) error
}
