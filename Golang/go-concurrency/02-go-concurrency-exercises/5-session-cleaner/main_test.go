package main

import (
	"testing"
	"time"
)

func TestSessionManagersCreationAndUpdate(t *testing.T) {
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		t.Error("Error CreateSession:", err)
	}

	data, err := m.GetSessionData(sID)
	if err != nil {
		t.Error("Error GetSessionData:", err)
	}

	data["website"] = "longhoang.de"
	err = m.UpdateSessionData(sID, data)
	if err != nil {
		t.Error("Error UpdateSessionData:", err)
	}

	data, err = m.GetSessionData(sID)
	if err != nil {
		t.Error("Error GetSessionData:", err)
	}

	if data["website"] != "longhoang.de" {
		t.Error("Expected website to be longhoang.de")
	}
}

func TestSessionManagersCleaner(t *testing.T) {
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		t.Error("Error CreateSession:", err)
	}

	time.Sleep(7 * time.Second)
	_, err = m.GetSessionData(sID)
	if err != ErrSessionNotFound {
		t.Error("Session still in memory after 7 seconds")
	}
}

func TestSessionManagersCleanerAfterUpdate(t *testing.T) {
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		t.Error("Error CreateSession:", err)
	}

	time.Sleep(3 * time.Second)

	err = m.UpdateSessionData(sID, make(map[string]interface{}))
	if err != nil {
		t.Error("Error UpdateSessionData:", err)
	}

	time.Sleep(3 * time.Second)

	_, err = m.GetSessionData(sID)
	if err == ErrSessionNotFound {
		t.Error("Session not found although has been updated 3 seconds earlier.")
	}

	time.Sleep(4 * time.Second)
	_, err = m.GetSessionData(sID)
	if err != ErrSessionNotFound {
		t.Error("Session still in memory 7 seconds after update")
	}
}
