package main

const (
	qGetUser            = "SELECT email, firstname, surname FROM users WHERE id='%s';"
	qGetUserUuidAndHash = "SELECT id, password FROM users WHERE username='%s';"
	qGetUserCompanies   = "SELECT id, name FROM companies WHERE owner_id='%s';"
	qGetCompany         = "SELECT name, field_id, tags FROM companies WHERE id='%s';"
)
