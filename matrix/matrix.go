package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows [][]int
	cols [][]int
}

func isInvalid(m [][]string) (bool, error) {
	// verify dimension
	dimension := 0
	for _, v := range m {
		if len(v) != dimension && dimension != 0 {
			return false, errors.New("invalid dimension")
		}
		dimension = len(v)
	}
	return true, nil
}

func New(s string) (*Matrix, error) {
	// remove space after \n, split by \n
	getRawRows := strings.Split(strings.ReplaceAll(s, "\n ", "\n"), "\n")

	// create a row structure
	rows := make([][]string, len(getRawRows))
	for i, r := range getRawRows {
		rows[i] = strings.Split(r, " ")
	}

	// validate input matrix
	if ok, err := isInvalid(rows); !ok {
		return &Matrix{}, err
	}

	m := Matrix{}
	m.rows = make([][]int, len(rows))
	for i, r := range rows {
		for _, v := range r {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Matrix{}, err
			}
			m.rows[i] = append(m.rows[i], n)
		}
	}

	m.cols = make([][]int, len(m.rows[0]))
	for i := 0; i < len(m.cols); i++ {
		for _, r := range m.rows {
			m.cols[i] = append(m.cols[i], r[i])
		}
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	duplicate := make([][]int, len(m.cols))
	for i := range m.cols {
		duplicate[i] = make([]int, len(m.cols[i]))
		copy(duplicate[i], m.cols[i])
	}
	return duplicate
}

func (m *Matrix) Rows() [][]int {
	duplicate := make([][]int, len(m.rows))
	for i := range m.rows {
		duplicate[i] = make([]int, len(m.rows[i]))
		copy(duplicate[i], m.rows[i])
	}
	return duplicate
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m.rows) || col >= len(m.cols) {
		return false
	}
	m.rows[row][col] = val
	m.cols[col][row] = val
	return true
}
