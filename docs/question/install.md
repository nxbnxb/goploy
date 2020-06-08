# 安装问题

# sql: unknown driver "" (forgotten import?)

原因是缺少 MySQL客户端，需要安装 MySQL客户端

# Error during Websocket handshake

原因是反向代理设置不正确
```nginx
# nginx需要设置升级http协议
proxy_set_header Upgrade         $http_upgrade;
proxy_set_header Connection      "upgrade";
```