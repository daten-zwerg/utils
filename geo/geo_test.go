package geo

import "testing"

func TestBoundingBoxList(t *testing.T) {
	tables := []struct {
		b    BoundingBox
		want []float64
	}{
		{BoundingBox{
			TopLeft:     Location{},
			BottomRight: Location{},
		},
			[]float64{0, 0, 0, 0}},
		{BoundingBox{
			TopLeft:     Location{Latitude: 51.38215337765274, Longitude: 5.401567891267973},
			BottomRight: Location{Latitude: 51.15716817451093, Longitude: 7.30774808471395},
		},
			[]float64{51.38215337765274, 5.401567891267973, 51.15716817451093, 7.30774808471395}},
	}

	for _, table := range tables {
		got := table.b.List()
		if len(got) != len(table.want) {
			t.Errorf("Invalid length of created list, got: %d, want: %d.", len(got), len(table.want))
		}

		for i := 0; i < len(got); i++ {
			if got[i] != table.want[i] {
				t.Errorf("Invalid element with index %d in created list, got: %f, want: %f.", i, got[i], table.want[i])
			}
		}
	}
}

func TestBoundingBoxString(t *testing.T) {
	tables := []struct {
		b    BoundingBox
		sep  string
		want string
	}{
		{BoundingBox{
			TopLeft:     Location{},
			BottomRight: Location{},
		},
			"",
			"0000"},
		{BoundingBox{
			TopLeft:     Location{Latitude: 51.38215337765274, Longitude: 5.401567891267973},
			BottomRight: Location{Latitude: 51.15716817451093, Longitude: 7.30774808471395},
		},
			",",
			"51.38215337765274,5.401567891267973,51.15716817451093,7.30774808471395"},
	}

	for _, table := range tables {
		got := table.b.String(table.sep)

		if got != table.want {
			t.Errorf("Created string does not match, got: %s, want: %s.", got, table.want)
		}
	}
}
