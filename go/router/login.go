package router

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	config "allyapps.com/memories/config"
	"context"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Level    string             `bson:"level"`
	Token    string             `bson:"token"`
}

type Token struct {
	Accesstoken string `json:"access_token"`
	Expires string `json:"expires_in"`
	Tokentype string `json:"token_type"`
}

// Login
func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		data := User{}
		json.NewDecoder(r.Body).Decode(&data)

		switch r.Method {
		case http.MethodPost:			
			// Validate user
			data, err := validateUser(data)
			if err == nil {
				// Generate Autho0 Token
				token := generateToken()
				data.Token = token.Accesstoken
				// w.Write([]byte(tokenString))
				json.NewEncoder(w).Encode(data)
			}else{
				w.Write([]byte("Access Denied"))
			}
		case http.MethodGet:
		default:
			w.Write([]byte("Wrong Method"))
		}
	})
}

func generateToken() *Token {
	godotenv.Load("/var/www/memories.allyapps.com/memories.env")
	url := os.Getenv("AUTH0URL")

	data := &Token{}
	payload := strings.NewReader("{\"client_id\":\""+os.Getenv("AUTH0CLIENTID")+"\",\"client_secret\":\"" + os.Getenv("AUTH0CLIENTSECRET") + "\",\"audience\":\"" + os.Getenv("AUTH0AUDIENCE") + "\",\"grant_type\":\"client_credentials\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	json.Unmarshal([]byte(string(body)), data)

	fmt.Println(data.Accesstoken)
	defer res.Body.Close()

	return data
}

func validateUser(user User) (User, error) {
	collection := config.Database.Collection("user")

	u := user.Email
	p := user.Password
	
	var data User

	// Search database by email
	filter := bson.M{"email": u}
	err := collection.FindOne(context.Background(), filter).Decode(&data)

	if err != nil {
		return data, err
	} else {
		// Compare the user password and the database password if the same
		errf := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(p))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
			fmt.Println("access denied")
			return data, errf
		} else {
			fmt.Println("access granted")
			return data, nil
		}
	}
}

// We use this for create user testing.
func createUser() {
	collection := config.Database.Collection("user")

	user := User{
		Email:    "admin",
		Password:  os.Getenv("TESTPASSWORD"),
	}

	// Encrypt password
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pass)

	// Insert new user to database
	collection.InsertOne(context.Background(), user)
	fmt.Print(user)
}