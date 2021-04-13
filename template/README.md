# qkgo-cli

### 运行之前
> 本地Go环境配置
> 
> mysql
> 
> [goose](https://bitbucket.org/liamstask/goose/src) 数据库迁移工具
> 
> [api-doc](https://apidocjs.com/) 自动生成文档工具

### 运行

1. 进入到`conf`文件夹中（MacOS, Linux执行以下命令，Windows将`cp`更改为`copy`）
```bash
    cp .database.json.example database.json
    cp .global.json.example global.json
    cp .http.json.example http.json
    cp .log.json.example log.json
```
根据实际情况更改配置参数
2. 将路径该问项目跟路径
```bash
    apidoc -i controller/ -o public/apidoc/
```
3. 将路径该问项目跟路径
```bash
    cp db/dbconf.yml.example db/dbconf.yml
    ## 根据实际数据库情况进行更改
    
    goose up
    ## 根据migrations中的sql文件自动在数据库中生成表
```
4. 启动项目
```bash
    go run main.go
```



