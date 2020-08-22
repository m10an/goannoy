package goannoy

import (
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

const minDistanceDiff float64 = 1e-5

func assertPanic(t *testing.T, shouldPanic func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
	}()
	shouldPanic()
}

func TestWrongUsage(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)

	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})
	assertPanic(t, func() { index.Save("go_test.ann", false) })

	index.Build(10)
	assertPanic(t, func() { index.Build(10) })
	index.Save("go_test.ann", false)

	index, _ = NewAnnoyIndex(Angular, 3)
	index.Load("go_test.ann", false)
	assertPanic(t, func() { index.Unbuild() })
	assertPanic(t, func() { index.AddItem(0, []float32{0, 0, 1}) })
	assertPanic(t, func() { index.Build(10) })

	if err := os.Remove("go_test.ann"); err != nil {
		t.Error(err.Error())
	}
}

func TestFileHandling(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)

	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})

	index.Build(10)
	nItems := index.GetNItems()

	index.Save("go_test.ann", false)
	index, _ = NewAnnoyIndex(Angular, 3)
	index.Load("go_test.ann", false)

	if nLoaded := index.GetNItems(); nItems != nLoaded {
		t.Errorf("Number of indexed items mismatch: expected %d, loaded %d", nItems, nLoaded)
	}

	if err := os.Remove("go_test.ann"); err != nil {
		t.Error(err.Error())
	}

	index.Save("go_test2.ann", false)
	index, _ = NewAnnoyIndex(Angular, 3)

	index.Load("go_test2.ann", false)

	if err := os.Remove("go_test2.ann"); err != nil {
		t.Error(err.Error())
	}

	index.Save("go_test3.ann", true)
	index, _ = NewAnnoyIndex(Angular, 3)

	index.Load("go_test3.ann", true)

	if err := os.Remove("go_test3.ann"); err != nil {
		t.Error(err.Error())
	}
}

func TestOnDiskBuild(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)
	index.OnDiskBuild("go_test.ann")

	if _, err := os.Stat("go_test.ann"); err != nil {
		t.Error(err.Error())
	}

	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})
	index.Build(10)

	index.Unload()
	index.Load("go_test.ann", false)

	if !reflect.DeepEqual([]int32{2, 1, 0}, index.GetNNsByVector([]float32{3, 2, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNNsByVector([]float32{1, 2, 3}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{2, 0, 1}, index.GetNNsByVector([]float32{2, 0, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if err := os.Remove("go_test.ann"); err != nil {
		t.Error(err.Error())
	}
}

func TestGetNnsByVector(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)
	index.AddItem(0, []float32{0, 0, 1})
	index.AddItem(1, []float32{0, 1, 0})
	index.AddItem(2, []float32{1, 0, 0})
	index.Build(10)

	if !reflect.DeepEqual([]int32{2, 1, 0}, index.GetNNsByVector([]float32{3, 2, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNNsByVector([]float32{1, 2, 3}, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{2, 0, 1}, index.GetNNsByVector([]float32{2, 0, 1}, 3, -1)) {
		t.Error("Wrong nns order!")
	}
}

func TestGetNnsByItem(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)
	index.AddItem(0, []float32{2, 1, 0})
	index.AddItem(1, []float32{1, 2, 0})
	index.AddItem(2, []float32{0, 0, 1})
	index.Build(10)

	if !reflect.DeepEqual([]int32{0, 1, 2}, index.GetNNsByItem(0, 3, -1)) {
		t.Error("Wrong nns order!")
	}

	if !reflect.DeepEqual([]int32{1, 0, 2}, index.GetNNsByItem(1, 3, -1)) {
		t.Error("Wrong nns order!")
	}
}

func TestGetItem(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 3)
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
}

func TestGetAngularDistance(t *testing.T) {
	index, _ := NewAnnoyIndex(Angular, 2)
	index.AddItem(0, []float32{0, 1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	d1 := math.Pow(2*(1.0-math.Pow(2, -0.5)), 0.5)
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
}

func TestGetManhattan(t *testing.T) {
	index, _ := NewAnnoyIndex(Manhattan, 2)
	index.AddItem(0, []float32{3, -1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	var d1 float64 = 4
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
}

func TestGetDotProductDistance(t *testing.T) {
	index, _ := NewAnnoyIndex(DotProduct, 2)
	index.AddItem(0, []float32{0, 1})
	index.AddItem(1, []float32{1, 1})
	index.Build(10)

	var d1 float64 = 1
	d2 := float64(index.GetDistance(0, 1))
	if diff := math.Abs(d1 - d2); diff > minDistanceDiff {
		t.Errorf("%f - %f (= %f) > %e", d1, d2, diff, minDistanceDiff)
	}
}

func TestLargeEuclideanIndex(t *testing.T) {
	index, _ := NewAnnoyIndex(Euclidean, 10)

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
		if !reflect.DeepEqual([]int32{int32(j), int32(j) + 1}, index.GetNNsByItem(j, 2, -1)) {
			t.Error("Wrong nns order!")
		}
		if !reflect.DeepEqual([]int32{int32(j) + 1, int32(j)}, index.GetNNsByItem(j+1, 2, -1)) {
			t.Error("Wrong nns order!")
		}
	}
}
