package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"bytes"
	"fmt"
	"io"
	"time"
	"regexp"
	"github.com/solcates/go-preso/models"
	"github.com/solcates/go-preso/globalsessions"
	"github.com/prometheus/common/log"
	"encoding/json"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) Post() {

	// Read the request
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	username := ob.Username
	password := ob.Password

	// Check the password.
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	now := time.Now()

	log.Info(username)
	log.Info(password)
	userInfo, err := models.GetUserInfo(username)
	if err != nil {
		return
	}
	if userInfo.Password == newPass {
		var users models.User
		users.LastLogintime = now
		models.UpdateUserById(&users)

		//Set the session successful login
		sess, _ := globalsessions.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		sess.Set("uid", userInfo.Id)
		sess.Set("uname", userInfo.Username)

		this.Ctx.Redirect(302, "/")
	} else {
		this.Data["JSON"] = &userInfo
		this.ServeJSON()
	}
}


type RegController struct {
	beego.Controller
}

func (c *RegController) URLMapping() {
	c.Mapping("Post", c.Post)
}

//Registration process
func (this *RegController) Post() {
	this.TplName = "reg.tpl"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	usererr := checkUsername(username)
	fmt.Println(usererr)
	if usererr == false {
		this.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := checkPassword(password)
	if passerr == false {
		this.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	now := time.Now()

	userInfo, _ := models.GetUserInfo(username)

	if userInfo.Username == "" {
		var users models.User
		users.Username = username
		users.Password = newPass
		users.Created = now
		users.LastLogintime = now
		models.AddUser(&users)

		//Set the session successful login
		sess, _ := globalsessions.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		sess.Set("uid", userInfo.Id)
		sess.Set("uname", userInfo.Username)
		this.Ctx.Redirect(302, "/")
	} else {
		this.Data["UsernameErr"] = "User already exists"
	}
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}
