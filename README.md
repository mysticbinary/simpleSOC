# 介绍
一个简单且小型的Golang WEB，前后端分离架构，一个学习DEMO；

运行效果图：
![avatar](./image/simpleSoc.png)


# foresoc 
参考vue-element-admin 、 element UI 来完成的前端样式。

开发、调试：
```
# 进入项目目录
cd vue-element-admin

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run dev
//浏览器访问 http://localhost:9527
```

部署：
```
# 构建测试环境
npm run build:stage

# 构建生产环境
npm run build:prod
```

# ginsoc
使用Gin封装的后端服务，主要是提供API接口。

开发、调试：
```
go run ./main.go
```

