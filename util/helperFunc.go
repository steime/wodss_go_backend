package util

import (
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
	"strconv"
)

func CheckID(r *http.Request) (bool,string) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	tk := ctx.Value("user")
	token := tk.(*persistence.Token)
	studId := strconv.Itoa(int(token.StudentID))
	return studId == id ,studId
}
