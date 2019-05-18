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
    surname varchar(20) NOT NULL
);

CREATE TABLE companies (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name varchar(60) NOT NULL,
    field_id UUID NOT NULL REFERENCES company_fields (id),
    t_size UUID REFERENCES tag_size (id) ON DELETE SET NULL,
    t_location UUID REFERENCES tag_location (id) ON DELETE SET NULL,
    t_investment UUID REFERENCES tag_investment (id) ON DELETE SET NULL
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

