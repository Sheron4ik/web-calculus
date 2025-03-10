package calculus

import (
	"go/token"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	calc, err := NewCalculator("2 + 3")
	if err != nil {
		t.Errorf("Ожидалась успешная инициализация, получена ошибка: %v", err)
	}
	if calc == nil {
		t.Error("Ожидался экземпляр калькулятора, получен nil")
	}

	_, err = NewCalculator("2 +")
	if err == nil {
		t.Error("Ожидалась ошибка для невалидного выражения, ошибка не получена")
	}
}

func TestBuildTasks(t *testing.T) {
	calc, _ := NewCalculator("2 + 3")
	err := calc.BuildTasks(calc.Expr)
	if err != nil {
		t.Errorf("Ожидалось успешное построение задач, получена ошибка: %v", err)
	}
	if len(calc.tasks) != 3 {
		t.Errorf("Ожидалось 3 задачи (2, 3, +), получено: %d", len(calc.tasks))
	}

	calc, _ = NewCalculator("2 + 3 * 4")
	err = calc.BuildTasks(calc.Expr)
	if err != nil {
		t.Errorf("Ожидалось успешное построение задач, получена ошибка: %v", err)
	}
	if len(calc.tasks) != 5 {
		t.Errorf("Ожидалось 5 задач (2, 3, 4, *, +), получено: %d", len(calc.tasks))
	}

	calc, _ = NewCalculator("(2 + 3) * 4")
	err = calc.BuildTasks(calc.Expr)
	if err != nil {
		t.Errorf("Ожидалось успешное построение задач, получена ошибка: %v", err)
	}
	if len(calc.tasks) != 5 {
		t.Errorf("Ожидалось 5 задач (2, 3, +, 4, *), получено: %d", len(calc.tasks))
	}
}

func TestGetTask(t *testing.T) {
	calc, _ := NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)

	id, left, right, op := calc.GetTask()
	if id == 0 {
		t.Error("Ожидалась задача, получено отсутствие задач")
	}
	if left != 2 || right != 3 || op != token.ADD {
		t.Errorf("Ожидалась задача c 2, 3 и +, получено: %f, %f, %v", left, right, op)
	}

	calc.UpdateTask(id, 5)
	id, _, _, _ = calc.GetTask()
	if id != 0 {
		t.Errorf("Ожидалось отсутствие задач после выполнения, получена задача c ID: %d", id)
	}
}

func TestUpdateTask(t *testing.T) {
	calc, _ := NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	id, _, _, _ := calc.GetTask()

	success := calc.UpdateTask(id, 5)
	if !success {
		t.Error("Ожидалось успешное обновление задачи, обновление не выполнено")
	}

	success = calc.UpdateTask(999, 0)
	if success {
		t.Error("Ожидалась неудача при обновлении несуществующей задачи, обновление выполнено")
	}
}

func TestGetResult(t *testing.T) {
	calc, _ := NewCalculator("2 + 3")
	calc.BuildTasks(calc.Expr)
	id, _, _, _ := calc.GetTask()
	calc.UpdateTask(id, 5)

	result, ok := calc.GetResult()
	if !ok || result != 5 {
		t.Errorf("Ожидался результат 5, получено: %f, успех: %v", result, ok)
	}

	calc, _ = NewCalculator("2 + 3 * 4")
	calc.BuildTasks(calc.Expr)
	_, ok = calc.GetResult()
	if ok {
		t.Error("Ожидалось отсутствие результата, так как не все задачи выполнены, результат получен")
	}
}
