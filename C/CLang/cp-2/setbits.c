#include <stdio.h>

unsigned setbits(unsigned x, int p, int n, unsigned y);

int main(void)
{
    printf("%u\n", setbits(5732, 6, 3, 9823));
    return 0;
}

unsigned setbits(unsigned x, int p, int n, unsigned y)
{
    unsigned a = x & ~(~(~0U << n) << (p + 1 - n));

    unsigned b = (y & ~(~0U << n)) << (p + 1 - n);

    return a | b;
}