package service

import "log"

// InputData — входные данные задачи Ax=b.
type InputData struct {
	N   int
	A   [][]float64
	B   []float64
	Eps float64
}

// SeidelConfig — конфигурация метода Гаусса–Зейделя.
type SeidelConfig struct {
	A        [][]float64
	B        []float64
	Eps      float64
	MaxIters int
	Logger   *log.Logger
}

// SeidelResult — результат работы метода.
type SeidelResult struct {
	X          []float64   // найденный вектор решения
	Iterations int         // число итераций
	LastDelta  []float64   // вектор погрешностей |x_i^(k)-x_i^(k-1)|
	History    [][]float64 // опционально: история итераций (можно отключить, если не нужна)
}

// Report — структура для печати результата.
type Report struct {
	Input     InputData
	A         [][]float64
	B         []float64
	C         [][]float64
	D         []float64
	CNormInf  float64
	Reordered bool
	RowOrder  []int
	Result    SeidelResult
}