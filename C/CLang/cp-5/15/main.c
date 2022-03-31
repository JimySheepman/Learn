#include <stdio.h>
#include "string.h"
#include "lines.h"
#include "qsort.h"
#include "cmp.h"

char *lineptr[MAXLINES]; /* pointers to text lines */

int main(int argc, char **argv)
{
    int nlines; /* number of input lines read */
    int numeric = 0; /* 1 if numeric sort */
    int order = 1; /* 1 if sorting in increasing order */
    int ignorecase = 0; /* 1 if case insensitive */
    int c;

    while (--argc > 0 && *(*++argv) == '-') {
        while ((c = *++(*argv)) != '\0') {
            switch (c) {
                case 'n':
                    numeric = 1;
                    break;
                case 'r':
                    order = 0;
                    break;
                case 'f':
                    ignorecase = 1;
                    break;
                default:
                    break;
            }
        }
    }

    if ((nlines = readlines(lineptr, MAXLINES)) >= 0) {
        qsort((void**) lineptr, 0, nlines-1,
            (int (*)(void*,void*))(numeric ? numcmp : (ignorecase ? strcmpignore : strcmp)), order);
        writelines(lineptr, nlines);
        return 0;
    } else {
        printf("error: input too big to sort\n");
        return 1;
    }
}
