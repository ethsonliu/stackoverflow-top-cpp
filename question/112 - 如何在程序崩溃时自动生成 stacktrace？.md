<https://stackoverflow.com/questions/77005/how-to-automatically-generate-a-stacktrace-when-my-program-crashes>

## 问题

有什么好的办法可以在 C/C++ 程序段错误退出时输出堆栈信息，来方便查找错误么？

## 回答

在 Linux 平台下可以使用 `<execinfo.h>` 里的 `backtrace_*` 函数，详见 [Backtraces](http://www.gnu.org/software/libc/manual/html_node/Backtraces.html)，例子如下，

```c++
#include <stdio.h>
#include <execinfo.h>
#include <signal.h>
#include <stdlib.h>
#include <unistd.h>


void handler(int sig) {
  void *array[10];
  size_t size;

  // get void*'s for all entries on the stack
  size = backtrace(array, 10);

  // print out all the frames to stderr
  fprintf(stderr, "Error: signal %d:\n", sig);
  backtrace_symbols_fd(array, size, STDERR_FILENO);
  exit(1);
}

void baz() {
 int *foo = (int*)-1; // make a bad pointer
  printf("%d\n", *foo);       // causes segfault
}

void bar() { baz(); }
void foo() { bar(); }


int main(int argc, char **argv) {
  signal(SIGSEGV, handler);   // install our handler
  foo(); // this will call foo, bar, and baz.  baz segfaults.
}
```

输入 `gcc -g -rdynamic ./test.c -o test` 进行编译，接着运行就会输出如下，

```
$ ./test
Error: signal 11:
./test(handler+0x19)[0x400911]
/lib64/tls/libc.so.6[0x3a9b92e380]
./test(baz+0x14)[0x400962]
./test(bar+0xe)[0x400983]
./test(foo+0xe)[0x400993]
./test(main+0x28)[0x4009bd]
/lib64/tls/libc.so.6(__libc_start_main+0xdb)[0x3a9b91c4bb]
./test[0x40086a]
```

可以看到是在 `baz` 中出现错误。


