<https://stackoverflow.com/questions/201718/concatenating-two-stdvectors>

## 问题

如何拼接两个 std::vector？

## 回答

```c++
// vector2 拷贝到 vector1
vector1.insert(vector1.end(), vector2.begin(), vector2.end());

// vector2 移动到 vector1，此时 vector2 不可再用
vector1.insert(vector1.end(), std::make_move_iterator(vector2.begin()), std::make_move_iterator(vector2.end()));
```
