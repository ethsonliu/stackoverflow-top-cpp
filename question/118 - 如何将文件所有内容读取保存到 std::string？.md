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

对此有篇文章写得很好，参见 <http://insanecoding.blogspot.com/2011/11/how-to-read-in-file-in-c.html>，

```c++
#include <string>
#include <cstdio>

std::string get_file_contents(const char *filename)
{
  std::FILE *fp = std::fopen(filename, "rb");
  if (fp)
  {
    std::string contents;
    std::fseek(fp, 0, SEEK_END);
    contents.resize(std::ftell(fp));
    std::rewind(fp);
    std::fread(&contents[0], 1, contents.size(), fp);
    std::fclose(fp);
    return(contents);
  }
  throw(errno);
}
```

或者

```c++
#include <fstream>
#include <string>

std::string get_file_contents(const char *filename)
{
  std::ifstream in(filename, std::ios::in | std::ios::binary);
  if (in)
  {
    std::string contents;
    in.seekg(0, std::ios::end);
    contents.resize(in.tellg());
    in.seekg(0, std::ios::beg);
    in.read(&contents[0], contents.size());
    in.close();
    return(contents);
  }
}
```

皆可。
