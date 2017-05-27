CREATE KEYSPACE chatbot
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE chatbot;

CREATE TABLE conversations (
    question text,
    response text,
    hits int,
    PRIMARY KEY (question, response)
);
