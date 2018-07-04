
# 定义form

`登录表单`

```html
<html>
<head>
<title>GopherFace - Sign Up</title>
<link rel="stylesheet" type="text/css" href="/static/css/pure.css" />
<link rel="stylesheet" type="text/css" href="/static/css/gopherface.css" />
</head>
<body>
	<h1>GopherFace - Sign Up</h1>

	<div class="formContainer">
	<form action="/signup" method="POST" class="pure-form pure-form-aligned">
		<fieldset>

			<div class="pure-control-group">
				<label for="username">Username</label>
				<input id="username" type="text" placeholder="Username" name="username" value="{{.Fields.username}}">
				<span id="usernameError" class="pure-form-message-inline">{{.Errors.usernameError}}</span>
			</div>

			<div class="pure-control-group">
				<label for="firstName">First Name</label>
				<input id="firstName" type="text" placeholder="First Name" name="firstName" value="{{.Fields.firstName}}">
				<span id="firstNameError" class="pure-form-message-inline">{{.Errors.firstNameError}}</span>
			</div>

			<div class="pure-control-group">
				<label for="lastName">Last Name</label>
				<input id="lastName" type="text" placeholder="Last Name" name="lastName" value="{{.Fields.lastName}}">
				<span id="lastNameError" class="pure-form-message-inline">{{.Errors.lastNameError}}</span>
			</div>

			<div class="pure-control-group">
				<label for="email">E-mail Address</label>
				<input id="email" type="text" placeholder="E-mail Address" name="email" value="{{.Fields.email}}">
				<span id="emailError" class="pure-form-message-inline">{{.Errors.emailError}}</span>
			</div>

			<div class="pure-control-group">
				<label for="password">Password</label>
				<input id="password" type="password" placeholder="Password" name="password" value="{{.Fields.password}}">
				<span id="passwordError" class="pure-form-message-inline">{{.Errors.passwordError}}</span>
			</div>

			<div class="pure-control-group">
				<label for="name">Confirm Password</label>
				<input id="confirmPassword" type="password" placeholder="Confirm Password" name="confirmPassword" value="{{.Fields.confirmPassword}}">
				<span id="confirmPasswordError" class="pure-form-message-inline">{{.Errors.confirmPasswordError}}</span>
			</div>

			<div class="pure-controls">
				<input class="pure-button pure-button-primary" type="submit" value="Sign Up" />
			</div>

		</fieldset>
	</form>
	</div>
</body>
</html>



```

`路由`

```
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")

```



`处理器`

```
package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume2/section3/gopherfaceform/validationkit"
)

type SignUpForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplaySignUpForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signupform.html", s)
}

func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signupconfirmation.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	PopulateFormFields(r, s)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	// Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	// Check if last name was filled out
	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if passord and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s)
	}

}

// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	// If we reached this point, that indicates that we had a successful form submission.
	// Later, we will include form processing logic here, in this case that would be
	// inserting the information from the form as an entry into the database.

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	s := SignUpForm{}
	s.FieldNames = []string{"username", "firstName", "lastName", "email"}
	s.Fields = make(map[string]string)
	s.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
	// 不同的方式,导航到不同的表单显示页面
		ValidateSignUpForm(w, r, &s)
	default:
		DisplaySignUpForm(w, r, &s)
	}

}

```

# 获取表单数据

`go  doc  http.Rrequest.FormValue`

```
s.FieldNames = []string{"username", "firstName", "lastName", "email"}

```

`检查form`

```go
// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	PopulateFormFields(r, s)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	// Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	// Check if last name was filled out
	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if passord and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
	// 如果错误项大于1 则会显示不同的form表单
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s)
	}

}

```


```
// ProcessSignUpForm -> 用户登录成功后
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	// If we reached this point, that indicates that we had a successful form submission.
	// Later, we will include form processing logic here, in this case that would be
	// inserting the information from the form as an entry into the database.

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

<html>
<head>
<title>GopherFace - Sign Up Confirmation</title>
<link rel="stylesheet" type="text/css" href="/static/css/gopherface.css" />
</head>
<body>
	<h1>GopherFace - Sign Up Confirmation</h1>


	<p>Thank you {{.Fields.firstName}}!</p>
	<p>We have received your form submission!</p>


</body>
</html>




```


`填充form表单`

```go
// 用户登陆后, 返回已经填过的值, 和错误等
func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}
```



# form 安全

```go
package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	ttemplate "text/template"
)

// Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func RenderUnsafeTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := ttemplate.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	w.Header().Set("X-XSS-Protection", "0")
	t.Execute(w, templateData)
}

func GenerateUUID() string {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Println("Encountered the following error when attempting to generate an UUID: ", err)
		return ""
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

```

```go
// ...
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signupform.html", s)
}
//...
```

# 媒体文件上传

`model`

