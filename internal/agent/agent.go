package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sheron4ik/web-calculus/internal/models"
)

func Worker(id int, port string) {
	client := &http.Client{}
	for {
		resp, err := client.Get(fmt.Sprintf("http://localhost:%s/internal/task", port))
		if err != nil || resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusInternalServerError {
			time.Sleep(1 * time.Second)
			continue
		}

		var data struct {
			Task models.Task `json:"task"`
		}
		json.NewDecoder(resp.Body).Decode(&data)
		resp.Body.Close()

		opTime, _ := strconv.ParseInt(data.Task.OperationTime, 10, 0)
		time.Sleep(time.Duration(opTime) * time.Millisecond)
		result := compute(data.Task)

		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     data.Task.Id,
			"result": result,
		})
		req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:%s/internal/task", port), bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		client.Do(req)
	}
}

func compute(t models.Task) float64 {
	switch t.Operation {
	case "addition":
		return t.Arg1 + t.Arg2
	case "subtraction":
		return t.Arg1 - t.Arg2
	case "multiplication":
		return t.Arg1 * t.Arg2
	case "division":
		return t.Arg1 / t.Arg2
	}
	return 0
}
