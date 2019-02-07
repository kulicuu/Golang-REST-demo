package main


import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "io"
    "strconv"
    "github.com/gorilla/mux"
)

var currentId int

var persons Persons

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func PersonsList(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(persons); err != nil {
        panic(err)
    }
}

func PersonFind(id int) Person {
    for _, p := range persons {
        if p.Id == id {
            return p
        }
    }
    return Person{}
}

func PersonRead(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var personId int
    var err error
    if personId, err = strconv.Atoi(vars["personId"]); err != nil {
        panic(err)
    }


    person := PersonFind(personId)

    if person.Id > 0 {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(person); err != nil {
            panic(err)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}



func PersonCreate(w http.ResponseWriter, r *http.Request) {
    var person Person
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &person); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }


    // fmt.Println(string(json.Marshal(person)))
    fmt.Printf("%+v\n", person)

    currentId += 1
    person.Id = currentId
    persons = append(persons, person)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(person); err != nil {
        panic(err)
    }




    // fmt.Fprint(w, "Placeholder for POST persons, create a person.")
}
