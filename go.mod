module github.com/izghua/go-blog

go 1.17

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-errors/errors v1.0.1
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-xorm/xorm v0.7.1
	github.com/gorilla/feeds v1.1.1
	github.com/izghua/zgh v0.0.29
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mojocn/base64Captcha v0.0.0-20190509095025-87c9c59224d8
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/speps/go-hashids v2.0.0+incompatible
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
)

require (
	github.com/Penglq/QLog v0.0.0-20190407122931-240a1c914e20 // indirect
	github.com/astaxie/beego v1.11.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/go-xorm/builder v0.3.2 // indirect
	github.com/go-xorm/core v0.6.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/robfig/cron v1.1.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/ugorji/go v1.1.4 // indirect
	golang.org/x/image v0.0.0-20190501045829-6d32002ffd75 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	qiniupkg.com/x v7.0.8+incompatible // indirect
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday v2.0.0+incompatible

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f

replace gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
