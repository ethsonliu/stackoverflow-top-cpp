<https://stackoverflow.com/questions/1538420/difference-between-malloc-and-calloc>

## 问题

下面的两句代码有什么区别？

```c
ptr = (char **) malloc (MAXELEMS * sizeof(char *));
```

vs

```c
ptr = (char **) calloc (MAXELEMS, sizeof(char*));
```

主要是 `malloc` 和 `calloc` 的区别。

## 回答

`calloc` 会申请内存，并全初始化为 0；而 `malloc` 只申请内存，并不作初始化。

所以 `calloc` 的执行会比 `malloc` 稍微费时，因为它多了初始化的步骤。
