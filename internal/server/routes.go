package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"damstudy-backend/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/rooms", s.AllRoomsHandler)

	r.Post("/rooms", s.CreateRoomHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)

	return r
}

// Get all documents from the rooms collection
func (s *Server) AllRoomsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Make this more DRY
	roomService := s.db.NewRoomService("damstudy", "rooms")
	resp, err := roomService.GetAll()
	if err != nil {
		log.Printf("error getting all rooms: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}

// Create a new document in the rooms collection
func (s *Server) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomService := s.db.NewRoomService("damstudy", "rooms")

	var room struct {
		models.Room
		Seats     string `json:"seats"`
		Tech      string `json:"technology"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("coordinates: %v, %v", room.Latitude, room.Longitude)

	latitude, err := strconv.ParseFloat(room.Latitude, 64)
	if err != nil {
		log.Printf("error converting latitude to float64: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(room.Longitude, 64)
	if err != nil {
		log.Printf("error converting longitude to float64: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	room.Coordinates = models.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}

	room.Room.Technology = strings.Split(room.Tech, ",")

	room.Room.Seats, err = strconv.Atoi(strings.TrimSpace(room.Seats))
	if err != nil {
		log.Printf("error converting seats to integer: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := roomService.Create(room.Room)
	if err != nil {
		log.Printf("error creating room: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	jsonResp, _ := json.Marshal(result)
	_, _ = w.Write(jsonResp)
}

// Delete a document from the rooms collection
func (s *Server) DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomService := s.db.NewRoomService("damstudy", "rooms")

	roomID := chi.URLParam(r, "id")

	result, err := roomService.Delete(roomID)
	if err != nil {
		log.Printf("error deleting room: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(result)
	_, _ = w.Write(jsonResp)
}

// Update a document in the rooms collection
func (s *Server) UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomService := s.db.NewRoomService("damstudy", "rooms")

	var room models.Room

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := roomService.Update(room)
	if err != nil {
		log.Printf("error updating room: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(result)
	_, _ = w.Write(jsonResp)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
