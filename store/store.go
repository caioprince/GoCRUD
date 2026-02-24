package store

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("User not found")
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
}

type Store struct {
	mu   sync.RWMutex
	data map[uuid.UUID]User
}

func New() *Store {
	return &Store{
		data: make(map[uuid.UUID]User),
	}
}

func (s *Store) FindAll() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.data))
	for _, u := range s.data {
		users = append(users, u)
	}
	return users
}

func (s *Store) FindById(id uuid.UUID) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &user, nil
}

func (s *Store) Insert(newUser User) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	newUser.ID = uuid.New()
	s.data[newUser.ID] = newUser
	return newUser
}

func (s *Store) Update(id uuid.UUID, userUpdates User) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.data[id]
	if !ok {
		return nil, ErrNotFound
	}

	user.FirstName = userUpdates.FirstName
	user.LastName = userUpdates.LastName
	user.Biography = userUpdates.Biography

	s.data[id] = user
	return &user, nil
}

func (s *Store) Delete(id uuid.UUID) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.data[id]
	if !ok {
		return nil, ErrNotFound
	}

	delete(s.data, id)
	return &user, nil
}
