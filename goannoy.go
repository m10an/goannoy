package goannoy

// #include "goannoy.h"
import "C"
import "unsafe"

type AnnoyIndex struct {
	index     *C.AnnoyIndex
	nFeatures int
}

func NewAnnoyIndexAngular(f int) *AnnoyIndex {
	return &AnnoyIndex{C.CreateAnnoyIndexAngular(C.int(f)), f}
}

func NewAnnoyIndexEuclidean(f int) *AnnoyIndex {
	return &AnnoyIndex{C.CreateAnnoyIndexEuclidean(C.int(f)), f}
}

func NewAnnoyIndexManhattan(f int) *AnnoyIndex {
	return &AnnoyIndex{C.CreateAnnoyIndexManhattan(C.int(f)), f}
}

func NewAnnoyIndexDotProduct(f int) *AnnoyIndex {
	return &AnnoyIndex{C.CreateAnnoyIndexDotProduct(C.int(f)), f}
}

func DeleteAnnoyIndex(a *AnnoyIndex) {
	C.DeleteAnnoyIndex(a.index)
}

func (a *AnnoyIndex) AddItem(item int, w []float32) {
	errMsg := new(*C.char)
	if !bool(C.AddItem(a.index, C.GoInt(item), (*C.float)(&w[0]), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) GetNItems() int {
	return int(C.GetNItems(a.index))
}

func (a *AnnoyIndex) Build(nTrees int) {
	errMsg := new(*C.char)
	if !bool(C.Build(a.index, C.int(nTrees), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) Unbuild() {
	errMsg := new(*C.char)
	if !bool(C.Unbuild(a.index, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) Save(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.Save(a.index, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) Unload() {
	C.Unload(a.index)
}

func (a *AnnoyIndex) Load(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.Load(a.index, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) GetDistance(i, j int) float32 {
	return float32(C.GetDistance(a.index, C.int(i), C.int(j)))
}

func (a *AnnoyIndex) GetNNsByItem(item, n, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.GetNNsByItem(a.index, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (a *AnnoyIndex) GetNNsByItemWithDistances(item, n, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.GetNNsByItemWithDistances(a.index, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (a *AnnoyIndex) GetNNsByVector(w []float32, n, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.GetNNsByVector(a.index, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (a *AnnoyIndex) GetNNsByVectorWithDistances(w []float32, n, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.GetNNsByVectorWithDistances(a.index, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (a *AnnoyIndex) GetItem(item int) []float32 {
	vector := make([]float32, a.nFeatures)
	C.GetItem(a.index, C.int(item), (*C.float)(&vector[0]))
	return vector
}

func (a *AnnoyIndex) OnDiskBuild(filename string) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.OnDiskBuild(a.index, chars, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (a *AnnoyIndex) Verbose(v bool) {
	C.Verbose(a.index, C.bool(v))
}
