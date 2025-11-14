#include <sys/wait.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#define ANSI_COLOR_RED     "\x1b[31m"

int ksh_cd(char **args);
int ksh_help(char **args);
int ksh_exit(char **args);

char *builtin_str[] = {
    "cd",
    "help",
    "exit"};

int (*builtin_func[])(char **) = {
    &ksh_cd,
    &ksh_help,
    &ksh_exit};

int ksh_num_builtins()
{
    return sizeof(builtin_str) / sizeof(char *);
}

int ksh_cd(char **args)
{
    if (args[1] == NULL)
    {
        fprintf(stderr, "ksh: expected argument to \"cd\"\n");
    }
    else
    {
        if (chdir(args[1]) != 0)
        {
            perror("ksh");
        }
    }
    return 1;
}

int ksh_help(char **args)
{
    int i;
    printf("Merlins KSH\n");
    printf("Type program names and arguments, and hit enter.\n");
    printf("The following are built in:\n");

    for (i = 0; i < ksh_num_builtins(); i++)
    {
        printf("  %s\n", builtin_str[i]);
    }

    printf("Use the man command for information on other programs.\n");
    return 1;
}

int ksh_exit(char **args)
{
    return 0;
}

int ksh_launch(char **args)
{
    pid_t pid;
    int status;

    pid = fork();
    if (pid == 0)
    {
        if (execvp(args[0], args) == -1)
        {
            perror("ksh");
        }
        exit(EXIT_FAILURE);
    }
    else if (pid < 0)
    {
        perror("ksh");
    }
    else
    {
        do
        {
            waitpid(pid, &status, WUNTRACED);
        } while (!WIFEXITED(status) && !WIFSIGNALED(status));
    }

    return 1;
}

int ksh_execute(char **args)
{
    int i;

    if (args[0] == NULL)
    {
        return 1;
    }

    for (i = 0; i < ksh_num_builtins(); i++)
    {
        if (strcmp(args[0], builtin_str[i]) == 0)
        {
            return (*builtin_func[i])(args);
        }
    }

    return ksh_launch(args);
}

char *ksh_read_line(void)
{
#ifdef KSH_USE_STD_GETLINE
    char *line = NULL;
    ssize_t bufsize = 0; // have getline allocate a buffer for us
    if (getline(&line, &bufsize, stdin) == -1)
    {
        if (feof(stdin))
        {
            exit(EXIT_SUCCESS); // We received an EOF
        }
        else
        {
            perror("ksh: getline\n");
            exit(EXIT_FAILURE);
        }
    }
    return line;
#else
#define KSH_RL_BUFSIZE 1024
    int bufsize = KSH_RL_BUFSIZE;
    int position = 0;
    char *buffer = malloc(sizeof(char) * bufsize);
    int c;

    if (!buffer)
    {
        fprintf(stderr, "ksh: allocation error\n");
        exit(EXIT_FAILURE);
    }

    while (1)
    {
        c = getchar();

        if (c == EOF)
        {
            exit(EXIT_SUCCESS);
        }
        else if (c == '\n')
        {
            buffer[position] = '\0';
            return buffer;
        }
        else
        {
            buffer[position] = c;
        }
        position++;

        if (position >= bufsize)
        {
            bufsize += KSH_RL_BUFSIZE;
            buffer = realloc(buffer, bufsize);
            if (!buffer)
            {
                fprintf(stderr, "Ksh: allocation error\n");
                exit(EXIT_FAILURE);
            }
        }
    }
#endif
}

#define KSH_TOK_BUFSIZE 64
#define KSH_TOK_DELIM " \t\r\n\a"

char **ksh_split_line(char *line)
{
    int bufsize = KSH_TOK_BUFSIZE, position = 0;
    char **tokens = malloc(bufsize * sizeof(char *));
    char *token, **tokens_backup;

    if (!tokens)
    {
        fprintf(stderr, "Ksh: allocation error\n");
        exit(EXIT_FAILURE);
    }

    token = strtok(line, KSH_TOK_DELIM);
    while (token != NULL)
    {
        tokens[position] = token;
        position++;

        if (position >= bufsize)
        {
            bufsize += KSH_TOK_BUFSIZE;
            tokens_backup = tokens;
            tokens = realloc(tokens, bufsize * sizeof(char *));
            if (!tokens)
            {
                free(tokens_backup);
                fprintf(stderr, "Ksh: allocation error\n");
                exit(EXIT_FAILURE);
            }
        }

        token = strtok(NULL, KSH_TOK_DELIM);
    }
    tokens[position] = NULL;
    return tokens;
}

void ksh_loop(void)
{
    char *line;
    char **args;
    int status;

    do
    {
        printf(ANSI_COLOR_RED     "merlins-> ");
        line = ksh_read_line();
        args = ksh_split_line(line);
        status = ksh_execute(args);
        free(line);
        free(args);
    } while (status);
}
int main(int argc, char **argv)
{
    ksh_loop();
    return EXIT_SUCCESS;
}