package userdata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subhroacharjee/auth/lib/model/user"
	"github.com/subhroacharjee/auth/lib/util"
	"gorm.io/gorm"
)

type RepositoryOptions struct {
	*gorm.DB
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(options RepositoryOptions) user.Repository {
	return UserRepositoryImpl{
		DB: options.DB,
	}
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login implements user.Repository.
func (u UserRepositoryImpl) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload LoginPayload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "invalid body provided", http.StatusBadRequest)
		return
	}
	var user user.User
	result := u.DB.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "invalid email or password", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			fmt.Println(result.Error)
			return
		}
	}
	var correct bool
	if correct, err = util.VerifyPassword(user.Password, payload.Password); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(result.Error)
		return
	}
	if !correct {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	userInfo := user.ToJson()
	accessToken, err := util.GenerateJWTToken(user, util.ACCESS_TOKEN)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	refreshToken, err := util.GenerateJWTToken(user, util.REFRESH_TOKEN)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	userInfo["access_token"] = accessToken
	userInfo["refresh_token"] = refreshToken

	jsonData, err := json.Marshal(userInfo)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type RegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register implements user.Repository.
func (u UserRepositoryImpl) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload RegisterPayload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "invalid body provided", http.StatusBadRequest)
		return
	}

	var newUser user.User
	isOldUser := true
	tx := u.DB.First(&newUser, "email = ?", payload.Email)
	if tx.Error != nil {
		if tx.Error != gorm.ErrRecordNotFound {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			fmt.Println(tx.Error)
			return
		} else {
			isOldUser = false
		}
	}

	if isOldUser {
		http.Error(w, "user already exists", http.StatusForbidden)
		return
	}

	newUser.Name = payload.Name
	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	newUser.Password = *hashedPassword
	newUser.Email = payload.Email
	result := u.DB.Create(&newUser)

	if result.Error != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(result.Error)
		return
	}
	userInfo := newUser.ToJson()
	accessToken, err := util.GenerateJWTToken(newUser, util.ACCESS_TOKEN)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	refreshToken, err := util.GenerateJWTToken(newUser, util.REFRESH_TOKEN)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	userInfo["access_token"] = accessToken
	userInfo["refresh_token"] = refreshToken

	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}
