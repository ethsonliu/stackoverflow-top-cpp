<https://stackoverflow.com/questions/1711631/improve-insert-per-second-performance-of-sqlite>

## 问题

SQLite 的优化比较棘手，就批量插入而言，其速度可以从每秒 85 条优化到每秒 96,000 条。下面我们来具体看下实验过程和结果，

**背景：**

1. 文件数据：[多伦多市全部交通时间表](http://www.toronto.ca/open/datasets/ttc-routes)，大小约 28MB，以 TAB 分隔的文本文件（约 865,000 条记录）
2. 机器环境： Windows XP 3.60 GHz P4
3. 编译环境：[Visual C++](http://en.wikipedia.org/wiki/Visual_C%2B%2B#32-bit_versions) 2005 Release，使用完全优化（/ Ox）和优先快速代码（/ Ot）
4. 数据库：SQLite 3.6.7

**实验一：建表 + 读取解析数据**

一个简单的 C 程序，逐行读取文本文件，将字符串拆分为值，但先不把数据插入到 SQLite 数据库中。代码如下：

```c++
/*************************************************************
    Baseline code to experiment with SQLite performance.

    Input data is a 28 MB TAB-delimited text file of the
    complete Toronto Transit System schedule/route info
    from http://www.toronto.ca/open/datasets/ttc-routes/

**************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
#include "sqlite3.h"

#define INPUTDATA "C:\\TTC_schedule_scheduleitem_10-27-2009.txt"
#define DATABASE "c:\\TTC_schedule_scheduleitem_10-27-2009.sqlite"
#define TABLE "CREATE TABLE IF NOT EXISTS TTC (id INTEGER PRIMARY KEY, Route_ID TEXT, Branch_Code TEXT, Version INTEGER, Stop INTEGER, Vehicle_Index INTEGER, Day Integer, Time TEXT)"
#define BUFFER_SIZE 256

int main(int argc, char **argv) {

    sqlite3 * db;
    sqlite3_stmt * stmt;
    char * sErrMsg = 0;
    char * tail = 0;
    int nRetCode;
    int n = 0;

    clock_t cStartClock;

    FILE * pFile;
    char sInputBuf [BUFFER_SIZE] = "\0";

    char * sRT = 0;  /* Route */
    char * sBR = 0;  /* Branch */
    char * sVR = 0;  /* Version */
    char * sST = 0;  /* Stop Number */
    char * sVI = 0;  /* Vehicle */
    char * sDT = 0;  /* Date */
    char * sTM = 0;  /* Time */

    char sSQL [BUFFER_SIZE] = "\0";

    /*********************************************/
    /* Open the Database and create the Schema */
    sqlite3_open(DATABASE, &db);
    sqlite3_exec(db, TABLE, NULL, NULL, &sErrMsg);

    /*********************************************/
    /* Open input file and import into Database*/
    cStartClock = clock();

    pFile = fopen (INPUTDATA,"r");
    while (!feof(pFile)) {

        fgets (sInputBuf, BUFFER_SIZE, pFile);

        sRT = strtok (sInputBuf, "\t");     /* Get Route */
        sBR = strtok (NULL, "\t");            /* Get Branch */
        sVR = strtok (NULL, "\t");            /* Get Version */
        sST = strtok (NULL, "\t");            /* Get Stop Number */
        sVI = strtok (NULL, "\t");            /* Get Vehicle */
        sDT = strtok (NULL, "\t");            /* Get Date */
        sTM = strtok (NULL, "\t");            /* Get Time */

        /* ACTUAL INSERT WILL GO HERE */

        n++;
    }
    fclose (pFile);

    printf("Imported %d records in %4.2f seconds\n", n, (clock() - cStartClock) / (double)CLOCKS_PER_SEC);

    sqlite3_close(db);
    return 0;
}
```

输出如下：

```
Imported 864913 records in 0.94 seconds
```

可以看出，原生 C 程序的 I/O 和字符串操作还是很快的。

**实验二：**

## 回答

