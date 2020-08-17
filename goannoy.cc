#include "annoylib.h"
#include "kissrandom.h"

using namespace Annoy;

// Annoy Class
template <class D>
using AnnoyC = AnnoyIndex<int32_t, float, D, Kiss64Random, AnnoyIndexSingleThreadedBuildPolicy>;

// Annoy Interface
using AnnoyI = AnnoyIndexInterface<int32_t, float>;

extern "C" {

  AnnoyC<Angular>* create_annidx_angular(int f) {
    return new AnnoyC<Angular>(f);
  }

  AnnoyC<Euclidean>* create_annidx_euclidean(int f) {
    return new AnnoyC<Euclidean>(f);
  }

  AnnoyC<Manhattan>* create_annidx_manhattan(int f) {
    return new AnnoyC<Manhattan>(f);
  }

  AnnoyC<DotProduct>* create_annidx_dot_product(int f) {
    return new AnnoyC<DotProduct>(f);
  }

  void free_annidx(AnnoyI *ptr) {
    delete ptr;
  }

  void add_item(AnnoyI *ptr, int item, const float *w) {
    ptr->add_item(item, w);
  }

  void build(AnnoyI *ptr, int q) {
    ptr->build(q, 1);
  }

  bool save(AnnoyI *ptr, const char *filename, bool prefault) {
    return ptr->save(filename, prefault);
  }

  void unload(AnnoyI *ptr) {
    ptr->unload();
  }

  bool load(AnnoyI *ptr, const char *filename, bool prefault) {
    return ptr->load(filename, prefault);
  }

  float get_distance(AnnoyI *ptr, int i, int j) {
    return ptr->get_distance(i, j);
  }

  int _results_to_arrays(vector<int32_t> *result_vec, vector<float> *distances_vec, int32_t *result, float *distances){
    int size = result_vec->size();
    for (int i = 0; i < size; ++i){
      result[i] = (*result_vec)[i];
      if (distances)
        distances[i] = (*distances_vec)[i];
    }
    return size;
  }

  int get_nns_by_item(AnnoyI *ptr, int item, int n, int search_k, int32_t *result, float *distances) {
    vector<int32_t> result_vec;
    vector<float> distances_vec;
    ptr->get_nns_by_item(item, n, search_k, &result_vec, distances ? &distances_vec : NULL);
    return _results_to_arrays(&result_vec, &distances_vec, result, distances);
  }

  int get_nns_by_vector(AnnoyI *ptr, const float *w, int n, int search_k, int32_t *result, float *distances) {
    vector<int32_t> result_vec;
    vector<float> distances_vec;
    ptr->get_nns_by_vector(w, n, search_k, &result_vec, distances ? &distances_vec : NULL);
    return _results_to_arrays(&result_vec, &distances_vec, result, distances);
  }

  int get_n_items(AnnoyI *ptr) {
    return (int)ptr->get_n_items();
  }

  void verbose(AnnoyI *ptr, bool v) {
    ptr->verbose(v);
  }

  void get_item(AnnoyI *ptr, int item, float *v) {
    ptr->get_item(item, v);
  }

  bool on_disk_build(AnnoyI *ptr, const char *filename) {
    return ptr->on_disk_build(filename);
  }

}
