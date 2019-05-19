\connect postgres;
DROP DATABASE IF EXISTS konsultant;

CREATE DATABASE konsultant;
\connect konsultant;


CREATE TABLE company_fields (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE tag_location (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE tag_size (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE tag_investment (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE user_roles (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(20) UNIQUE NOT NULL
);

CREATE TABLE users (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    username varchar(20) UNIQUE NOT NULL,
    email varchar(50) UNIQUE NOT NULL,
    password varchar(60) NOT NULL,
    company_id UUID, 
    role_id UUID NOT NULL REFERENCES user_roles (id) ON DELETE CASCADE,
    firstname varchar(20) NOT NULL,
    surname varchar(20) NOT NULL,
    e_eu_applied_before boolean,
    e_eu_funded_before boolean,
    e_eu_subsidy_before boolean,
    e_wrote_projects boolean,
    e_enlish_fluent boolean,
    e_edu_writing_projects boolean,
    e_edu_project_mngmt boolean,
    e_edu_financial_reporting boolean,
    e_edu_eu_standards boolean,
    e_edu_market_emerge boolean
);

CREATE TABLE companies (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name varchar(60) NOT NULL,
    created_at timestamp without time zone,
    employees integer DEFAULT 1,
    field_id UUID NOT NULL REFERENCES company_fields (id),
    t_size UUID REFERENCES tag_size (id) ON DELETE SET NULL,
    t_location UUID REFERENCES tag_location (id) ON DELETE SET NULL,
    t_investment UUID REFERENCES tag_investment (id) ON DELETE SET NULL,
    t_pib integer UNIQUE,
    t_mib integer UNIQUE
);

ALTER TABLE users ADD CONSTRAINT users_company_id_fkey FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE SET NULL;

CREATE TABLE applications (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    valid boolean DEFAULT true,
    title varchar(100) NOT NULL,
    overview text NOT NULL,
    tags UUID[] NOT NULL,
    from_date timestamp without time zone DEFAULT timezone('utc'::text, now()),
    to_date timestamp without time zone DEFAULT timezone('utc'::text, now())
);

CREATE TABLE checks (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    pre_checked boolean DEFAULT false,
    description text NOT NULL,
    assist text NOT NULL
);

CREATE TABLE calls (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    title varchar(100) NOT NULL,
    deadline timestamp without time zone NOT NULL,
    program varchar(20) NOT NULL,
    description text[] NOT NULL,
    included_checks UUID[] NOT NULL
)