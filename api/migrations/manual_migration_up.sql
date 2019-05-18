CREATE TABLE company_fields (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE companies (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(60) NOT NULL,
    field_id UUID NOT NULL REFERENCES company_fields (id),
    tags UUID[] NOT NULL
);

CREATE TABLE tags (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name varchar(40) UNIQUE NOT NULL 
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
    company_id UUID REFERENCES companies (id), 
    role_id UUID REFERENCES user_roles (id) ON DELETE CASCADE NOT NULL,
    firstname varchar(20) NOT NULL,
    surname varchar(20) NOT NULL
);

CREATE TABLE applications (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    valid boolean DEFAULT true,
    title varchar(100) NOT NULL,
    overview text NOT NULL,
    tags UUID[] NOT NULL,
    from_date timestamp without time zone DEFAULT timezone('utc'::text, now()),
    to_date timestamp without time zone DEFAULT timezone('utc'::text, now())
);
