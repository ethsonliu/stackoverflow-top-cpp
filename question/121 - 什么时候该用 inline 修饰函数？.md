<https://stackoverflow.com/questions/1759300/when-should-i-write-the-keyword-inline-for-a-function-method>

## 问题

什么时候该用 inline 函数？具体有以下几个问题，

1. 什么时候不应该用 inline 函数？
2. 怎么让编译器不去 inline 函数？
3. 如果一个 inline 函数被多个线程调用，会有性能上的影响么？

## 回答

先介绍下何谓 inline 函数，即内联函数。

inline 和宏定义 #define 的作用基本类似，都是替换或者展开。在程序编译阶段，如果遇到内联函数，则将内联函数的实现在当前位置展开。内联的目的是为了减少函数的调用开销，从而提高运行效率，但会增加代码体量。也就是说，对内联函数进行任何修改，都需要重新编译调用该函数的所有文件代码，因为编译器需要重新更换一次这个内联函数，否则将会继续使用旧的函数。

**注意: 内联只是一种建议，并不要求编译器必须执行。如果内联函数本身开销较大（比如含有 for、switch、递归等），编译器可能拒绝内联展开。再者，现代编译器在函数内联的决策处理会比人类手写来的更准确。**我个人只在类似以下这种情况显示 inline：

```c++
class Example
{
public:
    inline std::string getName() { return m_name; }
    inline void setName(const std::string& name) { m_name = name; }
private:
    std::string m_name;
};
```

接着再回答你上述的提问，

1. 什么时候该用 inline 函数？
  
    如果这个函数的定义也放在头文件，那么你应该用 inline 修饰它。

2. 什么时候不应该用 inline 函数？

    函数执行时间可能较长，比如含有 for、switch、递归等。

3. 怎么让编译器不去 inline 函数？

    

4. 如果一个 inline 被多个线程调用，会有性能上的影响么？





