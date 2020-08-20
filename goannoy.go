package goannoy

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include "gotypes.h"

void* create_annidx_angular(int);
void* create_annidx_euclidean(int);
void* create_annidx_manhattan(int);
void* create_annidx_dot_product(int);
void free_annidx(void *);
void add_item(void *, intgo_t, float_slice_t);
void build(void *, int);
bool save(void *, const char *, bool);
void unload(void *);
bool load(void *, const char *, bool);
float get_distance(void *, int, int);
int get_nns_by_item(void *, int, int, int, int32_t *, float *);
int get_nns_by_vector(void *, const float *w, int, int, int32_t *, float *);
int get_n_items(void *);
void verbose(void *, bool);
void get_item(void *, int, float *);
bool on_disk_build(void *, const char *);
*/
import "C"
import "unsafe"

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
	C.add_item(i.self, C.GoInt(item), *(*C.GoFloatSlice)(unsafe.Pointer(&w)))
}

func (i *Index) GetNItems() int {
	return int(C.get_n_items(i.self))
}

func (i *Index) Build(nTrees int) {
	C.build(i.self, C.int(nTrees))
}

func (i *Index) Save(filename string, prefault bool) bool {
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	return bool(C.save(i.self, chars, C.bool(prefault)))
}

func (i *Index) Unload() {
	C.unload(i.self)
}

func (i *Index) Load(filename string, prefault bool) bool {
	chars := C.CString(filename)
	defer C.free(unsafe.Pointer(chars))
	return bool(C.load(i.self, chars, C.bool(prefault)))
}

func (i *Index) GetDistance(firstItem int, secondItem int) float32 {
	return float32(C.get_distance(i.self, C.int(firstItem), C.int(secondItem)))
}

func (i *Index) GetNnsByItem(item int, n int, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_item(i.self,
		C.int(item),
		C.int(n),
		C.int(kSearch),
		(*C.int)(&result[0]),
		nil)
	return result[:found]
}

func (i *Index) GetNnsByItemWithDistances(item int, n int, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_item(i.self,
		C.int(item),
		C.int(n),
		C.int(kSearch),
		(*C.int)(&result[0]),
		(*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (i *Index) GetNnsByVector(w []float32, n int, kSearch int) []int32 {
	result := make([]int32, n)
	found := C.get_nns_by_vector(i.self,
		(*C.float)(&w[0]),
		C.int(n),
		C.int(kSearch),
		(*C.int)(&result[0]),
		nil)
	return result[:found]
}

func (i *Index) GetNnsByVectorWithDistances(w []float32, n int, kSearch int) ([]int32, []float32) {
	result := make([]int32, n)
	distances := make([]float32, n)
	found := C.get_nns_by_vector(i.self,
		(*C.float)(&w[0]),
		C.int(n),
		C.int(kSearch),
		(*C.int)(&result[0]),
		(*C.float)(&distances[0]))
	return result[:found], distances[:found]
}

func (i *Index) GetItem(item int) []float32 {
	vector := make([]float32, i.nFeatures)
	C.get_item(i.self, C.int(item), (*C.float)(&vector[0]))
	return vector
}

func (i *Index) OnDiskBuild(filename string) {
	chars := C.CString(filename)
	C.on_disk_build(i.self, chars)
	C.free(unsafe.Pointer(chars))
}

func (i *Index) Verbose(v bool) {
	C.verbose(i.self, C.bool(v))
}
