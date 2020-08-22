package goannoy

// #include "goannoy.h"
import "C"
import (
	"unsafe"
)

type Index struct {
	self      unsafe.Pointer
	nFeatures int
}

func NewAnnoyIndexAngular(f int) Index {
	return Index{C.create_annidx_angular(C.int(f)), f}
}

func NewAnnoyIndexEuclidean(f int) Index {
	return Index{C.create_annidx_euclidean(C.int(f)), f}
}

func NewAnnoyIndexManhattan(f int) Index {
	return Index{C.create_annidx_manhattan(C.int(f)), f}
}

func NewAnnoyIndexDotProduct(f int) Index {
	return Index{C.create_annidx_dot_product(C.int(f)), f}
}

func DeleteAnnoyIndex(i Index) {
	C.free_annidx(i.self)
}

func (i *Index) AddItem(item int, w []float32) {
	errMsg := new(*C.char)
	if !bool(C.add_item(i.self, C.GoInt(item), (*C.float)(&w[0]), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) GetNItems() int {
	return int(C.get_n_items(i.self))
}

func (i *Index) Build(nTrees int) {
	errMsg := new(*C.char)
	if !bool(C.build(i.self, C.int(nTrees), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) Unbuild() {
	errMsg := new(*C.char)
	if !bool(C.unbuild(i.self, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) Save(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.save(i.self, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) Unload() {
	C.unload(i.self)
}

func (i *Index) Load(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.load(i.self, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) GetDistance(firstItem int, secondItem int) float32 {
	return float32(C.get_distance(i.self, C.int(firstItem), C.int(secondItem)))
}

func (i *Index) GetNnsByItem(item int, n int, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_item(i.self, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (i *Index) GetNnsByItemWithDistances(item int, n int, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_item_with_dists(i.self, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (i *Index) GetNnsByVector(w []float32, n int, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_vector(i.self, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (i *Index) GetNnsByVectorWithDistances(w []float32, n int, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_vector_with_dists(i.self, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (i *Index) GetItem(item int) []float32 {
	vector := make([]float32, i.nFeatures)
	C.get_item(i.self, C.int(item), (*C.float)(&vector[0]))
	return vector
}

func (i *Index) OnDiskBuild(filename string) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.on_disk_build(i.self, chars, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (i *Index) Verbose(v bool) {
	C.verbose(i.self, C.bool(v))
}
