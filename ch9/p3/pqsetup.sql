DROP TABLE IF EXISTS post;
CREATE TABLE post (
  ID serial,
  TITLE varchar(40),
  CONTENT varchar(255),
  CONSTRAINT pk_post PRIMARY KEY(ID)
);
SELECT * FROM post;

INSERT INTO post(ID,TITLE,CONTENT) VALUES (1,"Title 1","Content 1"), (2,"Title 2","Content 2")