package main

import (
	_ "github.com/solcates/go-preso/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/solcates/go-preso/models"
	"crypto/md5"
	"bytes"
	"fmt"
	"io"
	"time"
	_ "github.com/astaxie/beego/session/mysql"
	//"github.com/solcates/go-preso/globalsessions"
	//"github.com/astaxie/beego/session"
)

func init() {

	/*
	Setup Database for object storage
	*/

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_url"), 30)

	/*
	Setup Sessions Manager
	*/
	//sessionConfig := &session.ManagerConfig{
	//	CookieName:     "presoid",
	//	Gclifetime:     3600,
	//	CookieLifeTime: 3600,
	//	Secure:         true,
	//	//ProviderConfig: beego.AppConfig.String("mysql_url"),
	//}
	var err error
	//globalsessions.GlobalSessions, err = session.NewManager("memory", sessionConfig)
	//if err != nil {
	//	log.Error(err)
	//}
	//go globalsessions.GlobalSessions.GC()

	/*
	Setup DB for first time
	*/
	orm.RunSyncdb("default", false, false)
	_, err = models.GetUserById(1)

	if err != nil {
		// Create default user of admin with password admin
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

		default_preso := &models.Preso{
			Name:    "Sample_Preso",
			Created: now,
		}
		_, err = models.AddPreso(default_preso)
		if err != nil {
			log.Error(err)
		}
	}

}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("mysql_url")
	beego.BConfig.WebConfig.Session.SessionName = "presoid"
	beego.Run()

}
