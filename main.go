package main

import (
	_ "github.com/solcates/go-preso/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"github.com/solcates/go-preso/globalsessions"
	"github.com/prometheus/common/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/beego/social-auth"
	"github.com/beego/social-auth/apps"
	"github.com/astaxie/beego/context"
	"github.com/solcates/go-preso/models"
	"crypto/md5"
	"bytes"
	"fmt"
	"io"
	"time"
)

func IsUserLogin(ctx *context.Context) (int, bool) {
	if id, ok := ctx.Input.CruSession.Get("login_user").(int); ok && id == 1 {
		return id, true
	}
	return 0, false
}

type socialAuther struct {
}

func (p *socialAuther) IsUserLogin(ctx *context.Context) (int, bool) {
	return IsUserLogin(ctx)
}

func (p *socialAuther) LoginUser(ctx *context.Context, uid int) (string, error) {
	// fake login the user
	if uid == 1 {
		ctx.Input.CruSession.Set("login_user", 1)
	}
	return "/login", nil
}

var SocialAuth *social.SocialAuth

func init() {

	/*
	Setup Sessions
	*/
	//orm.RegisterDataBase("default", "sqlite3", "db/database.db", 30)
	//orm.RegisterModel(new(models.User))

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_url"), 30)

	sessionConfig := &session.ManagerConfig{
		CookieName: "jigsessionid",
		Gclifetime: 3600,
	}
	var err error
	globalsessions.GlobalSessions, err = session.NewManager("memory", sessionConfig)
	if err != nil {
		log.Error(err)
	}
	globalsessions.GlobalSessions.SetSecure(true)
	go globalsessions.GlobalSessions.GC()

	// Setup DB for first time
	orm.RunSyncdb("default", false, false)
	_, err = models.GetUserById(1)

	if err != nil {
		md5Password := md5.New()
		io.WriteString(md5Password, "admin")
		buffer := bytes.NewBuffer(nil)
		fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
		newPass := buffer.String()
		now := time.Now()
		default_admin := models.User{
			Username:      "admin",
			Password:      newPass,
			LastLogintime: now,
			Created:       now,
		}
		_, err := models.AddUser(&default_admin)
		if err != nil {
			log.Error(err)
		}
		log.Info("Created default admin account of 'admin' with password of 'admin'... make sure you change it :)")

	}

	// OAuth
	var clientId, secret string

	appURL := beego.AppConfig.String("social_auth_url")
	if len(appURL) > 0 {
		social.DefaultAppUrl = appURL
	}

	clientId = beego.AppConfig.String("github_client_id")
	secret = beego.AppConfig.String("github_client_secret")
	err = social.RegisterProvider(apps.NewGithub(clientId, secret))
	if err != nil {
		beego.Error(err)
	}
	SocialAuth = social.NewSocial("/login/", new(socialAuther))
	SocialAuth.ConnectSuccessURL = "/settings/profile"
	SocialAuth.ConnectFailedURL = "/settings/profile"
	SocialAuth.ConnectRegisterURL = "/register/connect"
	SocialAuth.LoginURL = "/login"
}

func main() {
	//var FilterUser = func(ctx *context.Context) {
	//	if strings.HasPrefix(ctx.Input.URL(), "/login") {
	//		return
	//	}
	//	_, ok := ctx.Input.Session("uid").(int)
	//	if !ok {
	//		ctx.Redirect(302, "/login")
	//	}
	//}
	//
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}
