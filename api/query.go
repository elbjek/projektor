package main

const (
	qGetUser            = "SELECT email, firstname, surname FROM users WHERE id='%s';"
	qGetUserUuidAndHash = "SELECT id, password FROM users WHERE username='%s';"
)
