#define MAXLINES 50 /* max #lines to be sorted */
#define MAXLEN 100 /* max length of any input line */

int _getline(char *, int);
int readlines(char *lineptr[], int nlines);
void writelines(char *lineptr[], int nlines);