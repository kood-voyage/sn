package server

import (
	"net/http"
	"net/http/httptest"
	"social-network/internal/store/sqlstore"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	privacy     = "public"
	privacyFail = "publicprivate"
)

func TestUser_Register(t *testing.T) {
	tests := []struct {
		privacyValue string
		expectedCode int
	}{
		{privacy, http.StatusCreated},
		{privacyFail, http.StatusUnprocessableEntity},
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

func TestUser_Privacy_Update(t *testing.T) {
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

func TestUser_Follower(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT source_id FROM follower").
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

func TestUser_Following(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT target_id FROM follower").
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

func TestUser_Notifications(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM request").
		WithArgs(sourceID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type_id", "source_id", "target_id", "parent_id", "message", "created_at"}).
			AddRow("testID1", s.types.Request.Notification, targetID, sourceID, "parent_id", "test1", time.Now()).
			AddRow("testID2", s.types.Request.Notification, targetID, sourceID,
			"parent_id", "test2", time.Now()))

	req, err := http.NewRequest("GET", "/api/v1/auth/user/notifications", nil)
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

// func TestUser_Posts(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	store := sqlstore.New(db)
// 	s := newServer(store)

// 	mock.ExpectQuery("SELECT post.id, post.title, post.content, post.user_id, post.created_at FROM post JOIN user ON post.user_id = user.id JOIN privacy AS post_privacy ON post.id = post_privacy.id").
// 		WithArgs(targetID, sourceID, targetID, sourceID, targetID, sourceID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "user_id", "created_at"}).
// 			AddRow("postID1", "Its a cool post", "Today we are talking about posts", sourceID, time.Now()).
// 			AddRow("postID2", "Its another cool post", "Test description", targetID, time.Now()))

// 	req, err := http.NewRequest("GET", "/api/v1/auth/user/posts/"+targetID, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	token := generateValidToken(t, sourceID)
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	rec := httptest.NewRecorder()
// 	s.ServeHTTP(rec, req)
// 	if rec.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
// 	}
// }
