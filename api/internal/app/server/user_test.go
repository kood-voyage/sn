package server

import (
	"net/http"
	"net/http/httptest"
	"social-network/internal/store/sqlstore"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	privacy     = "public"
	privacyFail = "publicprivate"
)

func TestUserRegister(t *testing.T) {
	tests := []struct {
		privacyValue string
		expectedCode int
	}{
		{privacy, http.StatusCreated},
		{privacyFail, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.privacyValue, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error creating mock database: %v", err)
			}
			defer db.Close()

			store := sqlstore.New(db)
			s := newServer(store)

			mock.ExpectExec("INSERT INTO user").
				WithArgs(sourceID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectExec("INSERT INTO privacy").
				WithArgs(sourceID, s.types.Privacy.Values[tt.privacyValue]).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req, err := http.NewRequest("GET", "/api/v1/auth/user/create/"+tt.privacyValue, nil)
			if err != nil {
				t.Fatal(err)
			}

			token := generateValidToken(t, sourceID)
			req.Header.Set("Authorization", "Bearer "+token)

			rec := httptest.NewRecorder()

			s.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, rec.Code)
			}
		})
	}
}

func TestUserPrivacyUpdate(t *testing.T) {
	tests := []struct {
		privacyValue string
		expectedCode int
	}{
		{privacy, http.StatusOK},
		{privacyFail, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.privacyValue, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error creating mock database: %v", err)
			}
			defer db.Close()

			store := sqlstore.New(db)
			s := newServer(store)

			mock.ExpectExec("UPDATE privacy").
				WithArgs(s.types.Privacy.Values[tt.privacyValue], sourceID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req, err := http.NewRequest("GET", "/api/v1/auth/user/privacy/"+tt.privacyValue, nil)
			if err != nil {
				t.Fatal(err)
			}

			token := generateValidToken(t, sourceID)
			req.Header.Set("Authorization", "Bearer "+token)

			rec := httptest.NewRecorder()

			s.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, rec.Code)
			}
		})
	}
}

func TestHandleUser_GetFollowers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT source_id FROM").
		WithArgs(targetID).WillReturnRows(sqlmock.NewRows([]string{"source_id"}).AddRow("followerID1").AddRow("followerID2"))

	req, err := http.NewRequest("GET", "/api/v1/auth/user/followers/"+targetID, nil)
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

func TestHandleUser_GetFollowing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT target_id FROM").
		WithArgs(sourceID).WillReturnRows(sqlmock.NewRows([]string{"target_id"}).AddRow("followerID1").AddRow("followerID2"))

	req, err := http.NewRequest("GET", "/api/v1/auth/user/following/"+sourceID, nil)
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
