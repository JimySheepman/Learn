#include <stdio.h>

void _strncpy(char *s, char *t, int n);
int getchars(char *s, int max);

int main(void)
{
    char *t = "testring1234";
    char *s;

    _strncpy(s, t, 8);
    printf("strncpy: %s\n", s);

    return 0;
}

void _strncpy(char *s, char *t, int n)
{
    while (n-- > 0 && (*s++ = *t++))
        ;
    *s = '\0';
}