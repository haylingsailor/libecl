#include "libecl.h"

// Wrapper implementation - for converting arguments into and out of go world.

extern int elastic_activateFeature(const char* feature) {
    GoString gFeature = {(char *)feature, strlen(feature)};

    return (int) ActivateFeature(gFeature);
}

