package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var db *sql.DB

// SetDB sets the database connection for the API handlers.
func SetDB(database *sql.DB) {
	db = database
}

// SetupRoutes configures the API routes.
func SetupRoutes(router *mux.Router) {
	// Mothers
	router.HandleFunc("/api/mothers", getMothers).Methods("GET")
	router.HandleFunc("/api/mothers/{id}", getMother).Methods("GET")
	router.HandleFunc("/api/mothers", createMother).Methods("POST")
	router.HandleFunc("/api/mothers/{id}", updateMother).Methods("PUT")
	router.HandleFunc("/api/mothers/{id}", deleteMother).Methods("DELETE")

	// Leaders
	router.HandleFunc("/api/leaders", getLeaders).Methods("GET")
	router.HandleFunc("/api/leaders/{id}", getLeader).Methods("GET")
	router.HandleFunc("/api/leaders", createLeader).Methods("POST")
	router.HandleFunc("/api/leaders/{id}", updateLeader).Methods("PUT")
	router.HandleFunc("/api/leaders/{id}", deleteLeader).Methods("DELETE")

	// Groups
	router.HandleFunc("/api/groups", getGroups).Methods("GET")
	router.HandleFunc("/api/groups/{id}", getGroup).Methods("GET")
	router.HandleFunc("/api/groups", createGroup).Methods("POST")
	router.HandleFunc("/api/groups/{id}", updateGroup).Methods("PUT")
	router.HandleFunc("/api/groups/{id}", deleteGroup).Methods("DELETE")

	// Applications
	router.HandleFunc("/api/applications", getApplications).Methods("GET")
	router.HandleFunc("/api/applications/{id}", getApplication).Methods("GET")
	router.HandleFunc("/api/applications", createApplication).Methods("POST")
	router.HandleFunc("/api/applications/{id}", updateApplication).Methods("PUT")
	router.HandleFunc("/api/applications/{id}", deleteApplication).Methods("DELETE")
}

// Data structures for responses
type Mother struct {
	MotherID               int       `json:"mother_id,omitempty"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	Email                  string    `json:"email"`
	PasswordHash           string    `json:"password_hash,omitempty"`
	PhoneNumber            string    `json:"phone_number,omitempty"`
	Address                string    `json:"address,omitempty"`
	City                   string    `json:"city,omitempty"`
	State                  string    `json:"state,omitempty"`
	ZipCode                string    `json:"zip_code,omitempty"`
	DateOfBirth            string    `json:"date_of_birth,omitempty"`
	ChildrenInfo           string    `json:"children_info,omitempty"`
	PreferredMeetingTimes  string    `json:"preferred_meeting_times,omitempty"`
	PreferredLocations     string    `json:"preferred_locations,omitempty"`
	AccountStatus          string    `json:"account_status,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty"`
}

type Leader struct {
	LeaderID       int       `json:"leader_id,omitempty"`
	MotherID       *int      `json:"mother_id,omitempty"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"password_hash,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	LeaderStatus   string    `json:"leader_status,omitempty"`
	ApprovedByAdmin bool     `json:"approved_by_admin"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type Group struct {
	GroupID          int       `json:"group_id,omitempty"`
	Name             string    `json:"name"`
	Description      string    `json:"description,omitempty"`
	LeaderID         int       `json:"leader_id"`
	MaxMembers       int       `json:"max_members"`
	CurrentMembers   int       `json:"current_members"`
	Status           string    `json:"status,omitempty"`
	MeetingDetails   string    `json:"meeting_details,omitempty"`
	MeetingFrequency string    `json:"meeting_frequency,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

type Application struct {
	ApplicationID               int       `json:"application_id,omitempty"`
	MotherID                    int       `json:"mother_id"`
	GroupID                     int       `json:"group_id"`
	ApplicationDate             time.Time `json:"application_date,omitempty"`
	ApplicationStatus           string    `json:"application_status,omitempty"`
	InterviewDate               *time.Time `json:"interview_date,omitempty"`
	InterviewTimeSlot           *time.Time `json:"interview_time_slot,omitempty"`
	InterviewConfirmedByMother  bool      `json:"interview_confirmed_by_mother"`
	RejectionReason             string    `json:"rejection_reason,omitempty"`
	AcceptedDate                *time.Time `json:"accepted_date,omitempty"`
	ActivatedDate               *time.Time `json:"activated_date,omitempty"`
	NotesFromLeader             string    `json:"notes_from_leader,omitempty"`
	CreatedAt                   time.Time `json:"created_at,omitempty"`
	UpdatedAt                   time.Time `json:"updated_at,omitempty"`
}

// MOTHERS CRUD handlers
func getMothers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT mother_id, first_name, last_name, email, phone_number, city, state, account_status FROM MOTHERS")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mothers []Mother
	for rows.Next() {
		var mother Mother
		err := rows.Scan(&mother.MotherID, &mother.FirstName, &mother.LastName, &mother.Email, &mother.PhoneNumber, &mother.City, &mother.State, &mother.AccountStatus)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mothers = append(mothers, mother)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mothers)
}

