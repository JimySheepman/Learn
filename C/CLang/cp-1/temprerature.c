#include <stdio.h>

float ftoc(float fahr);

int main(void)
{
    float fahr, celsius;
    float lower, upper, step;
    lower = 0;
    upper = 300;
    step = 20;
    fahr = lower;

    while (fahr <= upper)
    {
        celsius = ftoc(fahr);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }

    return 0;
}

float ftoc(float fahr)
{
    float celsius;
    celsius = (5.0 / 9.0) * (fahr - 32.0);
    return celsius;
}