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
