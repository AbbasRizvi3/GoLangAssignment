package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbbasRizvi3/GoLangAssignment.git/core/app"
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/workers"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/routers"
)

type TestStruct struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
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

	var respondedData []workers.Task
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
	json.Unmarshal(rec.Body.Bytes(), &respondedData)
	fmt.Print(respondedData)
}
