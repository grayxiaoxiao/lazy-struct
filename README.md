# lazy-struct
自用型偷懒工具，用于生成golang的struct

```
git clone https://github.com/grayxiaoxiao/lazy-struct.git
make
make install
```

## 使用
> -d 指定文件最终存放位置
> -p 指定文件的package
> -n 执行结构体的名字
> -c 区分用于数据映射结构，还是配置映射结构体 (dm:数据映射用，sm:配置映射用)
```
lazy-struct gen -d tests -p business -n Customer -c sm serial_number:string name:string price:float64
```
