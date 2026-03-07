package service

import (
	"fmt"
	"io"
	"math"
)

// PrintReport — печатает отчет для итерационного метода.
func PrintReport(w io.Writer, r Report) {
	fmt.Fprintln(w, "\n==================== ОТЧЕТ ====================")

	if r.Reordered {
		fmt.Fprintf(w, "Перестановка строк выполнена: да (порядок строк: %v)\n", r.RowOrder)
	} else {
		fmt.Fprintln(w, "Перестановка строк выполнена: нет")
	}

	fmt.Fprintln(w, "\nМатрица A:")
	printMatrix(w, r.A)

	fmt.Fprintln(w, "\nВектор b:")
	printVector(w, r.B)

	fmt.Fprintln(w, "\nИтерационная форма x = Cx + d")
	fmt.Fprintln(w, "Матрица C:")
	printMatrix(w, r.C)

	fmt.Fprintln(w, "\nВектор d:")
	printVector(w, r.D)

	fmt.Fprintf(w, "\nНорма матрицы ||C||_inf: %.10f\n", r.CNormInf)

	fmt.Fprintln(w, "\nРешение (вектор неизвестных x):")
	printVector(w, r.Result.X)

	fmt.Fprintf(w, "\nКоличество итераций: %d\n", r.Result.Iterations)

	fmt.Fprintln(w, "\nВектор погрешностей |x_i^(k) - x_i^(k-1)|:")
	printVector(w, r.Result.LastDelta)

	// Печать вектора невязки полезна для самопроверки (хотя для итерационных в ТЗ не обязательно).
	residual := computeResidual(r.A, r.Result.X, r.B)
	fmt.Fprintln(w, "\nВектор невязки r = Ax - b (для самопроверки):")
	printVector(w, residual)

	fmt.Fprintln(w, "================================================")
}

func printMatrix(w io.Writer, m [][]float64) {
	for _, row := range m {
		fmt.Fprint(w, "[ ")
		for _, v := range row {
			fmt.Fprintf(w, "%12.6f ", v)
		}
		fmt.Fprintln(w, "]")
	}
}

func printVector(w io.Writer, v []float64) {
	fmt.Fprint(w, "[ ")
	for _, x := range v {
		fmt.Fprintf(w, "%12.6f ", x)
	}
	fmt.Fprintln(w, "]")
}

func computeResidual(A [][]float64, x, b []float64) []float64 {
	n := len(A)
	r := make([]float64, n)
	for i := 0; i < n; i++ {
		var s float64
		for j := 0; j < n; j++ {
			s += A[i][j] * x[j]
		}
		r[i] = s - b[i]
		// Нормализуем "шум" типа 1e-16 для красивого вывода
		if math.Abs(r[i]) < 1e-15 {
			r[i] = 0
		}
	}
	return r
}