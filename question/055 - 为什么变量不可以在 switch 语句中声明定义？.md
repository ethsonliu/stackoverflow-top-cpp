<https://stackoverflow.com/questions/92396/why-cant-variables-be-declared-in-a-switch-statement>

## 问题

就比如下面的代码，

```c++
switch (val)  
{  
case VAL:  
  // This won't work
  int newVal = 42;  
  break;
case ANOTHER_VAL:  
  ...
  break;
}  
```

会报如下的错，

```
initialization of 'newVal' is skipped by 'case' label
```

为什么会这样？

## 回答

`case` 语句其实就是标签（label），就像 goto 语句那样，解决这个问题其实很简单，只需加一对大括号，以表明作用域即可，

```c++
switch (val)
{   
case VAL:  
{
  // This will work
  int newVal = 42;  
  break;
}
case ANOTHER_VAL:  
...
break;
}
```
