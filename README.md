### 说明
这是自己在探究 RPC 原理过程中，自己实现的客户端服务端 RPC 通信原理的 Demo，底层通信使用的是 net 包的 tls 的加密通信

### 使用
- 生成证书

  ./makecert.sh mail@examole.com
  
- 启动服务端
  
  go run server/main.go
  
- 启动客户端
  
  go run client/main.go


