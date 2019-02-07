package main

// import "time"


type Person struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Address     string      `json:"address"`
    // Created     time.time   `json:"time"`
}


type Persons []Person
