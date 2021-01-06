# echo-demo

docker 编译时 go get 失败的问题
https://blog.csdn.net/asty9000/article/details/107720900
https://www.cnblogs.com/zhangmingcheng/p/12294156.html

go proxy
https://goproxy.cn/
https://goproxy.io/

# golang module

https://golang.org/doc/code.html
https://blog.golang.org/using-go-modules

# jwt
http://jwt.io/  
https://vertx.io/docs/vertx-auth-jwt/java/  
https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go  

# cron
https://github.com/go-co-op/gocron



# remote debug


    远程启动方式
    dlv exec ./main --headless --listen=:2345 --log --api-version=2

    本地 vscode launch.json 配置
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Connect to server",
                "type": "go",
                "request": "attach",
                "mode": "remote",
                "remotePath": "${workspaceFolder}",
                "port": 2345,
                "host": "127.0.0.1"
            }
        ]
    }
    