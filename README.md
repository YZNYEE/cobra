# cobra-test

#### 介绍
cobra学习用cobra做一个目录命令

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  此项目提供一些基本的ls功能
2.  默认使用go run main.go ls默认显示当前目录下的所有文件
#### 参数可选
1.  -d dir 指定目录
2. -u b|k|m|g 指定文件大小单位 b是byte,k为kb,m为MB，g为GB
3.-s n|s|t 指定排序 n我名字首字母排序,s文件大小排序,t修改时间排序
4.-a 是否显示隐藏文件，默认不显示
#### 子命令tree
go run main.go ls tree --L
L为所展示树的深度，默认为1，并继承父命令所有参数，为了方便显示目前只显示文件的大小与姓名
