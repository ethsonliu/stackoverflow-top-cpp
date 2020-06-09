<https://stackoverflow.com/questions/2602013/read-whole-ascii-file-into-c-stdstring>

## 问题

我需要把一个文件内的所有内容读取到一个 `std::string` 中。

如果是读到 `char[]` 中，那么很方便，

```c++
std::ifstream t;
int length;
t.open("file.txt");      // open input file
t.seekg(0, std::ios::end);    // go to the end
length = t.tellg();           // report location (this is the length)
t.seekg(0, std::ios::beg);    // go back to the beginning
buffer = new char[length];    // allocate memory for a buffer of appropriate dimension
t.read(buffer, length);       // read the whole file into the buffer
t.close();                    // close file handle

// ... Do stuff with buffer here ...
```

但现在我想做同样的事情，但不同的是，需要读到 `std::string` 中。我不想使用循环，也就是下面的代码，

```c++
std::ifstream t;
t.open("file.txt");
std::string buffer;
std::string line;
while(t){
std::getline(t, line);
// ... Append line to buffer and go on
}
t.close()
```

还有其他的办法么？

## 回答

```c++
string get_file_string() {
    std::ifstream ifs("path_to_file");
    return string((std::istreambuf_iterator<char>(ifs)),
                  (std::istreambuf_iterator<char>()));
}

string get_file_string2() {
    ifstream inFile;
    inFile.open("path_to_file"); // open the input file

    stringstream strStream;
    strStream << inFile.rdbuf(); // read the file
    return strStream.str(); // str holds the content of the file
}
```
