package orchestrator

import (
	"sync"

	"github.com/Sheron4ik/web-calculus/internal/config"
	"github.com/Sheron4ik/web-calculus/internal/models"
	"github.com/Sheron4ik/web-calculus/pkg/calculus"
)

var (
	Cfg                          = config.New()
	Exprs []*models.Expression   = make([]*models.Expression, 0)
	Calcs []*calculus.Calculator = make([]*calculus.Calculator, 0)
	Mu    sync.RWMutex
)
