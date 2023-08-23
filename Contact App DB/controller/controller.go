package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"login/contact"
	contactinfo "login/contactInfo"
	"login/services"
)

// Credentials represents the user credentials for login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles the login request
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Credentials:", credentials)
	user, err := services.GetUserByUsername(credentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password != credentials.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&user.ID)
}

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *services.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	newUser, err = services.CreateUser(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

// DeleteUser handles the deletion of a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // string to int
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	newUser, err := services.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

// GetUserById retrieves a user by their ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // string to int
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)

		return
	}

	newUser, err := services.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

// UpdateUser handles the update of a user's details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // string to int
	if err != nil {
		http.Error(w, "Inavalid user ID", http.StatusBadRequest)
		return
	}

	var newUser *services.User

	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newUser, err = services.UpdateUser(uint(id), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

// CreateContact handles the creation of a new contact for a user
func CreateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, "Failed to convert uint64 to uint", http.StatusInternalServerError)
		return
	}

	var newContact *contact.Contact

	err = json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newContact, err = contact.CreateContact(uint(uid), newContact.ContactName, newContact.ContactType, newContact.ContactValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newContact)
}

// DeleteContact handles the deletion of a contact
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, "Failed to get the ID", http.StatusBadRequest)
		return
	}

	var deleteContact *contact.Contact

	err = json.NewDecoder(r.Body).Decode(&deleteContact)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	deleteContact, err = contact.DeleteContact(deleteContact.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&deleteContact)
}

// GetContactById retrieves a contact by its ID
func GetContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["userid"]) // string to int
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	newContact, err := contact.GetContactById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newContact)
}

// UpdateContact handles the update of a contact's details
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, "Failed to get the ID", http.StatusBadRequest)
		return
	}

	var updateContact *contact.Contact

	err = json.NewDecoder(r.Body).Decode(&updateContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateContact, err = contact.UpdateContact(updateContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&updateContact)
}

// CreateContactInfo handles the creation of a new contact information for a user
func CreateContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newContactInfo *contactinfo.ContactInfo

	err = json.NewDecoder(r.Body).Decode(&newContactInfo)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newContactInfo, err = contactinfo.CreateContactInfo(uint(uid), newContactInfo.ContactInfoType, newContactInfo.ContactInfoValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newContactInfo)
}

// DeleteContactInfo handles the deletion of a contact information
func DeleteContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, "Failed to get the ID", http.StatusInternalServerError)
		return
	}

	var deleteContactInfo *contactinfo.ContactInfo

	err = json.NewDecoder(r.Body).Decode(&deleteContactInfo)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	deleteContactInfo, err = contactinfo.DeleteContactInfo(deleteContactInfo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&deleteContactInfo)
}

// GetContactInfoById retrieves a contact information by its ID
func GetContactInfoById(w http.ResponseWriter, r *http.Request) {
	var newContactInfo *contactinfo.ContactInfo

	err := json.NewDecoder(r.Body).Decode(&newContactInfo)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newContactInfo, err = contactinfo.GetContactInfoById(newContactInfo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newContactInfo)
}

// UpdateContactInfo handles the update of a contact information's details
func UpdateContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	uid, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(uid))
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var updateContactInfo *contactinfo.ContactInfo

	err = json.NewDecoder(r.Body).Decode(&updateContactInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateContactInfo, err = contactinfo.UpdateContactInfo(updateContactInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&updateContactInfo)
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	pageint, err := strconv.Atoi(page) // string to int
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}
	users, err := services.GetAllUsers(pageint, 10) // 10 is page size or limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&users)
}

// FindAllContacts retrieves all contacts with pagination
func FindAllContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	pageint, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}
	contacts, err := contact.FindAllContact(pageint, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&contacts)
}

// FindAllContactInfo retrieves all contact information with pagination
func FindAllContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	pageint, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}
	cinfo, err := contactinfo.FindAllContactInfo(pageint, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&cinfo)
}
