package main

const (
	qGetUser            = "SELECT * FROM users WHERE "
	qGetUserUuidAndHash = "SELECT id, password FROM users WHERE username='%s';"
)
