#include "annoylib.h"
#include "kissrandom.h"
#include "goannoy.h"

// Annoy Class
template <class D>
using AnnoyC = Annoy::AnnoyIndex<int32_t, float, D, Annoy::Kiss64Random, Annoy::AnnoyIndexSingleThreadedBuildPolicy>;

// Annoy Interface
using AnnoyI = Annoy::AnnoyIndexInterface<int32_t, float>;

typedef struct AnnoyIndex {
  AnnoyI* index;
} AnnoyIndex;

extern "C" {
  AnnoyIndex* CreateAnnoyIndexAngular(int f) {
    return new AnnoyIndex{new AnnoyC<Annoy::Angular>(f)};
  }

  AnnoyIndex* CreateAnnoyIndexEuclidean(int f) {
    return new AnnoyIndex{new AnnoyC<Annoy::Euclidean>(f)};
  }

  AnnoyIndex* CreateAnnoyIndexManhattan(int f) {
    return new AnnoyIndex{new AnnoyC<Annoy::Manhattan>(f)};
  }

  AnnoyIndex* CreateAnnoyIndexDotProduct(int f) {
    return new AnnoyIndex{new AnnoyC<Annoy::DotProduct>(f)};
  }

  void DeleteAnnoyIndex(AnnoyIndex *a) {
    a->index->unload();
    delete a->index;
  }

  bool AddItem(AnnoyIndex *a, intgo_t item, const float *w, char **error) {
    return a->index->add_item(item, w, error);
  }

  bool Build(AnnoyIndex *a, int q, char **error) {
    return a->index->build(q, 1, error);
  }

  bool Unbuild(AnnoyIndex *a, char **error) {
    return a->index->unbuild(error);
  }

  bool Save(AnnoyIndex *a, const char *filename, bool prefault, char **error) {
    return a->index->save(filename, prefault, error);
  }

  void Unload(AnnoyIndex *a) {
    a->index->unload();
  }

  bool Load(AnnoyIndex *a, const char *filename, bool prefault, char **error) {
    return a->index->load(filename, prefault, error);
  }

  float GetDistance(AnnoyIndex *a, int i, int j) {
    return a->index->get_distance(i, j);
  }

  int resultsToArrays(Annoy::vector<int32_t> *rv, Annoy::vector<float> *dv, int32_t *ra, float *da) {
    int size = rv->size();
    for (int i = 0; i < size; ++i) {
      ra[i] = (*rv)[i];
      da[i] = (*dv)[i];
    }
    return size;
  }

  int resultToArray(Annoy::vector<int32_t> *rv, int32_t *ra) {
    int size = rv->size();
    for (int i = 0; i < size; ++i)
      ra[i] = (*rv)[i];
    return size;
  }

  int GetNNsByItem(AnnoyIndex *a, int item, int n, int search_k, int32_t *result) {
    Annoy::vector<int32_t> *result_vec = new Annoy::vector<int32_t>();
    a->index->get_nns_by_item(item, n, search_k, result_vec, NULL);
    int size = resultToArray(result_vec, result);
    delete result_vec;
    return size;
  }

  int GetNNsByItemWithDistances(AnnoyIndex *a, int item, int n, int search_k, int32_t *result, float *distances) {
    Annoy::vector<int32_t>  *result_vec    = new Annoy::vector<int32_t>();
    Annoy::vector<float>    *distances_vec = new Annoy::vector<float>();
    a->index->get_nns_by_item(item, n, search_k, result_vec, distances_vec);
    int size = resultsToArrays(result_vec, distances_vec, result, distances);
    delete result_vec;
    delete distances_vec;
    return size;
  }

  int GetNNsByVector(AnnoyIndex *a, const float *w, int n, int search_k, int32_t *result) {
    Annoy::vector<int32_t> *result_vec = new Annoy::vector<int32_t>();
    a->index->get_nns_by_vector(w, n, search_k, result_vec, NULL);
    int size = resultToArray(result_vec, result);
    delete result_vec;
    return size;
  }

  int GetNNsByVectorWithDistances(AnnoyIndex *a, const float *w, int n, int search_k, int32_t *result, float *distances) {
    Annoy::vector<int32_t>  *result_vec    = new Annoy::vector<int32_t>();
    Annoy::vector<float>    *distances_vec = new Annoy::vector<float>();
    a->index->get_nns_by_vector(w, n, search_k, result_vec, distances_vec);
    int size = resultsToArrays(result_vec, distances_vec, result, distances);
    delete result_vec;
    delete distances_vec;
    return size;
  }

  int GetNItems(AnnoyIndex *a) {
    return (int)a->index->get_n_items();
  }

  void Verbose(AnnoyIndex *a, bool v) {
    a->index->verbose(v);
  }

  void GetItem(AnnoyIndex *a, int item, float *v) {
    a->index->get_item(item, v);
  }

  bool OnDiskBuild(AnnoyIndex *a, const char *filename, char **error) {
    return a->index->on_disk_build(filename, error);
  }

}
