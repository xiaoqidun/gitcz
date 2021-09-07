# gitcz
Golang版本Git Commitizen，commit规范工具
# 快速安装
go install github.com/xiaoqidun/gitcz@latest
# 编译安装
```
git clone https://github.com/xiaoqidun/gitcz.git
cd gitcz
go build gitcz.go
```
# 手动安装
1. 根据系统架构下载为你编译好的[二进制文件](https://aite.xyz/product/gitcz/)
2. 将下载好的二进制文件重命名为gitcz并保留后缀
3. 把gitcz文件移动到系统PATH环境变量中的目录下
4. windows外的系统需使用chmod命令赋予可执行权限
# 使用说明
```shell script
# 添加文件到本地仓库
git add .
# 使用-amend参数可覆盖最近一次提交信息，等同于git commit --amend
gitcz
# 推送文件到远程仓库
git push
```
# 规范文档
gitcz使用：[angular git commit规范](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines)