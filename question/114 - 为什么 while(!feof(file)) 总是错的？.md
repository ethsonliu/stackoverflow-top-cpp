<https://stackoverflow.com/questions/5431941/why-is-while-feof-file-always-wrong>

## 问题

我发现很多人在读取文件时都是这么用的，

```c
#include <stdio.h>
#include <stdlib.h>

int
main(int argc, char **argv)
{
    char *path = "stdin";
    FILE *fp = argc > 1 ? fopen(path=argv[1], "r") : stdin;

    if( fp == NULL ) {
        perror(path);
        return EXIT_FAILURE;
    }

    while( !feof(fp) ) {  /* THIS IS WRONG */
        /* Read and process data from file… */
    }
    if( fclose(fp) != 0 ) {
        perror(path);
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}
```

上面的循环有什么问题么？

## 回答
