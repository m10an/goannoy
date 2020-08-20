#include <stdint.h>

// int
// TODO: choose between int32 and int64 depending on system (32-bit or a 64-bit machine)
typedef int64_t intgo_t;
typedef int64_t GoInt;
typedef int32_t GoInt32;

// float32 slice
typedef const struct _GoFloatSlice { float *ptr; intgo_t len; intgo_t cap; } float_slice_t;
typedef float_slice_t GoFloatSlice;
// int slice
typedef const struct _GoIntSlice { intgo_t *ptr; intgo_t len; intgo_t cap; } intgo_slice_t;
typedef intgo_slice_t GoIntSlice;
