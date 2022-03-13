#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main()
{

    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    int n, i;
    scanf("%d", &n);
    int b[100], a;
    for (i = 0; i < 100; i++)
    {
        b[i] = 0;
    }
    for (i = 0; i < n; i++)
    {
        scanf("%d", &a);
        b[a]++;
    }
    for (i = 0; i < 100; i++)
    {
        printf("%d ", b[i]);
    }

    return 0;
}
