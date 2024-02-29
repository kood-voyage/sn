package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"social-network/internal/store/sqlstore"
	"social-network/pkg/jwttoken"
	"testing"
	"time"

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

	//create server to use this test on
	store := sqlstore.New(db)
	s := newServer(store)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/follow/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a recorder to capture the response
	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("TestHandleFollowSucceed: Expected status code %d, got %d", http.StatusCreated, rec.Code)
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

func TestHandleUnfollow(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/api/v1/unfollow/"+targetID, nil)
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

func TestHandleFollowRequestSucceed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO request").
		WithArgs(sqlmock.AnyArg(),"follow", sourceID, targetID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := sqlstore.New(db)
	s := newServer(store)

	req, err := http.NewRequest("GET", "/api/v1/follow/request/"+targetID, nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("TestHandleFollowRequestSucceed: Expected status code %d, got %d", http.StatusCreated, rec.Code)
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
