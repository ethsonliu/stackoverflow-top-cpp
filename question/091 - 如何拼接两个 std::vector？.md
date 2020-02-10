<https://stackoverflow.com/questions/201718/concatenating-two-stdvectors>

## 问题

如何拼接两个 std::vector？

## 回答

```c++
vector1.insert(vector1.end(), vector2.begin(), vector2.end());
```
