-- +goose Up
CREATE TABLE IF NOT EXISTS roles(
    id bigserial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL , 
    description text ,
    level INT NOT NULL DEFAULT 0
);

INSERT INTO ROLES(NAME,description,level )VALUES(
    'user','A user can create post or update or comment',1)
;
INSERT INTO ROLES(NAME,description,level )VALUES(
    'hr','Hr can  post review on users and all of the user privileges ',2);

    INSERT INTO ROLES(NAME,description,level )VALUES(
    'admin','An admin can create , edit or delete any post, comment and a,, privilege ',3);

-- +goose Down 
DROP TABLE IF EXISTS roles;
