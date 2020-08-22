#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include "gotypes.h"

void * create_annidx_angular(int);
void * create_annidx_euclidean(int);
void * create_annidx_manhattan(int);
void * create_annidx_dot_product(int);
void free_annidx(void *);
bool add_item(void *, intgo_t, const float *, char **);
bool build(void *, int, char **);
bool unbuild(void *, char **);
bool save(void *, const char *, bool, char **);
void unload(void *);
bool load(void *, const char *, bool, char **);
float get_distance(void *, int, int);
int get_nns_by_item(void *, int, int, int, int32_t *);
int get_nns_by_item_with_dists(void *, int, int, int, int32_t *, float *);
int get_nns_by_vector(void *, const float *, int, int, int32_t *);
int get_nns_by_vector_with_dists(void *, const float *, int, int, int32_t *, float *);
int get_n_items(void *);
void verbose(void *, bool);
void get_item(void *, int, float *);
bool on_disk_build(void *, const char *, char **);
