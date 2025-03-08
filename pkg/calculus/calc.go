package calculus

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"

	"github.com/Sheron4ik/web-calculus/pkg/errors"
)

var TaskID int64 = 1

const (
	TaskNumber = iota
	TaskBinaryOp
)

type Task struct {
	Type         int
	ID           int64
	IsInProgress bool
	Value        float64
	Op           token.Token
	LeftIdx      int
	RightIdx     int
}

type Calculator struct {
	Expr  ast.Expr
	tasks []*Task
	stack []int
}

func NewCalculator(expr string) (*Calculator, error) {
	fset := token.NewFileSet()
	astree, err := parser.ParseExprFrom(fset, "", []byte(expr), 0)
	if err != nil {
		return nil, errors.ErrInvalidExpression
	}

	return &Calculator{
		Expr:  astree,
		tasks: make([]*Task, 0),
		stack: make([]int, 0),
	}, nil
}

func (c *Calculator) BuildTasks(expr ast.Expr) error {
	switch node := expr.(type) {
	case *ast.BinaryExpr:
		if err := c.BuildTasks(node.X); err != nil {
			return err
		}
		if err := c.BuildTasks(node.Y); err != nil {
			return err
		}

		rightIdx := c.stack[len(c.stack)-1]
		leftIdx := c.stack[len(c.stack)-2]
		c.stack = c.stack[:len(c.stack)-2]

		c.tasks = append(c.tasks, &Task{
			Type:         TaskBinaryOp,
			ID:           TaskID,
			IsInProgress: false,
			Op:           node.Op,
			LeftIdx:      leftIdx,
			RightIdx:     rightIdx,
		})
		TaskID++
		c.stack = append(c.stack, len(c.tasks)-1)
		return nil

	case *ast.BasicLit:
		val, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return err
		}

		c.tasks = append(c.tasks, &Task{
			Type:  TaskNumber,
			Value: val,
		})
		c.stack = append(c.stack, len(c.tasks)-1)
		return nil

	case *ast.ParenExpr:
		return c.BuildTasks(node.X)

	default:
		return errors.ErrInvalidExpression
	}
}

func (c *Calculator) GetTask() (int64, float64, float64, token.Token) {
	for _, task := range c.tasks {
		switch task.Type {
		case TaskBinaryOp:
			if !task.IsInProgress && c.tasks[task.LeftIdx].Type == TaskNumber && c.tasks[task.RightIdx].Type == TaskNumber {
				task.IsInProgress = true
				return task.ID, c.tasks[task.LeftIdx].Value, c.tasks[task.RightIdx].Value, task.Op
			}
		}
	}
	return 0, 0, 0, token.EOF
}

func (c *Calculator) UpdateTask(id int64, result float64) bool {
	for _, task := range c.tasks {
		if task.ID == id {
			task.Type = TaskNumber
			task.Value = result
			return true
		}
	}
	return false
}

func (c *Calculator) GetResult() (float64, bool) {
	if len(c.tasks) > 0 && c.tasks[len(c.tasks)-1].Type == TaskNumber {
		return c.tasks[len(c.tasks)-1].Value, true
	}
	return 0, false
}
