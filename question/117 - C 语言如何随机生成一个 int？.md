<https://stackoverflow.com/questions/822323/how-to-generate-a-random-int-in-c>

## 问题

C 语言是否有一个函数可以随机生成一个整数？或者有其他的三方库可以实现的？

## 回答

```c
#include <time.h>
#include <stdlib.h>

srand(time(NULL));   // Initialization, should only be called once.
int r = rand();      // Returns a pseudo-random integer between 0 and RAND_MAX.
```

Linux 平台上建议使用 [random and srandom](https://linux.die.net/man/3/random)。

如果你需要更安全的随机数，建议使用 [libsodium](https://github.com/jedisct1/libsodium) 的接口 `randombytes`，

```c
#include "sodium.h"

int foo()
{
    char myString[32];
    uint32_t myInt;

    if (sodium_init() < 0) {
        /* panic! the library couldn't be initialized, it is not safe to use */
        return 1; 
    }


    /* myString will be an array of 32 random bytes, not null-terminated */        
    randombytes_buf(myString, 32);

    /* myInt will be a random number between 0 and 9 */
    myInt = randombytes_uniform(10);
}
```

Openssl 也提供了接口来实现，

```c
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <openssl/rand.h>

/* Random integer in [0, limit) */
unsigned int random_uint(unsigned int limit) {
    union {
        unsigned int i;
        unsigned char c[sizeof(unsigned int)];
    } u;

    do {
        if (!RAND_bytes(u.c, sizeof(u.c))) {
            fprintf(stderr, "Can't get random bytes!\n");
            exit(1);
        }
    } while (u.i < (-limit % limit)); /* u.i < (2**size % limit) */
    return u.i % limit;
}

/* Random double in [0.0, 1.0) */
double random_double() {
    union {
        uint64_t i;
        unsigned char c[sizeof(uint64_t)];
    } u;

    if (!RAND_bytes(u.c, sizeof(u.c))) {
        fprintf(stderr, "Can't get random bytes!\n");
        exit(1);
    }
    /* 53 bits / 2**53 */
    return (u.i >> 11) * (1.0/9007199254740992.0);
}

int main() {
    printf("Dice: %d\n", (int)(random_uint(6) + 1));
    printf("Double: %f\n", random_double());
    return 0;
}
```
