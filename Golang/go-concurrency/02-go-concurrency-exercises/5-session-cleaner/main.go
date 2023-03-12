package main

import (
	"errors"
	"log"
)

type SessionManager struct {
	sessions map[string]Session
}

type Session struct {
	Data map[string]interface{}
}

func NewSessionManager() *SessionManager {
	m := &SessionManager{
		sessions: make(map[string]Session),
	}

	return m
}

func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.sessions[sessionID] = Session{
		Data: make(map[string]interface{}),
	}

	return sessionID, nil
}

var ErrSessionNotFound = errors.New("SessionID does not exists")

func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
	session, ok := m.sessions[sessionID]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.Data, nil
}

func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	_, ok := m.sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}

	m.sessions[sessionID] = Session{
		Data: data,
	}

	return nil
}

func main() {
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	data := make(map[string]interface{})
	data["website"] = "longhoang.de"

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to longhoang.de")

	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}
