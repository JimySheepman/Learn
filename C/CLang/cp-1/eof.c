#include <stdio.h>

int main()
{
    printf("Please enter a character:\n");
    printf("The value of EOF is %d.\n", EOF);
    printf("The expression \"getchar() != EOF\" is %d.\n", getchar() != EOF);
    return 0;
}
