package main

import (
	"fmt"

	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/benmanns/goworker"
	"github.com/shiwork/favpostbot/config"
	"strconv"
)

var confPath = os.Getenv("BOT_CONFIG")
var conf config.BotConfig

func post(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args[0])
	api := anaconda.NewTwitterApi(conf.Token, conf.Secret)

	tweet_id, _ := strconv.ParseInt(args[0].(string), 10, 64)
	_, err := api.Retweet(tweet_id, true)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func init() {
	conf, _ = config.Parse(confPath)
	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)

	goworker.Register("Favpost", post)
}

func main() {
	//	goji.Get("/login", Login)
	//	goji.Get("/login/callback", LoginCallback)

	//	go goji.Serve()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}

/*
var session *sessions.Session

func InitSession(r *http.Request) {
	store := sessions.NewCookieStore([]byte("secret-secret"))
	session, _ = store.Get(r, "favpostbot")
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	InitSession(r)
	callbackURL := "http://" + r.Host + "/login/callback"
	authURL, tempCred, err := anaconda.AuthorizationURL(callbackURL)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	session.Values["temp_twitter_token"] = tempCred.Token
	session.Values["temp_twitter_secret"] = tempCred.Secret

	if err = sessions.Save(r, w); err != nil {
		fmt.Println("Error saving session: %v", err)
	}

	http.Redirect(w, r, authURL, 303)
}

func LoginCallback(c web.C, w http.ResponseWriter, r *http.Request) {
	InitSession(r)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error! %s\n", err), http.StatusInternalServerError)
		return
	}

	fmt.Println(session.Values["temp_twitter_token"])
	verifier := r.Form["oauth_verifier"][0]

	tempCred := &oauth.Credentials{
		Token:  session.Values["temp_twitter_token"].(string),
		Secret: session.Values["temp_twitter_secret"].(string),
	}

	credentials, values, err := anaconda.GetCredentials(tempCred, verifier)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error! %s\n", err), http.StatusInternalServerError)
		return
	}

	if credentials == nil {
		http.Error(w, "Credentials nil", http.StatusInternalServerError)
		return
	}

	fmt.Println(values)

	user_id, _ := strconv.ParseInt(values["user_id"][0], 10, 64)

	// login session
	session.Values["user_id"] = user_id
	if err = sessions.Save(r, w); err != nil {
		fmt.Println("Error saving session: %v", err)
	}

	fmt.Fprintf(w, "credentials, %v!", *credentials)
}
*/
