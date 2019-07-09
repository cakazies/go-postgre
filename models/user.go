package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	gorm.Model
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return map[string]interface{}{"status": "invalid", "message": "Email address format is incorrect"}, false
	}

	if len(user.Password) < 6 {
		return map[string]interface{}{"status": "invalid", "message": "Password is minimum 6 character"}, false
	}

	temp := &User{}

	sql := fmt.Sprintf("SELECT id,email,username FROM users	WHERE email = '%s'", user.Email)
	data, err := DB.Query(sql)
	if err != nil {
		log.Println("error query : ", err)
	}
	for data.Next() {
		err = data.Scan(&user.ID, &user.Email, &user.Username)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
			log.Println(saveError)
		}
	}
	if temp.Email != "" {
		return map[string]interface{}{"status": "invalid", "message": "Email address already in use by another user."}, false
	}

	return map[string]interface{}{"status": "Valid", "message": "Requirement passed"}, true
}

func (user *User) CreateAccount() map[string]interface{} {
	if rsp, status := user.Validate(); !status {
		return rsp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO users (email, username, password, created_at, updated_at, status) VALUES ('%s', '%s', '%s', '%s', '%s', '1'); ",
		user.Email, user.Username, user.Password, now, now)
	_, errs := DB.Query(sql)
	if errs != nil {
		log.Println("yang error adalah insert users", errs)
		return map[string]interface{}{"status": "invalid", "message": "Insert Errors call admin or developer "}
	}

	tk := &Token{UserId: 11}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("sapiperahhijau"))

	return map[string]interface{}{"status": "valid", "message": "Account is successfully created ", "token": tokenString}
}

func (user *User) Login() map[string]interface{} {
	return map[string]interface{}{"status": "invalid", "message": "percobaan "}
	// // registeredAccount := &User{}
	// // err := GetDB().Table("accounts").Where("email = ?", account.Email).First(registeredAccount).Error

	// // if err != nil {
	// // 	if err == gorm.ErrRecordNotFound {
	// // 		return util.MetaMsg(false, "Account is not recognized")
	// // 	}
	// // 	return util.MetaMsg(false, "There is something error")
	// // }

	// // err = bcrypt.CompareHashAndPassword([]byte(registeredAccount.Password), []byte(account.Password))

	// // if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// // 	return util.MetaMsg(false, "Password is Invalid")
	// // }

	// // tk := &Token{UserId: registeredAccount.ID}
	// // token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// // tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_secret")))

	// // registeredAccount.Token = tokenString
	// // registeredAccount.Password = ""

	// // response := util.MetaMsg(true, "Successfully Login")
	// // response["account"] = registeredAccount
	// return response
}
