<https://stackoverflow.com/questions/2354210/can-a-class-member-function-template-be-virtual>

## 问题

类成员模板函数为什么不能被 virtual 修饰？比如下面的代码会编译不通过，

```c++
class Animal{
  public:
      template<typename T>
      virtual void make_sound(){
        //...
      }
};
```

## 回答



