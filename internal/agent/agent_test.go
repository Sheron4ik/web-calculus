package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Sheron4ik/web-calculus/internal/models"
)

func TestWorkerBasic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/internal/task" && r.Method == "GET" {
			task := models.Task{
				Id:            1,
				Arg1:          2.0,
				Arg2:          3.0,
				Operation:     "addition",
				OperationTime: "1000",
			}
			resp := map[string]models.Task{"task": task}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				t.Fatalf("Ошибка при кодировании задачи: %v", err)
			}
		} else if r.URL.Path == "/internal/task" && r.Method == "POST" {
			var result struct {
				ID     int64   `json:"id"`
				Result float64 `json:"result"`
			}
			if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
				t.Fatalf("Ошибка при декодировании результата: %v", err)
			}
			if result.ID != 1 || result.Result != 5.0 {
				t.Errorf("Ожидался ID=1 и Result=5.0, получено ID=%d и Result=%f", result.ID, result.Result)
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	go Worker(1, server.URL[17:])

	time.Sleep(2 * time.Second)
}

func TestWorkerErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	done := make(chan struct{})

	go func() {
		client := &http.Client{}
		for range 2 {
			resp, err := client.Get(fmt.Sprintf("http://localhost:%s/internal/task", server.URL[17:]))
			if err != nil || resp.StatusCode == http.StatusNotFound {
				time.Sleep(1 * time.Second)
				continue
			}
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("Worker не обработал ошибку корректно, завис")
	}
}
