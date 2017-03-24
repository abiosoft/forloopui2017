CREATE TABLE guestbook (guestName VARCHAR(255), content VARCHAR(255),
                        id SERIAL PRIMARY KEY);
INSERT INTO guestbook (guestName, content) values ('first guest', 'I got here!');
INSERT INTO guestbook (guestName, content) values ('second guest', 'Me too!');
