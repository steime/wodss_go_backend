// Helper Functions for Handlers
package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
	"time"
)
// Check if param Student ID matches the Token Student ID
func CheckID(r *http.Request) (bool,string) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := int(claims["sub"].(float64))
		studId := strconv.Itoa(claimId)
		return studId == id ,studId
	} else {
		return false,"-1"
	}
}
// Check if query param Student ID matches the Token Student ID, ID as uint
func CheckBodyID(r *http.Request, id uint) bool {
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := uint(claims["sub"].(float64))
		return claimId == id
	} else {
		return false
	}
}
// Check if query param Student ID matches the Token Student ID, ID as string
func CheckQueryID(r *http.Request, id string) bool {
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := int(claims["sub"].(float64))
		studId := strconv.Itoa(claimId)
		return studId == id
	} else {
		return false
	}
}
// Extract Student ID from Token
func GetStudentIdFromToken(r *http.Request) (string,error) {
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := int(claims["sub"].(float64))
		studId := strconv.Itoa(claimId)
		return  studId ,nil
	} else {
		return "-1", errors.New("unable to extract id from token")
	}
}
// Email Input Validation
func ValidateMail(mail string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(mail)
}
// Log Error and send a bad request
func LogErrorAndSendBadRequest(w http.ResponseWriter,r *http.Request, err error) {
	LogError(err.Error(),r.Method,r.RequestURI,r.Proto, "400")
	w.WriteHeader(http.StatusBadRequest)
}
// Generate Tokenpair for Student, used in refresh handler func
func GenerateTokenPair(studentID uint) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = studentID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = studentID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token":  t,
		"refreshToken": rt,
	}, nil
}
