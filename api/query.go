package main

const (
	qGetRoleId          = "SELECT id FROM user_roles WHERE name='%s';"
	qFieldIDToName      = "SELECT name FROM company_fields WHERE id='%s';"
	qLocationIDToName   = "SELECT name FROM tag_location WHERE id='%s';"
	qSizeIDToName       = "SELECT name FROM tag_size WHERE id='%s';"
	qInvestmentIDToName = "SELECT name FROM tag_investment WHERE id='%s';"
)

const (
	qGetUser            = "SELECT email, firstname, surname, e_eu_applied_before, e_eu_funded_before, e_eu_subsidy_before, e_wrote_projects, e_enlish_fluent, e_edu_writing_projects, e_edu_project_mngmt, e_edu_financial_reporting, e_edu_eu_standards, e_edu_market_emerge FROM users WHERE id='%s';"
	qCreateUser         = "INSERT INTO users (id, username, email, password, role_id, firstname, surname) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')"
	qGetUserUuidAndHash = "SELECT id, password FROM users WHERE username='%s';"
	qGetUserCompanies   = "SELECT id, name FROM companies WHERE owner_id='%s';"
	qGetCompany         = "SELECT name, created_at, employees, field_id, t_size, t_location, t_investment, t_pib, t_mib FROM companies WHERE id='%s';"
)
