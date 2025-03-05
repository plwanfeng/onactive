# Orochi Network 工具安装与使用指南

## 环境要求
- Go 1.x
- 网络连接

## 安装说明
无需特殊安装步骤，只需确保您的系统已正确配置Go语言环境即可。

### 检查Go环境
1. 打开终端
2. 输入以下命令验证Go是否正确安装：
```bash
go version
```
如果显示Go版本信息，则说明环境配置正确。

## 使用说明

### 运行工具
1. 打开终端，进入项目目录
2. 抓包替换为你自己的cookie
   ![image](https://github.com/user-attachments/assets/e9c972a6-25f1-4526-bf47-86d8011b6aa1)

4. 执行以下命令运行程序：
```bash
go run onactive.go
```

### 使用方法
1. 程序启动后，按照提示输入code参数
2. 每次输入参数后，程序会自动发送请求并显示响应结果
3. 如需退出程序，使用Ctrl+C
![image](https://github.com/user-attachments/assets/1cec13ce-1866-4858-b4ef-cd5f8e9a5cb5)

## 编译方法
如果您想将程序编译成可执行文件，可以使用以下命令：

```bash
go build onactive.go
```

编译完成后会生成一个可执行文件，在Windows系统上是`onactive.exe`，在Unix/Linux系统上是`onactive`。

## 注意事项
1. 确保网络连接正常
2. 确认服务器状态
3. 验证请求参数格式
