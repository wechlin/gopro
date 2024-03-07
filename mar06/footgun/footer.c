#include "header.h"

int cube(int x)
{
#ifdef SYMBOL
    return x*x*x;
#else
    return x+x+x;
#endif
}