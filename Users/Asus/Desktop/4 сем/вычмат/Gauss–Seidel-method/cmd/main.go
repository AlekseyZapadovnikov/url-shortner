package main

import (
	"log"
	"os"

	"github.com/AlekseyZapadovnikov/Gaus-Seidel-method/service"
)

func main() {
	logger := log.New(os.Stdout, "[seidel] ", log.LstdFlags|log.Lshortfile)

	logger.Println("Application started")

	input, err := service.ReadInputInteractive(os.Stdin, os.Stdout)
	if err != nil {
		logger.Fatalf("failed to read input: %v", err)
	}

	logger.Printf("Input read: n=%d, eps=%g", input.N, input.Eps)

	A := service.CopyMatrix(input.A)
	b := service.CopyVector(input.B)

	// Проверка/достижение диагонального преобладания (перестановкой строк)
	diagBefore := service.HasStrictDiagonalDominance(A)
	logger.Printf("Diagonal dominance before reorder: %v", diagBefore)

	reordered, order, err := service.EnsureDiagonalDominanceByRowPermutation(A)
	if err != nil {
		logger.Fatalf("cannot ensure diagonal dominance: %v", err)
	}
	if reordered {
		b = service.ReorderVectorByRowOrder(b, order)
		logger.Printf("Rows reordered to achieve diagonal dominance. order=%v", order)
	} else {
		logger.Println("Row reorder not required")
	}

	diagAfter := service.HasStrictDiagonalDominance(A)
	logger.Printf("Diagonal dominance after reorder: %v", diagAfter)

	// Норма матрицы итерации C (берем ||C||_∞ = max sum по строкам)
	cNorm, C, d, err := service.BuildIterationFormAndNorm(A, b)
	if err != nil {
		logger.Fatalf("failed to build iteration form: %v", err)
	}
	logger.Printf("Iteration matrix norm ||C||_inf = %.10f", cNorm)

	// Сам метод Гаусса–Зейделя
	result, err := service.SolveGaussSeidel(service.SeidelConfig{
		A:        A,
		B:        b,
		Eps:      input.Eps,
		MaxIters: 10000,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatalf("Gauss-Seidel failed: %v", err)
	}

	service.PrintReport(os.Stdout, service.Report{
		Input:      input,
		A:          A,
		B:          b,
		C:          C,
		D:          d,
		CNormInf:   cNorm,
		Reordered:  reordered,
		RowOrder:   order,
		Result:     result,
	})
	logger.Println("Application finished successfully")
}