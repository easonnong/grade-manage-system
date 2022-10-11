## 目录结构

```shell
├─  .gitignore
│  go.mod 
│  go.sum
│  LICENSE
│  main.go 
│  README.md
│  tree.txt
│          
├─api         
├─config 
├─database  
├─log  
├─middleware  
├─model 
├─routes
│      router.go  
├─static
│  ├─admin          
│  └─front  
├─upload   
├─utils 
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web 
    ├─admin             
    └─front
```

```shell
首页
http://localhost:3000
后台管理页面
http://localhost:3000/admin

默认管理员:admin  密码:123456
```

## 实现功能

1. 简单的用户管理权限设置
2. 用户密码加密存储
3. 文章分类自定义
4. 列表分页
5. 图片上传七牛云
6. JWT 认证
7. 自定义日志功能
8. 跨域 cors 设置
9. 文章评论功能

## 技术栈

- golang
    - Gin web framework
    - gorm(v1 && v2)
    - jwt-go
    - scrypt
    - logrus
    - gin-contrib/cors
    - go-playground/validator/v10
    - go-ini
- JavaScript
    - vue
    - vue cli
    - vue router
    - ant design vue
    - vuetify
    - axios
    - tinymce
    - moment
- MySQL 