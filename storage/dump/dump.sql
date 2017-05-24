CREATE KEYSPACE sentence
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE sentence;

CREATE TABLE sentences (
    sentence text PRIMARY KEY
);

CREATE TABLE conversations (
    question text,
    response text,
    hits int,
    PRIMARY KEY (question, response)
);
