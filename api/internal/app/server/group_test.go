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

const (
	groupName        = "groupName"
	groupDescription = "groupDescription"
)

func TestGroup_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("INSERT INTO community").
		WithArgs(sqlmock.AnyArg(), sourceID, groupName, groupDescription).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO privacy").
		WithArgs(sqlmock.AnyArg(), s.types.Privacy.Public).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{
    "name":"%s",
    "description":"%s",
	"privacy":"%s"
}`, groupName, groupDescription, "public")

	req, err := http.NewRequest("POST", "/api/v1/auth/group/create", strings.NewReader(requestBody))
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

func TestGroup_Create_InvalidInput(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectExec("INSERT INTO community").
		WithArgs(sqlmock.AnyArg(), sourceID, groupName, groupDescription).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO privacy").
		WithArgs(sqlmock.AnyArg(), s.types.Privacy.Public).
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{
    "name":"%s",
    "description":"%s",
}`, groupName, groupDescription)

	req, err := http.NewRequest("POST", "/api/v1/auth/group/create", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestGroup_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM community").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "name", "description"}).AddRow("groupid", sourceID, groupName, groupDescription))

	mock.ExpectQuery("SELECT member.user_id, member.type_id FROM member JOIN member_type on  member.type_id = member_type.id").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_type"}).AddRow("user1", s.types.Member.Admin))

	mock.ExpectExec("UPDATE community").
		WithArgs(groupName+"test", groupDescription+"test", "groupid").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE privacy").
		WithArgs(s.types.Privacy.Private, "groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{
	"id":"groupid",
    "name":"%s",
    "description":"%s",
	"privacy":"%s"
}`, groupName+"test", groupDescription+"test", "private")

	req, err := http.NewRequest("PUT", "/api/v1/auth/group/update", strings.NewReader(requestBody))
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

func TestGroup_Update_InvalidInput(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM community").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "name", "description"}).AddRow("groupid", sourceID, groupName, groupDescription))

	mock.ExpectQuery("SELECT member.user_id, member.type_id FROM member JOIN member_type on  member.type_id = member_type.id").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_type"}).AddRow("user1", s.types.Member.Admin))

	mock.ExpectExec("UPDATE community").
		WithArgs(groupName+"test", groupDescription+"test", "groupid").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE privacy").
		WithArgs(s.types.Privacy.Private, "groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{
	"id":"groupid",
    "name":"%s",
    "description":"%s",
	"privacy": "%s"
}`, groupName, groupDescription, "private")

	req, err := http.NewRequest("PUT", "/api/v1/auth/group/update", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestGroup_Update_Invalid_User(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM community").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "name", "description"}).AddRow("groupid", targetID, groupName, groupDescription))

	mock.ExpectQuery("SELECT member.user_id, member.type_id FROM member JOIN member_type on  member.type_id = member_type.id").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_type"}).AddRow("user1", s.types.Member.Admin))

	mock.ExpectExec("UPDATE community").
		WithArgs(groupName+"test", groupDescription+"test", "groupid").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE privacy").
		WithArgs(s.types.Privacy.Private, "groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	requestBody := fmt.Sprintf(`{
	"id":"groupid",
    "name":"%s",
    "description":"%s",
	"privacy":"%s"
}`, groupName+"test", groupDescription+"test", "private")

	req, err := http.NewRequest("PUT", "/api/v1/auth/group/update", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, rec.Code)
	}
}

func TestGroup_Delete_Invalid_User(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM community").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "name", "description"}).AddRow("groupid", targetID, groupName, groupDescription))

	mock.ExpectQuery("SELECT member.user_id, member.type_id FROM member JOIN member_type on  member.type_id = member_type.id").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_type"}).AddRow("user1", s.types.Member.Admin))

	mock.ExpectExec("DELETE FROM community").
		WithArgs("groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM privacy").
		WithArgs("groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("DELETE", "/api/v1/auth/group/delete/"+"groupid", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := generateValidToken(t, sourceID)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, rec.Code)
	}
}

func TestGroup_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	mock.ExpectQuery("SELECT \\* FROM community").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "name", "description"}).AddRow("groupid", sourceID, groupName, groupDescription))

	mock.ExpectQuery("SELECT member.user_id, member.type_id FROM member JOIN member_type on  member.type_id = member_type.id").
		WithArgs("groupid").
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_type"}).AddRow("user1", s.types.Member.Admin))

	mock.ExpectExec("DELETE FROM community").
		WithArgs("groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM privacy").
		WithArgs("groupid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("DELETE", "/api/v1/auth/group/delete/"+"groupid", nil)
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
