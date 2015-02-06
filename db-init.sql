-- execute inside sqlite console
--.read db-init.sql
--or outside
--sqlite3 sqlite.db < db-init.sql

--turn on foreign keys
PRAGMA foreign_keys = ON;

--create tables
CREATE TABLE IF NOT EXISTS 'Contact' ('cid' INT UNIQUE NOT NULL, 'phone_num' TEXT NOT NULL, 'status' INT, 'first_name' TEXT, 'last_name' TEXT);
CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'email' TEXT NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT UNIQUE NOT NULL, 'password' TEXT NOT NULL, PRIMARY KEY(phone_num,email,user_id));
CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(user_id),PRIMARY KEY(contact_ref,contact_id));
CREATE TABLE IF NOT EXISTS 'Memo' ('id' INT NOT NULL, 'sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' INT NOT NULL, FOREIGN KEY(sender_id) REFERENCES User(user_id), PRIMARY KEY(id), FOREIGN KEY(recipient_id) REFERENCES Contact(cid));
CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' INT NOT NULL, 'user_id' INT, PRIMARY KEY(session_id,user_id));

INSERT INTO 'User' VALUES ("412-445-3171","thesnowmancometh@gmail.com","Robbie","McKinstry","1","foobar");
INSERT INTO 'Session' VALUES ("1","20150504111515","1");
INSERT INTO 'Contact_Reference' VALUES ("1","1");
INSERT INTO 'Contact' VALUES ("1","724-321-5520","2","Justin","Rushin III");
INSERT INTO 'Memo' VALUES ("1","1","1","lol","201502010111515");

INSERT INTO 'Contact_Reference' VALUES ("1","2");
INSERT INTO 'Contact' VALUES ("2","724-321-5521","2","Jason","Rustin IV");
