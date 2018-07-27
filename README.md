# Chunk - 网站 或 http/https 协议 web 应用服务器

###第三方包
- 使用了一些三方现成的包降低了开发难度(仓库里不包含,请自行下载)
    - 路由: gorilla/mux
    - mysql驱动: go-sql-driver
    - redis: garyburd/redigo
    - session: astaxie/beego/session
    - golang.org 的一些库
        - crypto
        - net
        - sys
        - text

### 启动方式    
- 服务器采用命令行形式启动
    - -c : 服务器执行命令, start|restart|stop 
    - -d : 服务器静默模式启动


### 配置文件
- 需要将 conf_example.json(服务器配置文件) 改名为 conf.json 并完成相应的配置才能启动服务器


### 编译
- 基于main.go文件编译服务器的启动文件


### 一些默认的目录设置
- 目录: ./static/ 默认保存所有 以 http://www.domain.com/static/xxx 访问的静态资源
- 目录: ./template/ 保存所有需要渲染的页面模板(模板以html为文件后缀) 


### https
- https 证书不一定放在服务器的目录内(certs目录), 在配置文件内给出证书绝对地址即可


### 开发说明
- 暂无,暂时没空写