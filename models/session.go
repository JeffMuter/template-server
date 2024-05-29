package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"template-server/internal/database"
)

var sessions = map[string]string{}

type Session struct {
	Id           int
	UserId       int
	SessionToken string
	Created      time.Time
	Expires      time.Time
}

func GenSessionId() (string, error) {
	byteSlice := make([]byte, 32)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(byteSlice), nil
}

func SetSession(email string, w http.ResponseWriter) error {
	sessionToken, err := GenSessionId()
	if err != nil {
		return err
	}

	// add session to db here, then delete session map stuff
	db := database.DatabaseOpen()

	defer db.Close()
	var session Session
	var user User
	user, err = GetUserByEmail(email)
	if err != nil {
		fmt.Println("get user failed in setSession()")
	}
	session.UserId = user.Id
	session.SessionToken = sessionToken
	session.Created = time.Now()
	session.Expires = time.Now().Add(time.Hour * 24)

	err = addession(db, session)
	if err != nil {
		log.Fatal("addSession() err in SetSession()")
	}

	// sessions[sessionId] = email

	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return nil
}

func addession(db *sql.DB, session Session) error {
	query := `INSERT INTO sessions(user_id, session_token, created_at, expires_at) VALUES($1, $2, $3, $4)`
	_, err := db.Exec(query, session.UserId, session.SessionToken, session.Created, session.Expires)
	return err
}

func GetSession(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", err
	}
	userName, exists := sessions[cookie.Value]
	if !exists {
		return "", http.ErrNoCookie
	}
	return userName, nil
}
