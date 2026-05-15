package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	// Set up a test database
	dbFile := "./test_wowmom.db"
	os.Remove(dbFile) // Clean up any old test db

	testDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Printf("Error opening test database: %v\n", err)
		os.Exit(1)
	}
	defer testDB.Close()

	SetDB(testDB)

	schema, err := os.ReadFile("../database/schema.sql")
	if err != nil {
		fmt.Printf("Error reading schema file: %v\n", err)
		os.Exit(1)
	}
	if _, err := testDB.Exec(string(schema)); err != nil {
		fmt.Printf("Error executing schema: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	exitCode := m.Run()

	// Clean up
	os.Remove(dbFile)
	os.Exit(exitCode)
}

func createTestRouter() *mux.Router {
	router := mux.NewRouter()
	SetupRoutes(router)
	return router
}

func TestCreateMother(t *testing.T) {
	router := createTestRouter()

	mother := Mother{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@example.com",
		PasswordHash: "hashedpassword",
	}
	jsonMother, _ := json.Marshal(mother)

	req, _ := http.NewRequest("POST", "/api/mothers", bytes.NewBuffer(jsonMother))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdMother Mother
	json.NewDecoder(rr.Body).Decode(&createdMother)
	if createdMother.FirstName != mother.FirstName {
		t.Errorf("handler returned unexpected body: got %v want %v",
			createdMother.FirstName, mother.FirstName)
	}
}

func TestGetMother(t *testing.T) {
	router := createTestRouter()

	// First, create a mother to retrieve
	mother := Mother{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		PasswordHash: "hashedpassword",
	}
	jsonMother, _ := json.Marshal(mother)

	req, _ := http.NewRequest("POST", "/api/mothers", bytes.NewBuffer(jsonMother))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	var createdMother Mother
	json.NewDecoder(rr.Body).Decode(&createdMother)

	// Now, get the created mother
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/mothers/%d", createdMother.MotherID), nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedMother Mother
	json.NewDecoder(rr.Body).Decode(&fetchedMother)
	if fetchedMother.FirstName != mother.FirstName {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fetchedMother.FirstName, mother.FirstName)
	}
}

func TestUpdateMother(t *testing.T) {
	router := createTestRouter()

	// First, create a mother to update
	mother := Mother{
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice.smith@example.com",
		PasswordHash: "hashedpassword",
	}
	jsonMother, _ := json.Marshal(mother)

	req, _ := http.NewRequest("POST", "/api/mothers", bytes.NewBuffer(jsonMother))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	var createdMother Mother
	json.NewDecoder(rr.Body).Decode(&createdMother)

	// Now, update the created mother
	updatedMother := createdMother
	updatedMother.FirstName = "Alicia"
	jsonUpdatedMother, _ := json.Marshal(updatedMother)

	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/mothers/%d", updatedMother.MotherID), bytes.NewBuffer(jsonUpdatedMother))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedMother Mother
	json.NewDecoder(rr.Body).Decode(&fetchedMother)
	if fetchedMother.FirstName != updatedMother.FirstName {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fetchedMother.FirstName, updatedMother.FirstName)
	}
}

func TestDeleteMother(t *testing.T) {
	router := createTestRouter()

	// First, create a mother to delete
	mother := Mother{
		FirstName: "Bob",
		LastName:  "Johnson",
		Email:     "bob.johnson@example.com",
		PasswordHash: "hashedpassword",
	}
	jsonMother, _ := json.Marshal(mother)

	req, _ := http.NewRequest("POST", "/api/mothers", bytes.NewBuffer(jsonMother))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	var createdMother Mother
	json.NewDecoder(rr.Body).Decode(&createdMother)

	// Now, delete the created mother
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/mothers/%d", createdMother.MotherID), nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Verify mother is deleted
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/mothers/%d", createdMother.MotherID), nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code after delete: got %v want %v",
			status, http.StatusNotFound)
	}
}