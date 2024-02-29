package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"social-network/internal/store/sqlstore"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	sourceID = "source"
	targetID = "target"
)

func TestHandleFollowSucceed(t *testing.T) {
	// new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	//mock expectation
	mock.ExpectExec("INSERT INTO follower").
		WithArgs(sqlmock.AnyArg(), sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), "notification", sourceID, targetID, "started following you.").
		WillReturnResult(sqlmock.NewResult(1, 1))

	//create server to use this test on
	store := sqlstore.New(db)
	s := newServer(store)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up a test context with a user ID
	ctx := context.WithValue(req.Context(), ctxUserID, sourceID)
	req = req.WithContext(ctx)

	// Create a recorder to capture the response
	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

}

func TestHandleFollowFail(t *testing.T) {
	// new sql mock
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	//create server to use this test on
	store := sqlstore.New(db)
	s := newServer(store)
	//check request without context
	req, err := http.NewRequest("GET", "/api/v1/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}
