
# 登录session

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



```
package authenticate

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/models"

	"github.com/gorilla/securecookie"
)

var hashKey []byte
var blockKey []byte
var s *securecookie.SecureCookie

func CreateSecureCookie(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	} else {
		log.Print(err)
		return err
	}

	return nil

}

func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		return nil, nil
	}
}

func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", 301)
}

func init() {

	hashKey = []byte(os.Getenv("GOPHERFACE_HASH_KEY"))
	blockKey = []byte(os.Getenv("GOPHERFACE_BLOCK_KEY"))

	s = securecookie.New(hashKey, blockKey)
}

```


```
package authenticate

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/models"

	"github.com/gorilla/securecookie"
)

var hashKey []byte
var blockKey []byte
var s *securecookie.SecureCookie

func CreateSecureCookie(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	} else {
		log.Print(err)
		return err
	}

	return nil

}

func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		return nil, nil
	}
}

func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", 301)
}

func init() {
// 可用和不可用key
	hashKey = []byte(os.Getenv("GOPHERFACE_HASH_KEY"))
	blockKey = []byte(os.Getenv("GOPHERFACE_BLOCK_KEY"))

	s = securecookie.New(hashKey, blockKey)
}

```