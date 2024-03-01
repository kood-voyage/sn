package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"social-network/internal/store/sqlstore"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestHandleCreateNotification_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), s.types.Request.Notification, sourceID, targetID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{"source_id": "%s", "target_id": "%s", "message": "%s"}`, sourceID, targetID, "TestMessage")
	req, err := http.NewRequest("POST", "/api/v1/auth/notification/create", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

}

func TestHandleCreateNotification_FromAnotherAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), s.types.Request.Notification, sourceID, targetID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{"source_id": "%s", "target_id": "%s", "message": "%s"}`, sourceID, targetID, "TestMessage")
	req, err := http.NewRequest("POST", "/api/v1/auth/notification/create", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, targetID)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}

}

func TestHandleCreateNotification_InvalidInput(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), s.types.Request.Notification, sourceID, targetID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{"source_id": "%s", "message": "%s"}`, sourceID, "TestMessage")
	req, err := http.NewRequest("POST", "/api/v1/auth/notification/create", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotAcceptable {
		t.Errorf("Expected status code %d, got %d", http.StatusNotAcceptable, rec.Code)
	}

}

func TestHandleDeleteNotification_Success(t *testing.T) {
	requestid := "somerequestid"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("DELETE FROM request").
		WithArgs(requestid).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("DELETE", "/api/v1/auth/notification/delete/"+requestid, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

}
