package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Starting app")

	userMap = make(map[string]User)
	config, _ = LoadConfig("config.json")
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	r.HandleFunc("/session/create", GenerateHMAC).
		Methods("GET")
	// HeadersRegexp("Authorization", "Basic.*")

	r.HandleFunc("/session/{username}/stats", UserStats).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

var config Config
var userMap map[string]User

// var allowedRoutesMap map[string][]string

//UserStats Returns attempted login stats for user
func UserStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	un := vars["username"]
	if un == "" {
		http.Error(w, "Not a valid path", http.StatusNotFound)
	}

	if user, ok := userMap[un]; !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		statsResult := Stats{user.Username, user.AuthorizedAttempt, user.UnauthorizedAttempt}
		result, _ := json.Marshal(statsResult)
		w.Write(result)
	}

}

//GenerateHMAC builds the MAC token for the provided user if they exist
func GenerateHMAC(w http.ResponseWriter, r *http.Request) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "Authorization failed", http.StatusUnauthorized)
		return
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		http.Error(w, "Authorization failed", http.StatusUnauthorized)
		return
	}

	if _, ok := userMap[pair[0]]; !ok {
		//User not found 404
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := userMap[pair[0]]

	userHash := []byte(user.PasswordHash)
	hashPass := []byte(pair[1])

	if bcrypt.CompareHashAndPassword(userHash, hashPass) != nil {
		//Wrong password 403
		user.UnauthorizedAttempt++
		userMap[user.Username] = user
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user.AuthorizedAttempt++
	userMap[user.Username] = user
	sharedKeyBytes := []byte(config.SharedKey)

	hmacToken := generateHMAC([]byte(user.Username), sharedKeyBytes)

	w.Write([]byte(hmacToken))
	w.WriteHeader(http.StatusOK)
}

func generateHMAC(un []byte, key []byte) string {
	minutes := 10
	today := time.Now()
	expires := today.Add(time.Duration(time.Minute) * time.Duration(minutes))
	message := fmt.Sprintf("%v%v%v", expires.Format("2020-12-31"), un, time.Now())
	mac := CalcMAC(message, key)
	message += "|"
	encode := append([]byte(message), mac...)
	b64 := base64.StdEncoding.EncodeToString(encode)

	return b64
}

//CalcMAC encrypts message and salt
func CalcMAC(message string, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(message))
	return mac.Sum(nil)
}

//LoadConfig parses a config file at given path, into the Config below
func LoadConfig(path string) (Config, error) {
	fmt.Println("Loading config . . .")
	var conf Config
	configFile, err := os.Open(path)
	if err != nil {
		return conf, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&conf)

	BuildUserMap(conf.Users)
	//TODO BuildAllowedRoutesMap(conf.AllowedRoutes)

	return conf, err
}

//BuildUserMap used for user lookup in app
func BuildUserMap(users []User) {
	for _, user := range users {
		if _, ok := userMap[user.Username]; !ok {
			var u User
			u.Username = user.Username
			u.PasswordHash = user.PasswordHash
			userMap[user.Username] = u
		}
	}
}

// func BuildAllowedRoutesMap(allowedRoutes []Route) {
//		TODO #4
// }

//HealthCheckHandler responds with simple status of service
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

//Config is the struct representing our config.json data
type Config struct {
	SharedKey     string  `json:"shared_key"`
	Users         []User  `json:"users"`
	AllowedRoutes []Route `json:"allowed_routes"`
}

//User holds basic encrypted creds
type User struct {
	Username            string `json:"username"`
	PasswordHash        string `json:"password_hash"`
	AuthorizedAttempt   int    `json:"authorized_attempt"`
	UnauthorizedAttempt int    `json:"unauthorized_attempt"`
}

//Route is a url with array of allowed(?) users
type Route struct {
	Destination string   `json:"destination"`
	Users       []string `json:"users"`
}

//Stats used for returning data for stats endpoint
type Stats struct {
	Name    string `json:"name"`
	Auths   int    `json:"auths"`
	Unauths int    `json:"unauths"`
}
