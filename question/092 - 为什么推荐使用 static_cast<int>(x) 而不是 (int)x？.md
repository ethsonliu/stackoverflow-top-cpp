<https://stackoverflow.com/questions/103512/why-use-static-castintx-instead-of-intx>

## 问题

我经常听见别人推荐用 static_cast 取代 C 语言的强制转换，这是对的么？

## 回答

是对的，C 式的强制转换看不出语义，也不利用编译器的错误检查，具体参考：[static_cast, dynamic_cast, const_cast 和 reinterpret_cast 怎么用？](https://github.com/EthsonLiu/stackoverflow-top-cpp/blob/master/question/010%20-%20static_cast%2C%20dynamic_cast%2C%20const_cast%20%E5%92%8C%20reinterpret_cast%20%E6%80%8E%E4%B9%88%E7%94%A8%EF%BC%9F.md)
