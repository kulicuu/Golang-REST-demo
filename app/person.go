package main

// import "time"


type Person struct {
    Id                  int         `json:"id"`
    FirstName           string      `json:"firstname"`
    LastName            string      `json:"lastname"`
    EmailAddress        string      `json:"emailaddress"`
    PhoneNumber         string      `json:"phonenumber"`
    // Created     time.time   `json:"time"`
}


type Persons []Person
