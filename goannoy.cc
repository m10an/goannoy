#include "annoylib.h"
#include "kissrandom.h"
#include "gotypes.h"

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

  bool add_item(AnnoyI *ptr, intgo_t item, const float *w, char **error) {
    return ptr->add_item(item, w, error);
  }

  bool build(AnnoyI *ptr, int q, char **error) {
    return ptr->build(q, 1, error);
  }

  bool unbuild(AnnoyI *ptr, char **error) {
    return ptr->unbuild(error);
  }

  bool save(AnnoyI *ptr, const char *filename, bool prefault, char **error) {
    return ptr->save(filename, prefault, error);
  }

  void unload(AnnoyI *ptr) {
    ptr->unload();
  }

  bool load(AnnoyI *ptr, const char *filename, bool prefault, char **error) {
    return ptr->load(filename, prefault, error);
  }

  float get_distance(AnnoyI *ptr, int i, int j) {
    return ptr->get_distance(i, j);
  }

  int _results_to_arrays(vector<int32_t> *rv, vector<float> *dv, int32_t *ra, float *da) {
    int size = rv->size();
    for (int i = 0; i < size; ++i) {
      ra[i] = (*rv)[i];
      da[i] = (*dv)[i];
    }
    return size;
  }

  int _result_to_array(vector<int32_t> *rv, int32_t *ra) {
    int size = rv->size();
    for (int i = 0; i < size; ++i)
      ra[i] = (*rv)[i];
    return size;
  }

  int get_nns_by_item(AnnoyI *ptr, int item, int n, int search_k, int32_t *result) {
    vector<int32_t> *result_vec = new vector<int32_t>();
    ptr->get_nns_by_item(item, n, search_k, result_vec, NULL);
    int size = _result_to_array(result_vec, result);
    delete result_vec;
    return size;
  }

  int get_nns_by_item_with_dists(AnnoyI *ptr, int item, int n, int search_k, int32_t *result, float *distances) {
    vector<int32_t>  *result_vec    = new vector<int32_t>();
    vector<float>    *distances_vec = new vector<float>();
    ptr->get_nns_by_item(item, n, search_k, result_vec, distances_vec);
    int size = _results_to_arrays(result_vec, distances_vec, result, distances);
    delete result_vec;
    delete distances_vec;
    return size;
  }

  int get_nns_by_vector(AnnoyI *ptr, const float *w, int n, int search_k, int32_t *result) {
    vector<int32_t> *result_vec = new vector<int32_t>();
    ptr->get_nns_by_vector(w, n, search_k, result_vec, NULL);
    int size = _result_to_array(result_vec, result);
    delete result_vec;
    return size;
  }

  int get_nns_by_vector_with_dists(AnnoyI *ptr, const float *w, int n, int search_k, int32_t *result, float *distances) {
    vector<int32_t>  *result_vec    = new vector<int32_t>();
    vector<float>    *distances_vec = new vector<float>();
    ptr->get_nns_by_vector(w, n, search_k, result_vec, distances_vec);
    int size = _results_to_arrays(result_vec, distances_vec, result, distances);
    delete result_vec;
    delete distances_vec;
    return size;
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
