package handler

import (
	"encoding/json"
	"go-todo-rest-api-example/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	tasks := []model.Task{}
	if err := db.Model(&project).Related(&tasks).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, tasks)
}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	task := model.Task{ProjectID: project.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusCreated, task)
}

func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}
	responseJSON(w, http.StatusOK, task)
}

func UpdateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, task)
}

func DeleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusNoContent, nil)
}

func CompleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}

	task.Complete()
	if err := db.Save(&task).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, task)
}

func UndoTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}

	task.Undo()
	if err := db.Save(&task).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, task)
}

func getTaskOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Task {
	task := model.Task{}
	if err := db.First(&task, id).Error; err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &task
}
