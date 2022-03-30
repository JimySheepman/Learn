#include <stdio.h>
#include <string.h>
#include "lines.h"

/* readlines: read input lines */
int readlines(char *lineptr[], int maxlines)
{
    int len, nlines;
    char line[MAXLEN];

    nlines = 0;
    while ((len = _getline(line, MAXLEN)) > 0)
        if (nlines >= maxlines)
            return -1;
        else {
            line[len-1] = '\0'; /* delete newline */
            strcpy(lineptr[nlines++], line);
        }
    return nlines;
}

/* writelines: write output lines */
void writelines(char *lineptr[], int nlines)
{
    int i;
    for (i = 0; i < nlines; i++)
        printf("%s\n", lineptr[i]);
}

int _getline(char s[], int max)
{
    int c, i, l;

    for (i = 0, l = 0; (c = getchar()) != EOF && c != '\n'; ++i)
        if (i < max - 1)
            s[l++] = c;
    if (c == '\n' && l < max - 1)
        s[l++] = c;
    s[l] = '\0';

    return l;
}