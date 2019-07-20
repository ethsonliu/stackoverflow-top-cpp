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

**实验二：在实验一的基础上，加上数据库插入操作**

```c++
sprintf(sSQL, "INSERT INTO TTC VALUES (NULL, '%s', '%s', '%s', '%s', '%s', '%s', '%s')", sRT, sBR, sVR, sST, sVI, sDT, sTM);
sqlite3_exec(db, sSQL, NULL, NULL, &sErrMsg);
```

输出结果：

```
Imported 864913 records in 9933.61 seconds
```

很慢，因为每个插入都是在自己的事务里，频率约为 85 条每秒。

**实验三：在实验二的基础上，加入事务（Transaction）**

```c++
sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, &sErrMsg);

pFile = fopen (INPUTDATA,"r");
while (!feof(pFile)) {

    ...

}
fclose (pFile);

sqlite3_exec(db, "END TRANSACTION", NULL, NULL, &sErrMsg);
```

输出如下：

```
Imported 864913 records in 38.03 seconds
```

加入事务之后速度提高不少，频率约为 23,000 条每秒。

**实验四：在实验三的基础上，加入预处理**

仔细观察会发现，插入语句的格式一样的，那么完全可以用`sqlite3_prepare_v2`来预处理优化，

```c++
/* Open input file and import into the database */
cStartClock = clock();

sprintf(sSQL, "INSERT INTO TTC VALUES (NULL, @RT, @BR, @VR, @ST, @VI, @DT, @TM)");
sqlite3_prepare_v2(db,  sSQL, BUFFER_SIZE, &stmt, &tail);

sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, &sErrMsg);

pFile = fopen (INPUTDATA,"r");
while (!feof(pFile)) {

    fgets (sInputBuf, BUFFER_SIZE, pFile);

    sRT = strtok (sInputBuf, "\t");   /* Get Route */
    sBR = strtok (NULL, "\t");        /* Get Branch */
    sVR = strtok (NULL, "\t");        /* Get Version */
    sST = strtok (NULL, "\t");        /* Get Stop Number */
    sVI = strtok (NULL, "\t");        /* Get Vehicle */
    sDT = strtok (NULL, "\t");        /* Get Date */
    sTM = strtok (NULL, "\t");        /* Get Time */

    sqlite3_bind_text(stmt, 1, sRT, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 2, sBR, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 3, sVR, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 4, sST, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 5, sVI, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 6, sDT, -1, SQLITE_TRANSIENT);
    sqlite3_bind_text(stmt, 7, sTM, -1, SQLITE_TRANSIENT);

    sqlite3_step(stmt);

    sqlite3_clear_bindings(stmt);
    sqlite3_reset(stmt);

    n++;
}
fclose (pFile);

sqlite3_exec(db, "END TRANSACTION", NULL, NULL, &sErrMsg);

printf("Imported %d records in %4.2f seconds\n", n, (clock() - cStartClock) / (double)CLOCKS_PER_SEC);

sqlite3_finalize(stmt);
sqlite3_close(db);

return 0;
```

输出如下：

```
Imported 864913 records in 16.27 seconds
```

速度更快了，频率约为 53,000 条每秒。

**实验五：在实验四的基础上，加入 PRAGMA synchronous = OFF**

默认情况下，SQLite 为了保证插入操作中的数据可以被写入磁盘，在调用系统 API 的`write`之后会暂停等待其完成，我们可以使用`PRAGMA synchronous = OFF`来关闭这个暂停等待。但注意，这个做法在系统崩溃或写入数据时意外断电的情况下数据库文件可能会损坏。

```c++
/* Open the database and create the schema */
sqlite3_open(DATABASE, &db);
sqlite3_exec(db, TABLE, NULL, NULL, &sErrMsg);
sqlite3_exec(db, "PRAGMA synchronous = OFF", NULL, NULL, &sErrMsg);
```

输出如下：

```
Imported 864913 records in 12.41 seconds
```

时间变得又少了点，频率约为 64,000 条每秒。

**实验六：在实验四的基础上，加入 PRAGMA journal_mode = MEMORY**

