package lof

//////////////////////////////////////////////////////////////////////////////
//
//  Interface for anything that must be treated as a point in a sample space.
//
//////////////////////////////////////////////////////////////////////////////

type ISample interface {
	GetID() int
	SetID(int)
	GetPoint() []float64
	SetPoint([]float64)
}

//////////////////////////////////////////////////////////////////////////////
//
//  This type is only used with GetSamplesFromFloat64s().
//
//////////////////////////////////////////////////////////////////////////////

type BasicSample struct {
	ID       int
	Point    []float64
	Distance float64
}

// NewBasicSample : Constructor for BasicSample type.
func NewBasicSample(id int, point []float64) *BasicSample {

	return &BasicSample{Point: point, ID: id}
}

// GetSamplesFromFloat64s : Given a slice of float64 slices, build BasicSamples
// treating each float64 slice as a point in sample space.
func GetSamplesFromFloat64s(points [][]float64) []ISample {
	bSpl := []ISample{}
	for idx, pt := range points {
		bSpl = append(bSpl, NewBasicSample(idx, pt))
	}
	return bSpl
}

// GetID : Satisfies ISample interface.
func (bs *BasicSample) GetID() int {
	return bs.ID
}

// SetID : Satisfies ISample interface.
func (bs *BasicSample) SetID(id int) {
	bs.ID = id
}

// GetPoint : Satisfies ISample interface.
func (bs *BasicSample) GetPoint() []float64 {
	return bs.Point
}

// SetPoint : Satisfies ISample interface.
func (bs *BasicSample) SetPoint(point []float64) {
	bs.Point = point
}
