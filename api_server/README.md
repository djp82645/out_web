## api_server
### golang版本 
    1.20.12
### 项目运行
    go mod tidy
    go run main.go
### 代码格式化 使用gofumpt替代gofmt(需要Go 1.20 或更高版本)
    Gofumpt 会执行比 gofmt 更严格的 Go 格式规范。同时确保向后兼容。
    安装命令：
    go install mvdan.cc/gofumpt@latest   
### upx安装 压缩编译后执行文件 
    https://github.com/upx/upx #下载源代码编译放到系统环境变量下
## 打包命令 BuildVersion为构建版本号
    ./setup.bat
## 测试数据
    61057944----http://xxx/api/getcode?token=e49debc93
    61057945----http://xxx/api/getcode?token=e49deb333