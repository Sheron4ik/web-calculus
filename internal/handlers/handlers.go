package handlers

import (
	"go/token"
	"net/http"
	"strconv"

	"github.com/Sheron4ik/web-calculus/internal/models"
	"github.com/Sheron4ik/web-calculus/internal/orchestrator"
	"github.com/Sheron4ik/web-calculus/pkg/calculus"
	"github.com/labstack/echo/v4"
)

func HandleCalculate(c echo.Context) error {
	var req struct {
		Expression string `json:"expression"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid request"})
	}

	calc, err := calculus.NewCalculator(req.Expression)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	err = calc.BuildTasks(calc.Expr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	orchestrator.Mu.Lock()
	id := int64(len(orchestrator.Exprs) + 1)
	orchestrator.Exprs = append(orchestrator.Exprs, &models.Expression{
		Id:     id,
		Status: models.Idle,
		Result: 0,
	})
	orchestrator.Calcs = append(orchestrator.Calcs, calc)
	orchestrator.Mu.Unlock()

	return c.JSON(http.StatusCreated, map[string]int64{"id": id})
}

func HandleListExpressions(c echo.Context) error {
	orchestrator.Mu.RLock()
	defer orchestrator.Mu.RUnlock()
	if len(orchestrator.Exprs) == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"expressions": "empty list expressions"})
	}

	for idx, calc := range orchestrator.Calcs {
		if orchestrator.Exprs[idx].Status == models.InProgress {
			if result, success := calc.GetResult(); success {
				orchestrator.Exprs[idx].Status = models.Completed
				orchestrator.Exprs[idx].Result = result
			}
		}
	}
	return c.JSON(http.StatusOK, map[string][]*models.Expression{"expressions": orchestrator.Exprs})
}

func HandleGetExpression(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid id"})
	}

	orchestrator.Mu.RLock()
	defer orchestrator.Mu.RUnlock()

	if id > len(orchestrator.Exprs) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "expression does not exist"})
	}

	if result, success := orchestrator.Calcs[id-1].GetResult(); success {
		orchestrator.Exprs[id-1].Status = models.Completed
		orchestrator.Exprs[id-1].Result = result
	}
	return c.JSON(http.StatusOK, map[string]*models.Expression{"expression": orchestrator.Exprs[id-1]})
}

func HandleGetTask(c echo.Context) error {
	orchestrator.Mu.Lock()
	defer orchestrator.Mu.Unlock()

	var task models.Task
	for idx, calc := range orchestrator.Calcs {
		if orchestrator.Exprs[idx].Status == models.Failed {
			continue
		}

		id, arg1, arg2, op := calc.GetTask()
		switch op {
		case token.ADD:
			task = models.Task{
				Id:            id,
				Arg1:          arg1,
				Arg2:          arg2,
				Operation:     "addition",
				OperationTime: orchestrator.Cfg.TimeAddMs,
			}
		case token.SUB:
			task = models.Task{
				Id:            id,
				Arg1:          arg1,
				Arg2:          arg2,
				Operation:     "subtraction",
				OperationTime: orchestrator.Cfg.TimeSubMs,
			}
		case token.MUL:
			task = models.Task{
				Id:            id,
				Arg1:          arg1,
				Arg2:          arg2,
				Operation:     "multiplication",
				OperationTime: orchestrator.Cfg.TimeMulMs,
			}
		case token.QUO:
			if arg2 == 0 {
				orchestrator.Exprs[idx].Status = models.Failed
				continue
			}
			task = models.Task{
				Id:            id,
				Arg1:          arg1,
				Arg2:          arg2,
				Operation:     "division",
				OperationTime: orchestrator.Cfg.TimeDivMs,
			}
		case token.EOF:
			continue
		default:
			orchestrator.Exprs[idx].Status = models.Failed
			continue
		}

		orchestrator.Exprs[idx].Status = models.InProgress
		return c.JSON(http.StatusOK, map[string]models.Task{"task": task})
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "no tasks available"})
}

func HandleUpdateTask(c echo.Context) error {
	var req struct {
		ID     int64   `json:"id"`
		Result float64 `json:"result"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
	}

	orchestrator.Mu.RLock()
	defer orchestrator.Mu.RUnlock()

	for _, calc := range orchestrator.Calcs {
		if success := calc.UpdateTask(req.ID, req.Result); success {
			return c.JSON(http.StatusOK, map[string]string{"status": "success"})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}
