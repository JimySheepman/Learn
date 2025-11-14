#include <stdio.h>

int main()
{
    int c, blank_recived;
    blank_recived = 0;
    printf("Input some characters, when you finishes, press Ctrl+D.\n");
    while ((c = getchar()) != EOF)
        if (c == ' ')
        {
            if (!blank_recived)
            {
                blank_recived++;
                putchar(c);
            }
        }
        else
        {
            blank_recived = 0;
            putchar(c);
        }

    return 0;
}