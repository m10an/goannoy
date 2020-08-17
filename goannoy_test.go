package goannoy

import (
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

const minDistanceDiff float64 = 1e-5

func TestFileHandling(t *testing.T) {
	index := NewAnnoyIndexAngular(3)

	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})

	index.Build(10)
	nItems := index.GetNItems()

	if !index.Save("go_test.ann", false) {
		t.Error("Failed to create file without prefault")
	}

	DeleteAnnoyIndex(index)
	index = NewAnnoyIndexAngular(3)

	if !index.Load("go_test.ann", false) {
		t.Error("Failed to load file without prefault")
	}

	if nLoaded := index.GetNItems(); nItems != nLoaded {
		t.Errorf("Number of indexed items mismatch: expected %d, loaded %d", nItems, nLoaded)
	}

	if err := os.Remove("go_test.ann"); err != nil {
		t.Error(err.Error())
	}

	if !index.Save("go_test2.ann", false) {
		t.Error("Failed to create file without prefault")
	}

	DeleteAnnoyIndex(index)
	index = NewAnnoyIndexAngular(3)

	if !index.Load("go_test2.ann", false) {
		t.Error("Failed to load file without prefault")
	}

	if err := os.Remove("go_test2.ann"); err != nil {
		t.Error(err.Error())
	}

	if !index.Save("go_test3.ann", true) {
		t.Error("Failed to create file with prefault")
	}

	DeleteAnnoyIndex(index)
	index = NewAnnoyIndexAngular(3)

	if !index.Load("go_test3.ann", true) {
		t.Error("Failed to load file allowing prefault")
	}

	DeleteAnnoyIndex(index)

	if err := os.Remove("go_test3.ann"); err != nil {
		t.Error(err.Error())
	}
}

func TestOnDiskBuild(t *testing.T) {
	index := NewAnnoyIndexAngular(3)
	index.OnDiskBuild("go_test.ann")

	if _, err := os.Stat("go_test.ann"); err != nil {
		t.Error(err.Error())
	}

	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})
	index.Build(10)

	index.Unload()
	if !index.Load("go_test.ann", false) {
		t.Error("Failed to load file")
	}

	if !reflect.DeepEqual([]int32{2, 1, 0}, index.GetNnsByVector([]float32{3, 2, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNnsByVector([]float32{1, 2, 3}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{2, 0, 1}, index.GetNnsByVector([]float32{2, 0, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	DeleteAnnoyIndex(index)

	if err := os.Remove("go_test.ann"); err != nil {
		t.Error(err.Error())
	}
}

func TestGetNnsByVector(t *testing.T) {
	index := NewAnnoyIndexAngular(3)
	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})
	index.Build(10)

	if !reflect.DeepEqual([]int32{2, 1, 0}, index.GetNnsByVector([]float32{3, 2, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNnsByVector([]float32{1, 2, 3}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{2, 0, 1}, index.GetNnsByVector([]float32{2, 0, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	DeleteAnnoyIndex(index)
}

func TestGetNnsByItem(t *testing.T) {
	index := NewAnnoyIndexAngular(3)
	index.AddItem(0, []float32{2, 1, 0})
	index.AddItem(1, []float32{1, 2, 0})
	index.AddItem(2, []float32{0, 0, 1})
	index.Build(10)

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNnsByItem(0, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{1, 0, 2}, index.GetNnsByItem(1, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	DeleteAnnoyIndex(index)
}

func TestGetItem(t *testing.T) {
	index := NewAnnoyIndexAngular(3)
	index.AddItem(0, []float32{2, 1, 0})
	index.AddItem(1, []float32{1, 2, 0})
	index.AddItem(2, []float32{0, 0, 1})
	index.Build(10)

	if !reflect.DeepEqual([]float32{2, 1, 0}, index.GetItem(0)) {
		t.Error("Wrong item got!")
	}

	if !reflect.DeepEqual([]float32{1, 2, 0}, index.GetItem(1)) {
		t.Error("Wrong item got!")
	}

	if !reflect.DeepEqual([]float32{0, 0, 1}, index.GetItem(2)) {
		t.Error("Wrong item got!")
	}

	DeleteAnnoyIndex(index)
}

func TestGetAngularDistance(t *testing.T) {
	index := NewAnnoyIndexAngular(2)
	index.AddItem(0, []float32{0, 1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	d1 := math.Pow(2*(1.0-math.Pow(2, -0.5)), 0.5)
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
	DeleteAnnoyIndex(index)
}

func TestGetManhattan(t *testing.T) {
	index := NewAnnoyIndexManhattan(2)
	index.AddItem(0, []float32{3, -1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	var d1 float64 = 4
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
	DeleteAnnoyIndex(index)
}

func TestGetDotProductDistance(t *testing.T) {
	index := NewAnnoyIndexDotProduct(2)
	index.AddItem(0, []float32{0, 1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	var d1 float64 = 1
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
	DeleteAnnoyIndex(index)
}

func TestLargeEuclideanIndex(t *testing.T) {
	index := NewAnnoyIndexEuclidean(10)

	for j := 0; j < 10000; j += 2 {
		p := make([]float32, 0, 10)
		for i := 0; i < 10; i++ {
			p = append(p, rand.Float32())
		}
		x := make([]float32, 0, 10)
		for i := 0; i < 10; i++ {
			x = append(x, 1+p[i]+rand.Float32()*1e-2)
		}
		y := make([]float32, 0, 10)
		for i := 0; i < 10; i++ {
			y = append(y, 1+p[i]+rand.Float32()*1e-2)
		}
		index.AddItem(j, x)
		index.AddItem(j+1, y)
	}
	index.Build(10)
	for j := 0; j < 10000; j += 2 {
		if !reflect.DeepEqual([]int32{int32(j), int32(j) + 1}, index.GetNnsByItem(j, 2, -1)) {
			t.Error("Wrong nns order!")
		}
		if !reflect.DeepEqual([]int32{int32(j) + 1, int32(j)}, index.GetNnsByItem(j+1, 2, -1)) {
			t.Error("Wrong nns order!")
		}
	}
	DeleteAnnoyIndex(index)
}