/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:53
 */
package conf

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/jwt"
	"github.com/izghua/zgh/utils/alarm"
	"github.com/izghua/zgh/utils/backup"
	"github.com/izghua/zgh/utils/hashid"
	"github.com/izghua/zgh/utils/mail"
	"github.com/izghua/zgh/utils/qq_captcha"
	string2 "github.com/izghua/zgh/utils/string"
	"github.com/speps/go-hashids"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
)

var (
	SqlServer   *xorm.Engine
	ZHashId     *hashids.HashID
	CacheClient *redis.Client
	MailClient  *mail.EmailParam
	Cnf         *Conf
	Env         string
)

func DefaultInit() {
	ZLogInit()
	CnfInit()
	DbInit()
	AlarmInit()
	MailInit()
	ZHashIdInit()
	RedisInit()
	JwtInit()
	QCaptchaInit()
	// the customer error code init
	conf.SetMsg(Msg)
	//BackUpInit()
}

func ZLogInit() {
	zog := new(zgh.ZLogParam)
	fileName := zog.SetFileName("go-blog")
	err := zog.ZLogInit(fileName)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	return
}

func DbInit() {
	sp := new(conn.Sp)
	dbUser := sp.SetDbUserName(Cnf.DbUser)
	dbPwd := sp.SetDbPassword(Cnf.DbPassword)
	dbPort := sp.SetDbPort(Cnf.DbPort)
	dbHost := sp.SetDbHost(Cnf.DbHost)
	dbdb := sp.SetDbDataBase(Cnf.DbDataBase)
	sqlServer, err := conn.InitMysql(dbUser, dbPwd, dbPort, dbHost, dbdb)
	SqlServer = sqlServer
	if err != nil {
		zgh.ZLog().Error("some errors", err.Error())
		panic(err.Error())
	}
	return
}

func BackUpInit() {
	bp := new(backup.BackUpParam)
	dest := "./zip/" + time.Now().Format("2006-01-02") + ".zip"
	backu := bp.SetFilePath(Cnf.BackUpFilePath).
		SetFiles("./backup", "./static/uploads/images").
		SetDest(dest).SetCronSpec(Cnf.BackUpDuration)
	data := make(map[string]string)
	data[time.Now().Format("2006-01-02")+".zip"] = dest
	bp.Ep = MailClient
	subject := time.Now().Format("2006-01-02") + "备份邮件"
	bp.Ep.SetSubject(mail.EmailType(subject)).SetAttaches(data).SetBody(mail.EmailType(
		`<html><body>
		<p><img src="https://golang.org/doc/gopher/doc.png"></p><br/>
		<h1>日常备份.</h1>
		` + string2.RandString(10) + `
		</body></html>`)).SetTo(mail.EmailType(Cnf.BackUpSentTo))
	err := backu.Backup()
	if err != nil {
		zgh.ZLog().Error("message", "backup has error", "error", err.Error())
	} else {
		zgh.ZLog().Info("message", "Congratulations for backup")
	}
	return
}

func AlarmInit() {
	a := new(alarm.AlarmParam)
	alarmT := a.SetType(alarm.AlarmType(Cnf.AlarmType))
	mailTo := a.SetMailTo("xzghua@gmail.com")
	err := a.AlarmInit(alarmT, mailTo)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	return
}

func MailInit() {
	m := new(mail.EmailParam)
	mailUser := m.SetMailUser(mail.EmailType(Cnf.MailUser))
	mailPwd := m.SetMailPwd(mail.EmailType(Cnf.MailPwd))
	mailHost := m.SetMailHost(mail.EmailType(Cnf.MailHost))
	mails, err := m.MailInit(mailPwd, mailHost, mailUser)
	if err != nil {
		zgh.ZLog().Error("message", err.Error())
	} else {
		MailClient = mails
		zgh.ZLog().Info("message", "begin to backup")
		BackUpInit()
		return
	}
	MailClient = mails
	return
}

func ZHashIdInit() {
	hd := new(hashid.HashIdParams)
	salt := hd.SetHashIdSalt(Cnf.HashIdSalt)
	hdLength := hd.SetHashIdLength(Cnf.HashIdLength)
	zHashId, err := hd.HashIdInit(hdLength, salt)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	ZHashId = zHashId
	return
}

func RedisInit() {
	rc := new(conn.RedisClient)
	addr := rc.SetRedisAddr(Cnf.RedisAddr)
	pwd := rc.SetRedisPwd(Cnf.RedisPwd)
	db := rc.SetRedisDb(Cnf.RedisDb)
	client, err := rc.RedisInit(addr, db, pwd)
	if err != nil {
		zgh.ZLog().Error(err.Error())
		panic(err.Error())
	}
	CacheClient = client
	return
}

func JwtInit() {
	jt := new(jwt.JwtParam)
	ad := jt.SetDefaultAudience(Cnf.JwtAudience)
	jti := jt.SetDefaultJti(Cnf.JwtJti)
	iss := jt.SetDefaultIss(Cnf.JwtIss)
	sk := jt.SetDefaultSecretKey(Cnf.JwtSecretKey)
	rc := jt.SetRedisCache(CacheClient)
	tl := jt.SetTokenLife(time.Hour * time.Duration(Cnf.JwtTokenLife))
	_ = jt.JwtInit(ad, jti, iss, sk, rc, tl)
	return
}

