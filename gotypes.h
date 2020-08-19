#include <stdint.h>

// int
// TODO: choose between int32 and int64 depending on system (32-bit or a 64-bit machine)
typedef int64_t GoInt;

// float  <=> float32
// double <=> float64
typedef float GoFloat32;

// slice
typedef struct { void *ptr; GoInt len; GoInt cap; } GoSlice;
typedef struct { GoFloat32 *ptr; GoInt len; GoInt cap; } GoFloat32Slice;
