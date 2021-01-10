# 部署说明

1、将项目拉取到本地

2、配置cousul环境

```
一、下载
https://www.consul.io/downloads.html 下载对应系统的文件

二、环境配置
1.将解压出来的consul.exe放在gopath下的bin目录中，因为gopath配置过环境变量，所以不要再配置

2.或者将consul.exe的目录单独配置个环境变量，建议使用第一种即可

三、验证
执行consul version命令，如果版本提示，说明正常
```

3、安装npm

4、启动consul

```
consul agent -dev
```

5、启动微服务

- 启动web目录下的web.exe
- 启动user_srv目录下的user_srv.exe
- 启动product_srv目录下的product_srv.exe
- 启动seckill_srv目录下的seckill_srv.exe

在http://localhost:8500/ui/dc1/services中确认微服务已启动

6、启动vue项目

- 进入vue_gin目录下
- 执行  npm run serve

