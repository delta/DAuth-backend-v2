// controller/student_controller.go
package controller

import (
	"encoding/json"
	"net/http"

	"https://github.com/delta/DAuth-backend-v2/tree/main/model"
)

func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	// Parse request data
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Validate the data (you can add more validation as needed)
	if student.RollNumber == "" || student.Name == "" || student.Email == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Save the student information to the database or perform necessary actions
	// ...

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Student registered successfully"))
}
