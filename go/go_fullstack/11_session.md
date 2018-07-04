
```
// ProcessLoginForm
func ProcessLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {

	authResult := authenticate.VerifyCredentials(e, r.FormValue("username"), r.FormValue("password"))
	fmt.Println("auth result: ", authResult)

	// Successful login, let's create a cookie for the user and redirect them to the feed route
	if authResult == true {

		sessionID := utility.GenerateUUID()
		fmt.Println("sessid: ", sessionID)
		u, err := e.DB.GetUser(r.FormValue("username"))
		if err != nil {
			log.Print("Encountered error when attempting to fetch user record: ", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

        // 检查session
		err = authenticate.CreateSecureCookie(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create secure cookie: ", err)
			http.Redirect(w, r, "/login", 302)
			return

		}

		err = authenticate.CreateUserSession(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create user session: ", err)
			http.Redirect(w, r, "/login", 302)
			return

		}

		http.Redirect(w, r, "/feed", 302)

	} else {

		l.Errors["usernameError"] = "Invalid login."
		DisplayLoginForm(w, r, l)

	}

}
```


`工具类`

```go
package authenticate

import (
	"log"
	"net/http"
	"os"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/models"

	"github.com/gorilla/sessions"
)

// 存储
var SessionStore *sessions.FilesystemStore

// 创建session
func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	gfSession, err := SessionStore.Get(r, "gopherface-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email

    // 写入请求
	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "gopherface-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)
}

func init() {
    // 存储, 获取hashkey
	SessionStore = sessions.NewFilesystemStore("/tmp/gopherface-sessions", []byte(os.Getenv("GOPHERFACE_HASH_KEY")))

}

```


```
package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/common/authenticate"
)

func GatedContentHandler(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		shouldRedirectToLogin := false

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)
		if err != nil {
			log.Print(err)
		}

		//fmt.Printf("secure cookie contents: %+v\n", secureCookieMap)

		// Check if the sid key which is used to store the session id value
		// has been populated in the map using the comma ok idiom
		if _, ok := secureCookieMap["sid"]; ok == true {

            // 获取session
			gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")

			fmt.Printf("gopherface session values: %+v\n", gfSession.Values)
			if err != nil {
				log.Print(err)
				return
			}

			// Check if the session id stored in the secure cookie matches
			// the id and username on the server-side session
			if gfSession.Values["sessionID"] == secureCookieMap["sid"] && gfSession.Values["username"] == secureCookieMap["username"] {
				next(w, r)
			} else {
				shouldRedirectToLogin = true
			}

		} else {
			shouldRedirectToLogin = true

		}

		if shouldRedirectToLogin == true {
			http.Redirect(w, r, "/login", 302)
		}

	})

}

```