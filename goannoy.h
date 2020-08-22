#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>

// int
// TODO: choose between int32 and int64 depending on system (32-bit or a 64-bit machine)
typedef int64_t intgo_t;
typedef int64_t GoInt;
typedef int32_t GoInt32;

#ifdef __cplusplus
extern "C" {
#endif

typedef struct AnnoyIndex AnnoyIndex;

AnnoyIndex * CreateAnnoyIndexAngular(int);
AnnoyIndex * CreateAnnoyIndexEuclidean(int);
AnnoyIndex * CreateAnnoyIndexManhattan(int);
AnnoyIndex * CreateAnnoyIndexDotProduct(int);
void DeleteAnnoyIndex(AnnoyIndex *);
bool AddItem(AnnoyIndex *, intgo_t, const float *, char **);
bool Build(AnnoyIndex *, int, char **);
bool Unbuild(AnnoyIndex *, char **);
bool Save(AnnoyIndex *, const char *, bool, char **);
void Unload(AnnoyIndex *);
bool Load(AnnoyIndex *, const char *, bool, char **);
float GetDistance(AnnoyIndex *, int, int);
int GetNNsByItem(AnnoyIndex *, int, int, int, int32_t *);
int GetNNsByItemWithDistances(AnnoyIndex *, int, int, int, int32_t *, float *);
int GetNNsByVector(AnnoyIndex *, const float *, int, int, int32_t *);
int GetNNsByVectorWithDistances(AnnoyIndex *, const float *, int, int, int32_t *, float *);
int GetNItems(AnnoyIndex *);
void Verbose(AnnoyIndex *, bool);
void GetItem(AnnoyIndex *, int, float *);
bool OnDiskBuild(AnnoyIndex *, const char *, char **);

#ifdef __cplusplus
} // end extern "C"
#endif
