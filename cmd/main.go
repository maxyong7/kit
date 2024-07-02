package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// var (
// 	cfg = viper.New()

// 	webserverCmd = &cobra.Command{
// 		Use: "webserver",
// 		Short: "run the webserver",
// 		Run: func(func(cmd *cobra.Command, args []string)){
// 		},
// 	}
// )

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getuserinfo", GetUserInfo).Methods(http.MethodGet)
	http.Handle("/", r)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	endpoint.GetUserInfo(w, r)
	// ctx := r.Context()

	// fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	// myName := r.PostFormValue("myName")
	// if myName == "" {
	// 	w.Header().Set("x-missing-field", "myName")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))

	http.Error(w, "Unable to GetUserInfo", http.StatusBadRequest)
}
