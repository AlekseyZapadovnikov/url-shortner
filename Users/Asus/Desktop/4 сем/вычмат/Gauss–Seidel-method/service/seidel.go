package service

import (
	"fmt"
	"math"
)

// SolveGaussSeidel решает Ax=b методом Гаусса–Зейделя.
// Критерий остановки: max_i |x_i^(k)-x_i^(k-1)| < eps
func SolveGaussSeidel(cfg SeidelConfig) (SeidelResult, error) {
	n := len(cfg.A)
	if n == 0 || len(cfg.B) != n {
		return SeidelResult{}, fmt.Errorf("invalid dimensions")
	}
	if cfg.Eps <= 0 {
		return SeidelResult{}, fmt.Errorf("eps must be > 0")
	}
	if cfg.MaxIters <= 0 {
		cfg.MaxIters = 10000
	}

	// Проверка диагонали
	for i := 0; i < n; i++ {
		if len(cfg.A[i]) != n {
			return SeidelResult{}, fmt.Errorf("matrix is not square at row %d", i)
		}
		if cfg.A[i][i] == 0 {
			return SeidelResult{}, fmt.Errorf("zero diagonal element at row %d", i)
		}
	}

	// Начальное приближение: x^0 = d,
	// где d_i = b_i / a_ii.
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = cfg.B[i] / cfg.A[i][i]
	}

	history := make([][]float64, 0, 16)
	history = append(history, CopyVector(x))

	if cfg.Logger != nil {
		cfg.Logger.Printf("Initial approximation x0 = %v", x)
	}

	for iter := 1; iter <= cfg.MaxIters; iter++ {
		prev := CopyVector(x)

		// Обновление "на месте": это и есть ключевая идея Зейделя.
		for i := 0; i < n; i++ {
			var sum1, sum2 float64

			// Используем уже обновленные значения x[0..i-1] (k+1)
			for j := 0; j < i; j++ {
				sum1 += cfg.A[i][j] * x[j]
			}
			// Используем старые значения prev[i+1..n-1] (k)
			for j := i + 1; j < n; j++ {
				sum2 += cfg.A[i][j] * prev[j]
			}

			x[i] = (cfg.B[i] - sum1 - sum2) / cfg.A[i][i]
		}

		delta := make([]float64, n)
		maxDelta := 0.0
		for i := 0; i < n; i++ {
			delta[i] = math.Abs(x[i] - prev[i])
			if delta[i] > maxDelta {
				maxDelta = delta[i]
			}
		}

		history = append(history, CopyVector(x))

		if cfg.Logger != nil {
			cfg.Logger.Printf("iter=%d x=%v max|dx|=%.10g", iter, x, maxDelta)
		}

		if maxDelta < cfg.Eps {
			return SeidelResult{
				X:          x,
				Iterations: iter,
				LastDelta:  delta,
				History:    history,
			}, nil
		}
	}

	return SeidelResult{}, fmt.Errorf("method did not converge within %d iterations", cfg.MaxIters)
}