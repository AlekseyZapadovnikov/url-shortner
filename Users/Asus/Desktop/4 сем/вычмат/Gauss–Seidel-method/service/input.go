package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadInputInteractive — выбор способа ввода (клавиатура/файл).
// Формат файла:
// n
// a11 a12 ... a1n
// ...
// an1 an2 ... ann
// b1 b2 ... bn
// eps
func ReadInputInteractive(in io.Reader, out io.Writer) (InputData, error) {
	reader := bufio.NewReader(in)

	fmt.Fprintln(out, "Выберите способ ввода:")
	fmt.Fprintln(out, "1 - с клавиатуры")
	fmt.Fprintln(out, "2 - из файла")
	fmt.Fprint(out, "> ")

	modeLine, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return InputData{}, fmt.Errorf("read mode: %w", err)
	}
	modeLine = strings.TrimSpace(modeLine)

	switch modeLine {
	case "1":
		return readFromKeyboard(reader, out)
	case "2":
		fmt.Fprint(out, "Введите путь к файлу: ")
		pathLine, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return InputData{}, fmt.Errorf("read file path: %w", err)
		}
		path := strings.TrimSpace(pathLine)
		return ReadInputFromFile(path)
	default:
		return InputData{}, fmt.Errorf("unknown input mode: %q", modeLine)
	}
}

func readFromKeyboard(reader *bufio.Reader, out io.Writer) (InputData, error) {
	var n int
	fmt.Fprint(out, "Введите размерность n (n <= 20): ")
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return InputData{}, fmt.Errorf("read n: %w", err)
	}
	if n <= 0 || n > 20 {
		return InputData{}, fmt.Errorf("invalid n=%d (must be 1..20)", n)
	}

	A := make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		fmt.Fprintf(out, "Введите строку матрицы A[%d] (%d чисел): ", i+1, n)
		for j := 0; j < n; j++ {
			if _, err := fmt.Fscan(reader, &A[i][j]); err != nil {
				return InputData{}, fmt.Errorf("read A[%d][%d]: %w", i, j, err)
			}
		}
	}

	b := make([]float64, n)
	fmt.Fprintf(out, "Введите вектор b (%d чисел): ", n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(reader, &b[i]); err != nil {
			return InputData{}, fmt.Errorf("read b[%d]: %w", i, err)
		}
	}

	var eps float64
	fmt.Fprint(out, "Введите точность eps: ")
	if _, err := fmt.Fscan(reader, &eps); err != nil {
		return InputData{}, fmt.Errorf("read eps: %w", err)
	}
	if eps <= 0 {
		return InputData{}, fmt.Errorf("eps must be > 0")
	}

	return InputData{
		N:   n,
		A:   A,
		B:   b,
		Eps: eps,
	}, nil
}

// ReadInputFromFile — чтение входных данных из файла.
func ReadInputFromFile(path string) (InputData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return InputData{}, fmt.Errorf("read file %q: %w", path, err)
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return InputData{}, fmt.Errorf("empty file")
	}

	pos := 0
	nextInt := func() (int, error) {
		if pos >= len(fields) {
			return 0, fmt.Errorf("unexpected EOF while reading int")
		}
		v, err := strconv.Atoi(fields[pos])
		pos++
		if err != nil {
			return 0, fmt.Errorf("parse int %q: %w", fields[pos-1], err)
		}
		return v, nil
	}

	nextFloat := func() (float64, error) {
		if pos >= len(fields) {
			return 0, fmt.Errorf("unexpected EOF while reading float")
		}
		v, err := strconv.ParseFloat(fields[pos], 64)
		pos++
		if err != nil {
			return 0, fmt.Errorf("parse float %q: %w", fields[pos-1], err)
		}
		return v, nil
	}

	n, err := nextInt()
	if err != nil {
		return InputData{}, fmt.Errorf("read n: %w", err)
	}
	if n <= 0 || n > 20 {
		return InputData{}, fmt.Errorf("invalid n=%d (must be 1..20)", n)
	}

	A := make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			A[i][j], err = nextFloat()
			if err != nil {
				return InputData{}, fmt.Errorf("read A[%d][%d]: %w", i, j, err)
			}
		}
	}

	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i], err = nextFloat()
		if err != nil {
			return InputData{}, fmt.Errorf("read b[%d]: %w", i, err)
		}
	}

	eps, err := nextFloat()
	if err != nil {
		return InputData{}, fmt.Errorf("read eps: %w", err)
	}
	if eps <= 0 {
		return InputData{}, fmt.Errorf("eps must be > 0")
	}

	return InputData{
		N:   n,
		A:   A,
		B:   b,
		Eps: eps,
	}, nil
}

// CopyMatrix — глубокая копия матрицы.
func CopyMatrix(src [][]float64) [][]float64 {
	dst := make([][]float64, len(src))
	for i := range src {
		dst[i] = append([]float64(nil), src[i]...)
	}
	return dst
}

// CopyVector — копия вектора.
func CopyVector(src []float64) []float64 {
	return append([]float64(nil), src...)
}