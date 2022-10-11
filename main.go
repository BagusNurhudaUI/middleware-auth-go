package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// res, err := http.Get("https://randomuser.me/api")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(res.Body)
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// defer res.Body.Close()
	// sb := string(body)
	// log.Println (sb)

	http.HandleFunc("/student", ActionStudent)
	server := new(http.Server)
	server.Addr = ":9000"
	fmt.Println("server started at localhost: ", server.Addr)
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return

	}
	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
