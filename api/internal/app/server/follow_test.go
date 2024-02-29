package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"social-network/internal/store/sqlstore"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestHandleFollow(t *testing.T) {
	// Create a new SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Set up a mock expectation for your SQL query
	mock.ExpectExec("INSERT INTO follower").
		WithArgs(sqlmock.AnyArg(), "test", "").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), "notification", "test", "", "started following you.").
		WillReturnResult(sqlmock.NewResult(1, 1))
	// Create a new Server instance
	store := sqlstore.New(db)
	s := newServer(store)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/follow/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up a test context with a user ID
	ctx := context.WithValue(req.Context(), ctxUserID, "test")
	req = req.WithContext(ctx)

	// Create a recorder to capture the response
	rec := httptest.NewRecorder()

	// Call the handleFollow function with the request and recorder
	s.handleFollow()(rec, req)

	// Check the response status code
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

}
