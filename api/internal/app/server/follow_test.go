package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"social-network/internal/store/sqlstore"
	"social-network/pkg/jwttoken"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	sourceID = "source"
	targetID = "target"
)

func TestFollow_Privacy_Public(t *testing.T) {
	// new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	//create server to use this test on
	store := sqlstore.New(db)
	s := newServer(store)
	//mock expectation
	mock.ExpectQuery("SELECT type_id FROM privacy").
		WithArgs(targetID).
		WillReturnRows(sqlmock.NewRows([]string{"type_id"}).AddRow(1))
	mock.ExpectExec("INSERT INTO follower").
		WithArgs(sqlmock.AnyArg(), sourceID, targetID, sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/auth/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a recorder to capture the response
	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

}

func TestFollow_Privacy_Private(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT type_id FROM privacy").
		WithArgs(targetID).
		WillReturnRows(sqlmock.NewRows([]string{"type_id"}).AddRow(2))
	mock.ExpectQuery("SELECT \\* FROM request").
		WithArgs(s.types.Privacy.Private, sourceID, targetID).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), s.types.Request.Follow, sourceID, targetID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("GET", "/api/v1/auth/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

}

func TestFollow_Privacy_Invalid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT type_id FROM privacy").
		WithArgs(targetID).
		WillReturnRows(sqlmock.NewRows([]string{"type_id"}).AddRow(5))
	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(), s.types.Request.Follow, sourceID, targetID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("GET", "/api/v1/auth/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}

}

func TestFollow_Invalid_Unauthorized_Request(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	req, err := http.NewRequest("GET", "/api/v1/auth/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestUnfollow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM follower").
		WithArgs(sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := sqlstore.New(db)
	s := newServer(store)

	req, err := http.NewRequest("GET", "/api/v1/auth/unfollow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestFollow_Request_Reject(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("DELETE FROM request").
		WithArgs(s.types.Request.Follow, sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	requestBody := fmt.Sprintf(`{"target_id": "%s", "option": "reject"}`, sourceID)
	req, err := http.NewRequest("POST", "/api/v1/auth/follow/request", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, targetID)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

}

func TestFollow_Request_Accept(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM request").
		WithArgs(s.types.Privacy.Private, sourceID, targetID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type_id", "source_id", "target_id", "parent_id", "message", "created_at"}).AddRow(1, s.types.Privacy.Private, sourceID, targetID, "", "Test Message", time.Now()))
	mock.ExpectExec("DELETE FROM request").
		WithArgs(s.types.Request.Follow, sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO follower").
		WithArgs(sqlmock.AnyArg(), sourceID, targetID, sourceID, targetID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{"target_id": "%s", "option": "accept"}`, sourceID)

	req, err := http.NewRequest("POST", "/api/v1/auth/follow/request", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, targetID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

}

func generateValidToken(t *testing.T, userID string) string {
	claims := jwttoken.NewClaims()
	claims.SetTime("exp", time.Now().Add(time.Hour*2))
	claims.Set("user_id", userID)

	alg := jwttoken.HmacSha256(os.Getenv(jwtKey))
	token, err := alg.Encode(claims)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	return token
}
