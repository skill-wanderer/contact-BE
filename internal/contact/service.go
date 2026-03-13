package contact

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Contact struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
}

type createContactRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Service struct {
	mu       sync.Mutex
	nextID   int64
	contacts []Contact
}

func NewService() *Service {
	return &Service{nextID: 1, contacts: make([]Contact, 0)}
}

func (s *Service) ListContacts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s.mu.Lock()
	defer s.mu.Unlock()

	_ = json.NewEncoder(w).Encode(s.contacts)
}

func (s *Service) CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req createContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, `{"error":"name and email are required"}`, http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	contact := Contact{
		ID:        s.nextID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now().UTC(),
	}
	s.nextID++
	s.contacts = append(s.contacts, contact)
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(contact)
}
