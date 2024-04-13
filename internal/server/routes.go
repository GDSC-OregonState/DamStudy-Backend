package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	r.Get("/mock", s.mockHandler)

	r.Get("/all", s.allRoomsHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)

	return r
}

func (s *Server) allRoomsHandler(w http.ResponseWriter, r *http.Request) {
	roomService := s.db.NewRoomService("damstudy", "rooms")
	resp, err := roomService.GetAll()
	if err != nil {
		log.Printf("error getting all rooms: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}

func (s *Server) mockHandler(w http.ResponseWriter, r *http.Request) {
	roomService := s.db.NewRoomService("damstudy", "rooms")

	room, err := roomService.Create(models.Room{
		Name:       "room1",
		Seats:      10,
		Technology: []string{"tech1", "tech2"},
		NoiseLevel: "noise1",
		Location:   "location1",
		Seating:    "seating1",
		Image:      "image1",
	})
	if err != nil {
		log.Printf("error creating room: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(room)
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
