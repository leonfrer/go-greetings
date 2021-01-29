module github.com/go-greeting

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210122235752-a8b976e07c7b // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/go-greeting/conf => ./conf
	github.com/go-greeting/middleware => ./middleware
	github.com/go-greeting/model => ./model
	github.com/go-greeting/pkg/e => ./pkg/e
	github.com/go-greeting/pkg/setting => ./pkg/setting
	github.com/go-greeting/pkg/util => ./pkg/util
	github.com/go-greeting/router => ./router
	github.com/go-greeting/runtime => ./runtime
)
