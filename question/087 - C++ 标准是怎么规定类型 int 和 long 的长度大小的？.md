<https://stackoverflow.com/questions/589575/what-does-the-c-standard-state-the-size-of-int-long-type-to-be>

## 问题

 C++ 标准是怎么规定类型 int 和 long 的长度大小的？

## 回答

C++ 标准并没有规定它们的固定大小，只规定了下限。

```
sizeof(char) == 1
sizeof(char) <= sizeof(short) <= sizeof(int) <= sizeof(long) <= sizeof(long long)

sizeof(signed char)   == 1
sizeof(unsigned char) == 1

sizeof(short)     >= 2
sizeof(int)       >= 2
sizeof(long)      >= 4
sizeof(long long) >= 8
```

