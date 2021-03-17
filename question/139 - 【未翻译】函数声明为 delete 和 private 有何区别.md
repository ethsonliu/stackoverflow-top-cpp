<https://stackoverflow.com/questions/18847957/delete-modifier-vs-declaring-function-as-private>

## 问题

```c++
class A 
{
public:
    A (const A&) = delete; 
};
```

```c++
class A 
{
private:
    A (const A&) {}
};
```

上述的两段代码实现的功能一致，但两种写法有何区别？

## 回答

