package handlers

import (
    "encoding/json"
    "net/http"
    "student-crud/models"
)

// StudentStore is our in-memory database
var StudentStore = make(map[string]models.Student)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    http.ServeFile(w, r, "static/index.html")
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case "GET":
        studentList := make([]models.Student, 0)
        for _, student := range StudentStore {
            studentList = append(studentList, student)
        }
        json.NewEncoder(w).Encode(studentList)
        
    case "POST":
        var student models.Student
        if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        StudentStore[student.ID] = student
        json.NewEncoder(w).Encode(student)
        
    case "PUT":
        var student models.Student
        if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if _, exists := StudentStore[student.ID]; !exists {
            http.Error(w, "Student not found", http.StatusNotFound)
            return
        }
        StudentStore[student.ID] = student
        json.NewEncoder(w).Encode(student)
        
    case "DELETE":
        id := r.URL.Query().Get("id")
        if _, exists := StudentStore[id]; !exists {
            http.Error(w, "Student not found", http.StatusNotFound)
            return
        }
        delete(StudentStore, id)
        w.WriteHeader(http.StatusOK)
        
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