func QCaptchaInit() {
	qc := new(qq_captcha.QQCaptcha)
	aid := qc.SetAid(Cnf.QCaptchaAid)
	sk := qc.SetSecretKey(Cnf.QCaptchaSecretKey)
	_ = qc.QQCaptchaInit(aid, sk)
	return
}

func CnfInit() {
	cf := &Conf{
		AppUrl:                "http://www.lengleng.cloud",
		AppImgUrl:             "http://www.lengleng.cloud/static/uploads/images/",
		DefaultLimit:          "20",
		DefaultIndexLimit:     "10",
		DbUser:                "root",
		DbPassword:            "Blzg-301a",
		DbPort:                "3306",
		DbDataBase:            "lengleng-blog",
		DbHost:                "mysql",
		AlarmType:             "mail,wechat",
		MailUser:              "nengg@foxmail.com",
		MailPwd:               "GQ13364077664",
		MailHost:              "nengg@foxmail.com",
		HashIdSalt:            "i must add a salt what is only for me",
		HashIdLength:          8,
		JwtIss:                "go-blog",
		JwtAudience:           "blog",
		JwtJti:                "go-blog",
		JwtSecretKey:          "go-blog",
		JwtTokenLife:          3,
		RedisAddr:             "redis",
		RedisPwd:              "redis",
		RedisDb:               12,
		QCaptchaAid:           "",
		QCaptchaSecretKey:     "**",
		BackUpFilePath:        "./backup/",
		BackUpDuration:        "* * */1 * *",
		BackUpSentTo:          "xzghua@gmail.com",
		DataCacheTimeDuration: 720,
		ImgUploadUrl:          "http://www.lengleng.cloud/console/post/imgUpload",
		ImgUploadDst:          "./static/uploads/images/",
		ImgUploadBoth:         true, // img will upload to qiniu and your server local
		QiNiuUploadImg:        false,
		QiNiuHostName:         "",
		QiNiuAccessKey:        "",
		QiNiuSecretKey:        "",
		QiNiuBucket:           "go-blog",
		QiNiuZone:             "HUABEI",
		CateListKey:           "all:cate:sort",
		TagListKey:            "all:tag",
		Theme:                 0,
		Title:                 "默认Title",
		Keywords:              "默认关键词,Lengleng",
		Description:           "",
		RecordNumber:          "000-0000",
		UserCnt:               2,
		Author:                "Lengleng",
		Email:                 "",
		PostIndexKey:          "index:all:post:list",
		TagPostIndexKey:       "index:all:tag:post:list",
		CatePostIndexKey:      "index:all:cate:post:list",
		LinkIndexKey:          "index:all:link:list",
		SystemIndexKey:        "index:all:system:list",
		PostDetailIndexKey:    "index:post:detail",
		ArchivesKey:           "index:archives:list",
		GithubName:            "lengleng_blog",
		GithubRepo:            "",
		GithubClientId:        "1481d4a1a2d763eb5a61",
		GithubClientSecret:    "f16c12726adbb925447fde25de483a480c474eb0",
		GithubLabels:          "Gitalk",
		ThemeJs:               "/static/home/assets/js",
		ThemeCss:              "/static/home/assets/css",
		ThemeImg:              "/static/home/assets/img",
		ThemeFancyboxCss:      "/static/home/assets/fancybox",
		ThemeFancyboxJs:       "/static/home/assets/fancybox",
		ThemeHLightCss:        "/static/home/assets/highlightjs",
		ThemeHLightJs:         "/static/home/assets/highlightjs",
		ThemeShareCss:         "/static/home/assets/css",
		ThemeShareJs:          "/static/home/assets/js",
		ThemeArchivesJs:       "/static/home/assets/js",
		ThemeArchivesCss:      "/static/home/assets/css",
		ThemeNiceImg:          "/static/home/assets/img",
		ThemeAllCss:           "/static/home/assets/css",
		ThemeIndexImg:         "/static/home/assets/img",
		ThemeCateImg:          "/static/home/assets/img",
		ThemeTagImg:           "/static/home/assets/img",
	}

	files, _ := filepath.Glob("./env.*.yaml")
	dev := false
	prod := false
	for _, v := range files {
		switch v {
		case "env.dev.yaml":
			dev = true
		case "env.prod.yaml":
			prod = true
		default:
			continue
		}
	}

	var fileName string
	var env string
	if dev {
		fileName = "/env.dev.yaml"
		env = "dev"
	} else if prod {
		fileName = "/env.prod.yaml"
		env = "prod"
	} else {
		fileName = "default"
		env = "dev"
	}

	if fileName == "default" {
		Cnf = cf
		Env = env
		return
	}

	res, err := filepath.Abs(filepath.Dir("./main.go"))
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}

	//读取yaml配置文件
	yamlFile, err := ioutil.ReadFile(res + fileName)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &cf)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}

	Cnf = cf
	Env = env
	return
}
