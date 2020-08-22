package goannoy

// #include "goannoy.h"
import "C"
import (
	"unsafe"
)

type Index struct {
	c         unsafe.Pointer
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

func DeleteAnnoyIndex(idx Index) {
	C.free_annidx(idx.c)
}

func (idx *Index) AddItem(item int, w []float32) {
	errMsg := new(*C.char)
	if !bool(C.add_item(idx.c, C.GoInt(item), (*C.float)(&w[0]), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) GetNItems() int {
	return int(C.get_n_items(idx.c))
}

func (idx *Index) Build(nTrees int) {
	errMsg := new(*C.char)
	if !bool(C.build(idx.c, C.int(nTrees), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) Unbuild() {
	errMsg := new(*C.char)
	if !bool(C.unbuild(idx.c, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) Save(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.save(idx.c, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) Unload() {
	C.unload(idx.c)
}

func (idx *Index) Load(filename string, prefault bool) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.load(idx.c, chars, C.bool(prefault), errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) GetDistance(i, j int) float32 {
	return float32(C.get_distance(idx.c, C.int(i), C.int(j)))
}

func (idx *Index) GetNnsByItem(item, n, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_item(idx.c, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (idx *Index) GetNnsByItemWithDistances(item, n, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_item_with_dists(idx.c, C.int(item), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (idx *Index) GetNnsByVector(w []float32, n, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_vector(idx.c, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]))
	return result[:found]
}

func (idx *Index) GetNnsByVectorWithDistances(w []float32, n, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_vector_with_dists(idx.c, (*C.float)(&w[0]), C.int(n), C.int(kSearch), (*C.GoInt32)(&result[0]), (*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (idx *Index) GetItem(item int) []float32 {
	vector := make([]float32, idx.nFeatures)
	C.get_item(idx.c, C.int(item), (*C.float)(&vector[0]))
	return vector
}

func (idx *Index) OnDiskBuild(filename string) {
	errMsg := new(*C.char)
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	if !bool(C.on_disk_build(idx.c, chars, errMsg)) {
		defer C.free(unsafe.Pointer(*errMsg))
		panic(C.GoString(*errMsg))
	}
}

func (idx *Index) Verbose(v bool) {
	C.verbose(idx.c, C.bool(v))
}
