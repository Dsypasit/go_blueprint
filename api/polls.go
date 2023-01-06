package main

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type poll struct {
	ID      bson.ObjectId  `json:"id" bson:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results"`
	APIKey  string         `json:"api_key"`
}

func (s *Server) handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handlePollsGet(w, r)
		return
	case "POST":
		s.handlePollsPost(w, r)
		return
	case "DELETE":
		s.handlePollsDelete(w, r)
		return
	}
	respondHTTPErr(w, r, http.StatusNotFound)
}

func (s *Server) handlePollsGet(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implmented"))
}

func (s *Server) handlePollsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implmented"))
}

func (s *Server) handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implmented"))
}
