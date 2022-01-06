CREATE USER dreamit_root WITH password 'qwerty123';

DROP DATABASE IF EXISTS dreamit_api_db;
CREATE DATABASE dreamit_api_db
    WITH OWNER dreamit_root
    ENCODING 'utf8';
GRANT ALL PRIVILEGES ON database dreamit_api_db TO dreamit_root;
\connect dreamit_api_db;

CREATE EXTENSION IF NOT EXISTS citext;

DROP TABLE IF EXISTS rooms CASCADE;
CREATE TABLE rooms (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL
);
GRANT ALL PRIVILEGES ON TABLE rooms TO dreamit_root;

INSERT INTO rooms(id, title) VALUES (1, 'dad');
INSERT INTO rooms(id, title) VALUES (2, 'dad');
INSERT INTO rooms(id, title) VALUES (3, 'dad');

DROP TABLE IF EXISTS intervals CASCADE;
CREATE TABLE intervals (
   id SERIAL NOT NULL PRIMARY KEY,
   start TIMESTAMP(3) WITH TIME ZONE NOT NULL,
   "end" TIMESTAMP(3) WITH TIME ZONE NOT NULL
);
GRANT ALL PRIVILEGES ON TABLE intervals TO dreamit_root;


DROP TABLE IF EXISTS schedules CASCADE;
CREATE TABLE schedules (
   roomId BIGINT NOT NULL REFERENCES rooms(id),
   intervalId INT NOT NULL REFERENCES intervals(id),
   isBooked BOOLEAN NOT NULL
);
GRANT ALL PRIVILEGES ON TABLE schedules TO dreamit_root;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO dreamit_root;

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
   id SERIAL NOT NULL PRIMARY KEY,
   login CITEXT NOT NULL,
   first_name TEXT,
   last_name TEXT,
   email CITEXT,
   password BYTEA NOT NULL DEFAULT E'\\000'::bytea,
   avatar TEXT NOT NULL DEFAULT '',
   background TEXT NOT NULL DEFAULT '',
   CONSTRAINT email_unique UNIQUE (email),
   CONSTRAINT login_unique UNIQUE (login)
);
GRANT ALL PRIVILEGES ON TABLE users TO dreamit_root;
INSERT INTO users(login, email, password) VALUES ('user_login_01', 'dummy_email_01@mail.ru', 'SoMePassWorD2007');
INSERT INTO users(login, email, password) VALUES ('user_login_02', 'dummy_email_02@mail.ru', 'SoMePassWorD2007');
INSERT INTO users(login, email, password) VALUES ('user_login_03', 'dummy_email_03@mail.ru', 'SoMePassWorD2007');
INSERT INTO users(login, email, password) VALUES ('user_login_04', 'dummy_email_04@mail.ru', 'SoMePassWorD2007');

-- List of auth services in app
DROP TABLE IF EXISTS auth_services CASCADE;
CREATE TABLE auth_services (
   id SERIAL NOT NULL PRIMARY KEY,
   service CITEXT NOT NULL,

   CONSTRAINT service_unique UNIQUE (service)
);

INSERT INTO auth_services(service) VALUES ('keycloak');


-- Auth tokens from other app
DROP TABLE IF EXISTS auth_tokens CASCADE;
CREATE TABLE auth_tokens (
     auth_id INTEGER NOT NULL,
     service_id INTEGER NOT NULL,
     access_token TEXT NOT NULL,
     user_id INTEGER NOT NULL,
     refresh_token TEXT,

     PRIMARY KEY (auth_id, service_id, access_token),

     FOREIGN KEY (service_id) REFERENCES auth_services(id),
     FOREIGN KEY (user_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS calendar CASCADE;
CREATE TABLE calendar (
                          id SERIAL NOT NULL PRIMARY KEY,
                          roomId BIGINT NOT NULL REFERENCES rooms(id),
                          title TEXT NOT NULL,
                          start TIMESTAMP(3) WITH TIME ZONE NOT NULL,
                          "end" TIMESTAMP(3) WITH TIME ZONE NOT NULL,
                          author BIGINT NOT NULL REFERENCES users(id)
);
GRANT ALL PRIVILEGES ON TABLE calendar TO dreamit_root;

-- ОСТАВЬ НАДЕЖДУ ВСЯК СЮДА ВХОДЯЩИЙ --

-- справочник статутов инвайта 
DROP TABLE IF EXISTS invite_status_dict CASCADE;
CREATE TABLE invite_status_dict (
   id SERIAL NOT NULL PRIMARY KEY, 
   "status" TEXT NOT NULL
);
GRANT ALL PRIVILEGES ON TABLE invite_status_dict TO dreamit_root;

INSERT INTO invite_status_dict(id, "status") VALUES (0, 'pending');
INSERT INTO invite_status_dict(id, "status") VALUES (1, 'accepted');
INSERT INTO invite_status_dict(id, "status") VALUES (2, 'declined');
INSERT INTO invite_status_dict(id, "status") VALUES (3, 'expired');

-- табличка инвайтов 
-- автором инвайта считаем автора ивента (только он может кидать инвайт)
-- у одного инвайта может быть много получателй 
-- у одного получателя можеть быть много инвайтов 
DROP TABLE IF EXISTS invites CASCADE;
CREATE TABLE invites (
   id SERIAL NOT NULL PRIMARY KEY, 
   eventId BIGINT NOT NULL REFERENCES calendar(id),
   receiverId BIGINT NOT NULL REFERENCES users(id),
   statusId BIGINT NOT NULL REFERENCES invite_status_dict(id)
);
GRANT ALL PRIVILEGES ON TABLE invites TO dreamit_root;

-- табличка нотифаев -- сообщенек участникам ивента --
DROP TABLE IF EXISTS notifies CASCADE;
CREATE TABLE notifies (
   id SERIAL NOT NULL PRIMARY KEY, 
   eventId BIGINT NOT NULL REFERENCES calendar(id),
   "subject" TEXT NOT NULL,  
   "message" TEXT NOT NULL 
);
GRANT ALL PRIVILEGES ON TABLE notifies TO dreamit_root;

-- теги
DROP TABLE IF EXISTS notify_tags CASCADE;
CREATE TABLE notify_tags (
   if SERIAL NOT NULL PRIMARY KEY,
   tag TEXT NOT NULL,
   notifyId BIGINT NOT NULL REFERENCES notifies(id)
);
GRANT ALL PRIVILEGES ON TABLE notify_tags TO dreamit_root;

DROP TABLE IF EXISTS comments CASCADE;
CREATE TABLE comments (
   id SERIAL NOT NULL PRIMARY KEY,
   authorId BIGINT NOT NULL REFERENCES users(id),
   notifyId BIGINT NOT NULL REFERENCES notifies(id),
   "message" TEXT NOT NULL
);
