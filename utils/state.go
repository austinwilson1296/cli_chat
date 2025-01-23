package utils

import (
	"sync"
	"time"
	
)


type UserState struct {
	mu         sync.RWMutex
	currentUser *User
	isLoggedIn  bool
	lastActive  time.Time
}

type User struct{
	Username string
}


var State = &UserState{}

func (s *UserState) SetUser(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.currentUser = user
	s.isLoggedIn = user != nil
	s.lastActive = time.Now()
}

func (s *UserState) GetCurrentUser() *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currentUser
}

func (s *UserState) IsLoggedIn() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isLoggedIn
}

func (s *UserState) Logout() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.currentUser = nil
	s.isLoggedIn = false
}

func (s *UserState) UpdateLastActive() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastActive = time.Now()
}