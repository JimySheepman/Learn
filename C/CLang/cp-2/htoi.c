#include<stdio.h>

#define MAXLINE 1000

int get_line(char line[], int maxline);
unsigned long htio(char s[]);

int main(){
    int len;
    char line[MAXLINE];

    while ((len = get_line(line,MAXLINE))>0)
    {
        printf("%lu\n",htio(line));
    }

    return 0;    
}

int get_line(char s[],int lim){
    int c,i,l;

    for(i = 0, l=0;(c = getchar())!= EOF && c!= '\n';i++)
        if (i<lim-1)
            s[l++] = c;
    
    s[l] = '\0';

    return 1;
}

unsigned long htoi(char s[])
{
    unsigned long n;
    int i, c;

    n = 0;
    for (i = 0; s[i] != '\0'; ++i) {
        c = s[i];
        if (i == 0 && c == '0' && (s[1] == 'x' || s[1] == 'X')) {
            i = 1;
            continue;
    }
    n *= 16;
    if (c >= '0' && c <= '9')
        n += c - '0';
    else if (c >= 'a' && c <= 'f')
        n += c - 'a';
    else if (c >= 'A' && c <= 'F')
        n += c - 'A';
    }

    return n;
}