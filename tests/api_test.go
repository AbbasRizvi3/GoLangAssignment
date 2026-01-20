package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/core/app"
	routers "github.com/AbbasRizvi3/GoLangAssignment.git/internal/router"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
)

type TestStruct struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}
type GetTaskResponse struct {
	Tasks []tasks.Task `json:"tasks"`
}

func TestHandleAddTask(t *testing.T) {
	routers.SetupRoutes()
	sampleData := TestStruct{
		Name:     "Sample Task",
		Priority: 1,
	}
	marshaledData, err := json.Marshal(sampleData)
	if err != nil {
		t.Fatal(err)
	}
	finalData := bytes.NewBuffer(marshaledData)
	req, err := http.NewRequest("POST", "/tasks", finalData)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.Router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleGetAllTasks(t *testing.T) {

	var respondedData []tasks.Task
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.Router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	err = json.Unmarshal(rec.Body.Bytes(), &respondedData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(respondedData)
}

func TestHandleGetSpecificTask(t *testing.T) {
	app.Tasks.Tasks = nil
	sampleTask := TestStruct{
		Name:     "Specific Task",
		Priority: 2,
	}
	marshaledData, err := json.Marshal(sampleTask)
	if err != nil {
		t.Fatal(err)
	}
	finalData := bytes.NewBuffer(marshaledData)
	reqAdd, err := http.NewRequest("POST", "/tasks", finalData)
	if err != nil {
		t.Fatal(err)
	}
	reqAdd.Header.Set("Content-Type", "application/json")
	recAdd := httptest.NewRecorder()
	app.Router.ServeHTTP(recAdd, reqAdd)

	var addedTaskResponse struct {
		Message string     `json:"message"`
		Task    tasks.Task `json:"task"`
	}
	err = json.Unmarshal(recAdd.Body.Bytes(), &addedTaskResponse)
	if err != nil {
		t.Fatal(err)
	}
	addedTaskID := addedTaskResponse.Task.ID

	reqGet, err := http.NewRequest("GET", "/task/"+addedTaskID, nil)
	if err != nil {
		t.Fatal(err)
	}
	reqGet.Header.Set("Content-Type", "application/json")
	recGet := httptest.NewRecorder()
	app.Router.ServeHTTP(recGet, reqGet)

	if status := recGet.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var getTaskResponse GetTaskResponse

	err = json.Unmarshal(recGet.Body.Bytes(), &getTaskResponse)
	if err != nil {
		t.Fatal(err)
	}
	if len(getTaskResponse.Tasks) == 0 || getTaskResponse.Tasks[0].ID != addedTaskID {
		t.Errorf("Expected to retrieve task with ID %s, but got different result", addedTaskID)
	}
}
