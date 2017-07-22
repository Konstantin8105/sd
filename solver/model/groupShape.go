package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/shape"
)

type shapeGroup struct {
	shape       shape.Shape
	beamIndexes []element.ElementIndex
}