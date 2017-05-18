CREATE KEYSPACE question
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

USE question;

CREATE TABLE questions (
    question text,
    response text,
    hits int,
    PRIMARY KEY (question, response)
);
