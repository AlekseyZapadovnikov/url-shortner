package service

import (
	"errors"
	"fmt"
	"math"
)

// HasStrictDiagonalDominance проверяет строгое диагональное преобладание по строкам:
// |a_ii| > sum_{j != i} |a_ij|
func HasStrictDiagonalDominance(A [][]float64) bool {
	n := len(A)
	for i := 0; i < n; i++ {
		var sum float64
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			sum += math.Abs(A[i][j])
		}
		if math.Abs(A[i][i]) <= sum {
			return false
		}
	}
	return true
}

// EnsureDiagonalDominanceByRowPermutation пытается добиться диагонального преобладания перестановкой строк.
// Возвращает: была ли перестановка, порядок строк (newRow[i] = oldRowIndex), ошибку при невозможности.
// ВАЖНО: для n<=20 полный перебор допустим, но может быть дорогим на 20.
// Здесь — backtracking с отсечением, обычно хватает.
func EnsureDiagonalDominanceByRowPermutation(A [][]float64) (bool, []int, error) {
	n := len(A)
	if n == 0 {
		return false, nil, errors.New("empty matrix")
	}

	// Если уже есть преобладание — порядок идентичный.
	if HasStrictDiagonalDominance(A) {
		order := make([]int, n)
		for i := 0; i < n; i++ {
			order[i] = i
		}
		return false, order, nil
	}

	used := make([]bool, n)
	order := make([]int, n) // order[targetRow] = sourceRow
	found := false

	var backtrack func(col int)
	backtrack = func(col int) {
		if found {
			return
		}
		if col == n {
			found = true
			return
		}

		for r := 0; r < n; r++ {
			if used[r] {
				continue
			}
			// Проверяем, подходит ли строка r на позицию col (т.е. диагональный элемент будет A[r][col]).
			diag := math.Abs(A[r][col])
			var sum float64
			for j := 0; j < n; j++ {
				if j == col {
					continue
				}
				sum += math.Abs(A[r][j])
			}
			if diag <= sum {
				continue
			}

			used[r] = true
			order[col] = r
			backtrack(col + 1)
			used[r] = false
		}
	}

	backtrack(0)

	if !found {
		return false, nil, fmt.Errorf("невозможно добиться диагонального преобладания перестановкой строк")
	}

	// Применяем порядок к матрице A in-place (через копию).
	old := CopyMatrix(A)
	for i := 0; i < n; i++ {
		copy(A[i], old[order[i]])
	}
	return true, order, nil
}

// ReorderVectorByRowOrder переставляет b в соответствии с перестановкой строк матрицы.
func ReorderVectorByRowOrder(b []float64, order []int) []float64 {
	res := make([]float64, len(b))
	for i := range b {
		res[i] = b[order[i]]
	}
	return res
}

// BuildIterationFormAndNorm строит C и d из Ax=b:
// x_i = -sum_{j!=i}(a_ij/a_ii)x_j + b_i/a_ii
// Возвращает ||C||_inf = max_i sum_j |c_ij|
func BuildIterationFormAndNorm(A [][]float64, b []float64) (normInf float64, C [][]float64, d []float64, err error) {
	n := len(A)
	if n == 0 || len(b) != n {
		return 0, nil, nil, fmt.Errorf("invalid sizes")
	}

	C = make([][]float64, n)
	d = make([]float64, n)

	for i := 0; i < n; i++ {
		if len(A[i]) != n {
			return 0, nil, nil, fmt.Errorf("matrix is not square at row %d", i)
		}
		if A[i][i] == 0 {
			return 0, nil, nil, fmt.Errorf("zero diagonal element at row %d", i)
		}

		C[i] = make([]float64, n)
		var rowSum float64

		for j := 0; j < n; j++ {
			if i == j {
				C[i][j] = 0
				continue
			}
			C[i][j] = -A[i][j] / A[i][i]
			rowSum += math.Abs(C[i][j])
		}
		d[i] = b[i] / A[i][i]
		if rowSum > normInf {
			normInf = rowSum
		}
	}

	return normInf, C, d, nil
}