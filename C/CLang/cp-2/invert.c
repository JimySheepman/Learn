#include <stdio.h>

unsigned invert(unsigned x, int p, int n);

int main(void)
{
    printf("%u\n", invert(5732, 6, 3));
    return 0;
}

unsigned invert(unsigned x, int p, int n)
{   
    unsigned mask = ~(~0U << n) << (p + 1 - n);
     
    return x ^ mask;
}