func getMother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var mother Mother
	err := db.QueryRow("SELECT mother_id, first_name, last_name, email, phone_number, address, city, state, zip_code, date_of_birth, children_info, preferred_meeting_times, preferred_locations, account_status FROM MOTHERS WHERE mother_id = ?", id).Scan(
		&mother.MotherID, &mother.FirstName, &mother.LastName, &mother.Email, &mother.PhoneNumber, &mother.Address, &mother.City, &mother.State, &mother.ZipCode, &mother.DateOfBirth, &mother.ChildrenInfo, &mother.PreferredMeetingTimes, &mother.PreferredLocations, &mother.AccountStatus,
	)

	if err != nil {
		http.Error(w, "Mother not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mother)
}

func createMother(w http.ResponseWriter, r *http.Request) {
	var mother Mother
	err := json.NewDecoder(r.Body).Decode(&mother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mother.PasswordHash == "" {
		mother.PasswordHash = "hashed_" + mother.Email
	}

	result, err := db.Exec("INSERT INTO MOTHERS (first_name, last_name, email, password_hash, phone_number, address, city, state, zip_code, date_of_birth, children_info, preferred_meeting_times, preferred_locations) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		mother.FirstName, mother.LastName, mother.Email, mother.PasswordHash, mother.PhoneNumber, mother.Address, mother.City, mother.State, mother.ZipCode, mother.DateOfBirth, mother.ChildrenInfo, mother.PreferredMeetingTimes, mother.PreferredLocations)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	mother.MotherID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mother)
}

func updateMother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var mother Mother
	err := json.NewDecoder(r.Body).Decode(&mother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE MOTHERS SET first_name = ?, last_name = ?, email = ?, phone_number = ?, address = ?, city = ?, state = ?, zip_code = ?, date_of_birth = ?, children_info = ?, preferred_meeting_times = ?, preferred_locations = ?, updated_at = CURRENT_TIMESTAMP WHERE mother_id = ?",
		mother.FirstName, mother.LastName, mother.Email, mother.PhoneNumber, mother.Address, mother.City, mother.State, mother.ZipCode, mother.DateOfBirth, mother.ChildrenInfo, mother.PreferredMeetingTimes, mother.PreferredLocations, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mother.MotherID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mother)
}

func deleteMother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := db.Exec("DELETE FROM MOTHERS WHERE mother_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// LEADERS CRUD handlers
func getLeaders(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT leader_id, first_name, last_name, email, phone_number, leader_status, approved_by_admin FROM LEADERS")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var leaders []Leader
	for rows.Next() {
		var leader Leader
		err := rows.Scan(&leader.LeaderID, &leader.FirstName, &leader.LastName, &leader.Email, &leader.PhoneNumber, &leader.LeaderStatus, &leader.ApprovedByAdmin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		leaders = append(leaders, leader)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaders)
}

func getLeader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var leader Leader
	err := db.QueryRow("SELECT leader_id, mother_id, first_name, last_name, email, phone_number, bio, leader_status, approved_by_admin FROM LEADERS WHERE leader_id = ?", id).Scan(
		&leader.LeaderID, &leader.MotherID, &leader.FirstName, &leader.LastName, &leader.Email, &leader.PhoneNumber, &leader.Bio, &leader.LeaderStatus, &leader.ApprovedByAdmin,
	)

	if err != nil {
		http.Error(w, "Leader not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leader)
}

func createLeader(w http.ResponseWriter, r *http.Request) {
	var leader Leader
	err := json.NewDecoder(r.Body).Decode(&leader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if leader.PasswordHash == "" {
		leader.PasswordHash = "hashed_" + leader.Email
	}

	result, err := db.Exec("INSERT INTO LEADERS (mother_id, first_name, last_name, email, password_hash, phone_number, bio, leader_status, approved_by_admin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		leader.MotherID, leader.FirstName, leader.LastName, leader.Email, leader.PasswordHash, leader.PhoneNumber, leader.Bio, leader.LeaderStatus, leader.ApprovedByAdmin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	leader.LeaderID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(leader)
}

func updateLeader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var leader Leader
	err := json.NewDecoder(r.Body).Decode(&leader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE LEADERS SET first_name = ?, last_name = ?, email = ?, phone_number = ?, bio = ?, leader_status = ?, approved_by_admin = ?, updated_at = CURRENT_TIMESTAMP WHERE leader_id = ?",
		leader.FirstName, leader.LastName, leader.Email, leader.PhoneNumber, leader.Bio, leader.LeaderStatus, leader.ApprovedByAdmin, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	leader.LeaderID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leader)
}

func deleteLeader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := db.Exec("DELETE FROM LEADERS WHERE leader_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GROUPS CRUD handlers
func getGroups(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT group_id, name, description, leader_id, max_members, current_members, status, meeting_frequency FROM GROUPS")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.GroupID, &group.Name, &group.Description, &group.LeaderID, &group.MaxMembers, &group.CurrentMembers, &group.Status, &group.MeetingFrequency)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		groups = append(groups, group)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var group Group
	err := db.QueryRow("SELECT group_id, name, description, leader_id, max_members, current_members, status, meeting_details, meeting_frequency FROM GROUPS WHERE group_id = ?", id).Scan(
		&group.GroupID, &group.Name, &group.Description, &group.LeaderID, &group.MaxMembers, &group.CurrentMembers, &group.Status, &group.MeetingDetails, &group.MeetingFrequency,
	)

	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	var group Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO GROUPS (name, description, leader_id, max_members, current_members, status, meeting_details, meeting_frequency) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		group.Name, group.Description, group.LeaderID, group.MaxMembers, group.CurrentMembers, group.Status, group.MeetingDetails, group.MeetingFrequency)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	group.GroupID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var group Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE GROUPS SET name = ?, description = ?, leader_id = ?, max_members = ?, current_members = ?, status = ?, meeting_details = ?, meeting_frequency = ?, updated_at = CURRENT_TIMESTAMP WHERE group_id = ?",
		group.Name, group.Description, group.LeaderID, group.MaxMembers, group.CurrentMembers, group.Status, group.MeetingDetails, group.MeetingFrequency, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group.GroupID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := db.Exec("DELETE FROM GROUPS WHERE group_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// APPLICATIONS CRUD handlers
func getApplications(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT application_id, mother_id, group_id, application_status, interview_confirmed_by_mother FROM APPLICATIONS")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var applications []Application
	for rows.Next() {
		var app Application
		err := rows.Scan(&app.ApplicationID, &app.MotherID, &app.GroupID, &app.ApplicationStatus, &app.InterviewConfirmedByMother)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		applications = append(applications, app)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(applications)
}

func getApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var app Application
	err := db.QueryRow("SELECT application_id, mother_id, group_id, application_status, interview_date, interview_time_slot, interview_confirmed_by_mother, rejection_reason, accepted_date, activated_date, notes_from_leader FROM APPLICATIONS WHERE application_id = ?", id).Scan(
		&app.ApplicationID, &app.MotherID, &app.GroupID, &app.ApplicationStatus, &app.InterviewDate, &app.InterviewTimeSlot, &app.InterviewConfirmedByMother, &app.RejectionReason, &app.AcceptedDate, &app.ActivatedDate, &app.NotesFromLeader,
	)

	if err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app)
}

func createApplication(w http.ResponseWriter, r *http.Request) {
	var app Application
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO APPLICATIONS (mother_id, group_id, application_status, notes_from_leader) VALUES (?, ?, ?, ?)",
		app.MotherID, app.GroupID, "Pending", app.NotesFromLeader)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	app.ApplicationID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(app)
}

func updateApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var app Application
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE APPLICATIONS SET application_status = ?, interview_date = ?, interview_time_slot = ?, interview_confirmed_by_mother = ?, rejection_reason = ?, accepted_date = ?, activated_date = ?, notes_from_leader = ?, updated_at = CURRENT_TIMESTAMP WHERE application_id = ?",
		app.ApplicationStatus, app.InterviewDate, app.InterviewTimeSlot, app.InterviewConfirmedByMother, app.RejectionReason, app.AcceptedDate, app.ActivatedDate, app.NotesFromLeader, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.ApplicationID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app)
}

func deleteApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := db.Exec("DELETE FROM APPLICATIONS WHERE application_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}