```
// 表情
// All possible mood states.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

// ....


// Map that holds the various mood states with keys to serve as
// aliases to their respective mood states.
var Moods map[string]MoodState
var MoodsEmoji map[MoodState]string

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}

	MoodsEmoji = map[MoodState]string{MoodStateNeutral: "\xF0\x9F\x98\x90", MoodStateHappy: "\xF0\x9F\x98\x8A", MoodStateSad: "\xF0\x9F\x98\x9E", MoodStateAngry: "\xF0\x9F\x98\xA0", MoodStateHopeful: "\xF0\x9F\x98\x8C", MoodStateThrilled: "\xF0\x9F\x98\x81", MoodStateBored: "\xF0\x9F\x98\xB4", MoodStateShy: "\xF0\x9F\x98\xB3", MoodStateComical: "\xF0\x9F\x98\x9C", MoodStateOnCloudNine: "\xF0\x9F\x98\x82"}

}
```

```
package handlers

import (
	"net/http"
	"strconv"

	"github.com/EngineerKamesh/gofullstack/volume2/section3/gopherfaceform/models/socialmedia"
)

type PostForm struct {
	FieldNames  []string
	Fields      map[string]string
	Errors      map[string]string
	Moods       map[string]int
	MoodEmoji   map[string]string
	CurrentMood int
}

var MoodStates map[string]int
var MoodEmoji map[string]string

func DisplayPostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {
	RenderTemplate(w, "./templates/postform.html", p)
}

func DisplayPostPreview(w http.ResponseWriter, r *http.Request, p *PostForm) {
// 没有提交前显示

	moodState, _ := strconv.Atoi(p.Fields["mood"])
	post := socialmedia.NewPost("Anonymous Gopher", socialmedia.MoodState(moodState), p.Fields["caption"], p.Fields["messageBody"], "", "", "", nil)
	RenderTemplate(w, "./templates/socialmediapost.html", post)
}

func PopulatePostFormFields(r *http.Request, p *PostForm) {
// 处理提交的表单
	for _, fieldName := range p.FieldNames {
		p.Fields[fieldName] = r.FormValue(fieldName)
	}

}

func ValidatePostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {
// 检查表单的合法性
	p.CurrentMood, _ = strconv.Atoi(r.FormValue("mood"))

	PopulatePostFormFields(r, p)

	if r.FormValue("caption") == "" {
		p.Errors["captionError"] = "The caption field is required."
	}

	if r.FormValue("messageBody") == "" {
		p.Errors["messageBodyError"] = "The post message body is required."
	}

	if len(p.Errors) > 0 {
		DisplayPostForm(w, r, p)
	} else {
		DisplayPostPreview(w, r, p)
	}

}

func PostPreviewHandler(w http.ResponseWriter, r *http.Request) {

	p := PostForm{}
	p.FieldNames = []string{"caption", "messageBody", "mood"}
	p.Fields = make(map[string]string)
	p.Errors = make(map[string]string)
	p.Moods = MoodStates
	p.MoodEmoji = MoodEmoji
	p.CurrentMood = 0

	switch r.Method {

	case "GET":
		DisplayPostForm(w, r, &p)
	case "POST":
		ValidatePostForm(w, r, &p)
	default:
		DisplayPostForm(w, r, &p)
	}

}

func init() {

	MoodStates = make(map[string]int)
	MoodStates["Neutral"] = int(socialmedia.MoodStateNeutral)
	MoodStates["Happy"] = int(socialmedia.MoodStateHappy)
	MoodStates["Sad"] = int(socialmedia.MoodStateSad)
	MoodStates["Angry"] = int(socialmedia.MoodStateAngry)
	MoodStates["Hopeful"] = int(socialmedia.MoodStateHopeful)
	MoodStates["Thrilled"] = int(socialmedia.MoodStateThrilled)
	MoodStates["Bored"] = int(socialmedia.MoodStateBored)
	MoodStates["Shy"] = int(socialmedia.MoodStateShy)
	MoodStates["Comical"] = int(socialmedia.MoodStateComical)
	MoodStates["On Cloud Nine"] = int(socialmedia.MoodStateOnCloudNine)

	MoodEmoji = make(map[string]string)
	MoodEmoji["Neutral"] = socialmedia.MoodsEmoji[socialmedia.MoodStateNeutral]
	MoodEmoji["Happy"] = socialmedia.MoodsEmoji[socialmedia.MoodStateHappy]
	MoodEmoji["Sad"] = socialmedia.MoodsEmoji[socialmedia.MoodStateSad]
	MoodEmoji["Angry"] = socialmedia.MoodsEmoji[socialmedia.MoodStateAngry]
	MoodEmoji["Hopeful"] = socialmedia.MoodsEmoji[socialmedia.MoodStateHopeful]
	MoodEmoji["Thrilled"] = socialmedia.MoodsEmoji[socialmedia.MoodStateThrilled]
	MoodEmoji["Bored"] = socialmedia.MoodsEmoji[socialmedia.MoodStateBored]
	MoodEmoji["Shy"] = socialmedia.MoodsEmoji[socialmedia.MoodStateShy]
	MoodEmoji["Comical"] = socialmedia.MoodsEmoji[socialmedia.MoodStateComical]
	MoodEmoji["On Cloud Nine"] = socialmedia.MoodsEmoji[socialmedia.MoodStateOnCloudNine]

}

```