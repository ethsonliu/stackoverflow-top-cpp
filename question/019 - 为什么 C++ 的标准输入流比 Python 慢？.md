<https://stackoverflow.com/questions/9371238/why-is-reading-lines-from-stdin-much-slower-in-c-than-python>

## 问题

我想比较一下 C++ 和 Python 的标准输入，但实验的结果让人大吃一惊，C++ 慢了许多。下面是我的实验代码：

**C++ 代码**

```c++
#include <iostream>
#include <time.h>

using namespace std;

int main() {
    string input_line;
    long line_count = 0;
    time_t start = time(NULL);
    int sec;
    int lps;

    while (cin) {
        getline(cin, input_line);
        if (!cin.eof())
            line_count++;
    };

    sec = (int) time(NULL) - start;
    cerr << "Read " << line_count << " lines in " << sec << " seconds.";
    if (sec > 0) {
        lps = line_count / sec;
        cerr << " LPS: " << lps << endl;
    } else
        cerr << endl;
    return 0;
}

// Compiled with:
// g++ -O3 -o readline_test_cpp foo.cpp
```

**Python 代码**

```python
#!/usr/bin/env python
import time
import sys

count = 0
start = time.time()

for line in  sys.stdin:
    count += 1

delta_sec = int(time.time() - start_time)
if delta_sec >= 0:
    lines_per_sec = int(round(count/delta_sec))
    print("Read {0} lines in {1} seconds. LPS: {2}".format(count, delta_sec,
       lines_per_sec))
```

**下面是实验结果**

```shell
$ cat test_lines | ./readline_test_cpp
Read 5570000 lines in 9 seconds. LPS: 618889

$cat test_lines | ./readline_test.py
Read 5570000 lines in 1 seconds. LPS: 5570000
```

我在 Mac OS X v10.6.8 和 Linux 2.6.32 (Red Hat Linux 6.2) 都测试过，

```shell
$ for i in {1..5}; do echo "Test run $i at `date`"; echo -n "CPP:"; cat test_lines | ./readline_test_cpp ; echo -n "Python:"; cat test_lines | ./readline_test.py ; done
Test run 1 at Mon Feb 20 21:29:28 EST 2012
CPP:   Read 5570001 lines in 9 seconds. LPS: 618889
Python:Read 5570000 lines in 1 seconds. LPS: 5570000
Test run 2 at Mon Feb 20 21:29:39 EST 2012
CPP:   Read 5570001 lines in 9 seconds. LPS: 618889
Python:Read 5570000 lines in 1 seconds. LPS: 5570000
Test run 3 at Mon Feb 20 21:29:50 EST 2012
CPP:   Read 5570001 lines in 9 seconds. LPS: 618889
Python:Read 5570000 lines in 1 seconds. LPS: 5570000
Test run 4 at Mon Feb 20 21:30:01 EST 2012
CPP:   Read 5570001 lines in 9 seconds. LPS: 618889
Python:Read 5570000 lines in 1 seconds. LPS: 5570000
Test run 5 at Mon Feb 20 21:30:11 EST 2012
CPP:   Read 5570001 lines in 10 seconds. LPS: 557000
Python:Read 5570000 lines in  1 seconds. LPS: 5570000
```

## 回答

默认情况下，cin 与 stdin 总是保持同步的，也就是说这两种方法可以混用，而不必担心文件指针混乱，同时 cout 和 stdout 也一样，两者混用不会输出顺序错乱。正因为这个兼容性的特性，导致 cin 有许多额外的开销，如何禁用这个特性呢？

```c++
std::ios_base::sync_with_stdio(false);
```

这样就可以取消 cin 于 stdin 的同步了。

通常，输入流都是从缓冲区读取内容，而 `stdio` 和 `iostreams` 都有自己的缓冲区，如果一起使用就会出现未知的问题。比如：

```c++
int myvalue1;
cin >> myvalue1;
int myvalue2;
scanf("%d",&myvalue2);
```

如果在控制台同时输入`1 2`，按我们的预想，cin 拿到的值是 1，scanf 拿到的是 2，但事实可能并非如此：scanf 可能拿不到 2，因为 2 这个值在 cin 的缓冲区那里，scanf 缓冲区什么也没有。（如果调用 `std::ios_base::sync_with_stdio(false)`，程序就需要考虑到这点，以免出现未知错误）

为了避免这种情况，C++ 默认使 cin 与 stdio 同步，这样就不会出现问题。
