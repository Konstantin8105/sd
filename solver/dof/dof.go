package dof

import (
	"fmt"
	"sort"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/utils"
)

// Dim - dimension unit
type Dim int

// Type of dimension
const (
	Dim2d Dim = 3 // 3 degree of freedom for point in 2d. Dx, Dy, M
	Dim3d     = 6 // 6 degree of freedom for point in 3d. Dx, Dy, Dz, Mx, My, Mz
)

// AxeNumber - axe of number
type AxeNumber int

// DoF - degree of freedom
type DoF struct {
	dofArray  []int
	dimension Dim
}

// ForElements - add new beam
func ForElements(elements []element.Elementer, dim Dim) (d DoF) {
	var array []int
	for _, e := range elements {
		switch e.(type) {
		case element.Beam:
			beam := e.(element.Beam)
			for i := range beam.PointIndexes {
				array = append(array, int(beam.PointIndexes[i]))
			}
		default:
			panic("")
		}
	}
	utils.UniqueInt(&array)
	d.dofArray = array
	d.dimension = dim
	return d
}

// GetDoF - get degree of freedom for point index
func (d *DoF) GetDoF(index point.Index) []AxeNumber {
	if d.dimension == Dim2d {
		axes := make([]AxeNumber, int(d.dimension), int(d.dimension))
		number := d.found(index)
		for i := 0; i < int(d.dimension); i++ {
			axes[i] = AxeNumber(i + number*int(Dim2d))
		}
		return axes
	}
	panic("Please add algorithm")
}

func (d *DoF) found(index point.Index) int {
	i := sort.Search(len(d.dofArray), func(i int) bool { return d.dofArray[i] >= int(index) })
	if i >= 0 && i < len(d.dofArray) && d.dofArray[i] == int(index) {
		// index is present at array[i]
		return i
	}
	// index is not present in array,
	// but i is the index where it would be inserted.
	panic("Not correct binary searching")
}

// ConvertToInt - convert to int
func ConvertToInt(axes []AxeNumber) []int {
	result := make([]int, len(axes), len(axes))
	for i := 0; i < len(axes); i++ {
		result[i] = int(axes[i])
	}
	return result
}

// ConvertToAxe - convert to axe
func ConvertToAxe(ins []int) []AxeNumber {
	result := make([]AxeNumber, len(ins), len(ins))
	for i := 0; i < len(ins); i++ {
		result[i] = AxeNumber(ins[i])
	}
	return result
}

// RemoveIndexes - remove indexex for axeNumber slice
// without reallocation matrix
func RemoveIndexes(a *[]AxeNumber, indexes ...int) {
	if len(indexes) == 0 {
		return
	}
	// sorting indexes for optimization of algoritm
	sort.Ints(indexes)
	// global checking indexes
	if indexes[len(indexes)-1] >= len(*a) {
		panic(fmt.Errorf("indexes is outside of matrix. Indexes = %v", indexes))
	}
	// modify values
	positionIndex := 0
	newPositionInSlice := 0
	for i := 0; i < len(*a); i++ {
		if positionIndex != len(indexes) && i == indexes[positionIndex] {
			positionIndex++
			continue
		}
		(*a)[newPositionInSlice] = (*a)[i]
		newPositionInSlice++
	}

	(*a) = (*a)[0 : len(*a)-len(indexes)]
}