# GO数据类型

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526153203.png)

## 整型int

默认和操作系统的位数有关



## 浮点数

由符号位 + 指数位 + 尾数位组成，尾数部分会丢失，造成精度损失。要保存高精度的数应该选用float64

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526160753.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526161252.png)

## 字符型

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526165406.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526165447.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526162612.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526170330.png)

## bool类型

![image-20210526170652877](/Users/mr.x/Library/Application Support/typora-user-images/image-20210526170652877.png)

一字节为最小单位

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526171053.png)

## string 类型

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526171942.png)

## 零值

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526172736.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526173243.png)

## 类型转换

没有自动转换，只能显示转换

语法 

`var n1 float32 = float32(i)`

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526174902.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526175011.png)

## string类型和基本数据类型的转换

基本类型转string两种方式

- `fmt.Sprintf("%参数"，表达式)`

  ![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526180603.png)

- 

- strconv包的函数

  ![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526180915.png)

![](https://raw.githubusercontent.com/SeaSoonKeun/Picture/main/Blog_Pic/20210526182102.png)