回滚日志文件（Rollback Journals），用于实现数据库的原子提交和回滚。 此文件和数据库文件总是在同一个目录，并且有相同的文件名，但是在文件名中添加了一个`-journal`字符串。此文件一般在`transaction`开始时创建，`transaction`结束时删除。

如果系统 crash，Rollback Journals 文件将被保留，下次打开数据库文件时，系统会检查有没有 Rollback journals 文件存在，如果有就用它来恢复数据库。

SQLite 默认会把回滚日志文件保存在磁盘上，现在改为保存在内存中，避免了磁盘 I/O。但注意，如果系统 crash，数据库文件可能也会 crash。

```c++
/* Open the database and create the schema */
sqlite3_open(DATABASE, &db);
sqlite3_exec(db, TABLE, NULL, NULL, &sErrMsg);
sqlite3_exec(db, "PRAGMA journal_mode = MEMORY", NULL, NULL, &sErrMsg);
```

输出如下：

```
Imported 864913 records in 13.50 seconds
```

比实验五稍微慢了点，但比实验四快了点，频率约为 64,000 条每秒。

**实验七：在实验四的基础上，同时加上 PRAGMA synchronous = OFF 和 PRAGMA journal_mode = MEMORY**

这次我们把实验五和实验六合并在一起再看看，

```c++
/* Open the database and create the schema */
sqlite3_open(DATABASE, &db);
sqlite3_exec(db, TABLE, NULL, NULL, &sErrMsg);
sqlite3_exec(db, "PRAGMA synchronous = OFF", NULL, NULL, &sErrMsg);
sqlite3_exec(db, "PRAGMA journal_mode = MEMORY", NULL, NULL, &sErrMsg);
```

输出如下：

```
Imported 864913 records in 12.00 seconds
```

变得更好了，频率约为 72,000 条每秒。

**实验八：在实验七的基础上，做些代码重构**

让`strtok`直接赋值给`sqlite3_bind_text`，

```c++
pFile = fopen (INPUTDATA,"r");
while (!feof(pFile)) {

    fgets (sInputBuf, BUFFER_SIZE, pFile);

    sqlite3_bind_text(stmt, 1, strtok (sInputBuf, "\t"), -1, SQLITE_TRANSIENT); /* Get Route */
    sqlite3_bind_text(stmt, 2, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Branch */
    sqlite3_bind_text(stmt, 3, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Version */
    sqlite3_bind_text(stmt, 4, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Stop Number */
    sqlite3_bind_text(stmt, 5, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Vehicle */
    sqlite3_bind_text(stmt, 6, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Date */
    sqlite3_bind_text(stmt, 7, strtok (NULL, "\t"), -1, SQLITE_TRANSIENT);    /* Get Time */

    sqlite3_step(stmt);        /* Execute the SQL Statement */
    sqlite3_clear_bindings(stmt);    /* Clear bindings */
    sqlite3_reset(stmt);        /* Reset VDBE */

    n++;
}
fclose (pFile);
```

输出如下：

```
Imported 864913 records in 8.94 seconds
```

一个小小的变动，频率就达到了 96,700 条每秒。

**实验九：在实验七的基础上，使用 In-Memory Databases**

数据库定义在内存中（除非有特殊用途，否则还是乖乖地定义在磁盘上），

```c++
#define DATABASE ":memory:"
```

输出如下：

```
Imported 864913 records in 10.94 seconds
```

频率约为 79,000 条每秒。

**总结**

实验结果已说明一切了，实际应用各取所需即可。

值得一提的是，如果加入索引（Index）的顺序不同也会导致速度有所差异。在实验八的基础上，我们加入索引，

```c++
sqlite3_exec(db, "CREATE  INDEX 'TTC_Stop_Index' ON 'TTC' ('Stop')", NULL, NULL, &sErrMsg);
sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, &sErrMsg);
...
```

先创建索引，再插入数据：输出为`Imported 864913 records in 18.13 seconds`。

```c++
..
sqlite3_exec(db, "END TRANSACTION", NULL, NULL, &sErrMsg);
sqlite3_exec(db, "CREATE  INDEX 'TTC_Stop_Index' ON 'TTC' ('Stop')", NULL, NULL, &sErrMsg);
```

先插入数据，再创建索引：输出为`Imported 864913 records in 13.66 seconds`。

## 回答

