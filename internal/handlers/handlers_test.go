package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sheron4ik/web-calculus/internal/models"
	"github.com/Sheron4ik/web-calculus/internal/orchestrator"
	"github.com/Sheron4ik/web-calculus/pkg/calculus"
	"github.com/labstack/echo/v4"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	e.POST("/calculate", HandleCalculate)
	e.GET("/expressions", HandleListExpressions)
	e.GET("/expressions/:id", HandleGetExpression)
	e.GET("/task", HandleGetTask)
	e.POST("/task", HandleUpdateTask)
	return e
}

func TestHandleCalculate(t *testing.T) {
	e := setupEcho()

	reqBody, _ := json.Marshal(map[string]string{"expression": "2 + 3"})
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := HandleCalculate(c); err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if rec.Code != http.StatusCreated {
		t.Errorf("Ожидался код состояния 201, получен: %d", rec.Code)
	}

	reqBody, _ = json.Marshal(map[string]string{"expression": "2 +"})
	req = httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	HandleCalculate(c)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("Ожидался код состояния 422, получен: %d", rec.Code)
	}
}

func TestHandleListExpressions(t *testing.T) {
	e := setupEcho()

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil

	calc, _ := calculus.NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	orchestrator.Exprs = append(orchestrator.Exprs, &models.Expression{Id: 1, Status: models.Idle, Result: 0})
	orchestrator.Calcs = append(orchestrator.Calcs, calc)

	req := httptest.NewRequest(http.MethodGet, "/expressions", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := HandleListExpressions(c); err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния 200, получен: %d", rec.Code)
	}

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil
	req = httptest.NewRequest(http.MethodGet, "/expressions", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	HandleListExpressions(c)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Ожидался код состояния 500, получен: %d", rec.Code)
	}
}

func TestHandleGetExpression(t *testing.T) {
	e := setupEcho()

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil

	calc, _ := calculus.NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	orchestrator.Exprs = append(orchestrator.Exprs, &models.Expression{Id: 1, Status: models.Idle, Result: 0})
	orchestrator.Calcs = append(orchestrator.Calcs, calc)

	req := httptest.NewRequest(http.MethodGet, "/expressions/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	if err := HandleGetExpression(c); err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния 200, получен: %d", rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/expressions/2", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")
	HandleGetExpression(c)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Ожидался код состояния 404, получен: %d", rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/expressions/invalid", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")
	HandleGetExpression(c)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Ожидался код состояния 500, получен: %d", rec.Code)
	}
}

func TestHandleGetTask(t *testing.T) {
	e := setupEcho()

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil

	calc, _ := calculus.NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	orchestrator.Exprs = append(orchestrator.Exprs, &models.Expression{Id: 1, Status: models.Idle, Result: 0})
	orchestrator.Calcs = append(orchestrator.Calcs, calc)

	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := HandleGetTask(c); err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния 200, получен: %d", rec.Code)
	}

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil
	req = httptest.NewRequest(http.MethodGet, "/task", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	HandleGetTask(c)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Ожидался код состояния 404, получен: %d", rec.Code)
	}
}

func TestHandleUpdateTask(t *testing.T) {
	e := setupEcho()

	orchestrator.Exprs = nil
	orchestrator.Calcs = nil

	calc, _ := calculus.NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	orchestrator.Exprs = append(orchestrator.Exprs, &models.Expression{Id: 1, Status: models.Idle, Result: 0})
	orchestrator.Calcs = append(orchestrator.Calcs, calc)

	id, _, _, _ := calc.GetTask()

	reqBody, _ := json.Marshal(map[string]interface{}{"id": float64(id), "result": 5.0})
	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := HandleUpdateTask(c); err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния 200, получен: %d", rec.Code)
	}

	reqBody, _ = json.Marshal(map[string]interface{}{"id": 999.0, "result": 0.0})
	req = httptest.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	HandleUpdateTask(c)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Ожидался код состояния 404, получен: %d", rec.Code)
	}
}
