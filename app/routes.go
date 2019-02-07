package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Persons List",
        "GET",
        "/persons",
        PersonsList,
    },
    Route{
        "Persons Create",
        "POST",
        "/persons",
        PersonCreate,
    },
    Route{
        "Person Read",
        "GET",
        "/persons/{personId}",
        PersonRead,
    },
    Route{
        "Person Update",
        "PUT",
        "/persons/{personId}",
        PersonUpdate,
    },
    Route{
        "Person Delete",
        "DELETE",
        "/persons/{personId}",
        PersonDelete,
    },
}
