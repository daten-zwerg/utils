package geo

import "strconv"

type Location struct {
	Latitude  float64
	Longitude float64
}

type BoundingBox struct {
	TopLeft     Location
	BottomRight Location
}

func (bbox *BoundingBox) List() []float64 {
	return []float64{
		bbox.TopLeft.Latitude,
		bbox.TopLeft.Longitude,
		bbox.BottomRight.Latitude,
		bbox.BottomRight.Longitude,
	}
}

func (bbox *BoundingBox) String(sep string) string {
	var res string
	values := bbox.List()
	for i, v := range values {
		res += strconv.FormatFloat(v, 'f', -1, 64)
		if i+1 < len(values) {
			res += sep
		}
	}
	return res
}
