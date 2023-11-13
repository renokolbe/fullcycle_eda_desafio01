DROP DATABASE IF EXISTS wallet;
CREATE DATABASE wallet;  
USE wallet;  

DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    id varchar(255), 
    name varchar(255), 
    email varchar(255), 
    created_at date
);

INSERT INTO clients (id, name, email) VALUES ('C1', 'John Doe', 'john.doe@email.com');
INSERT INTO clients (id, name, email) VALUES ('C2', 'Jane Doe', 'jane.doe@email.com');

DROP TABLE IF EXISTS accounts;   
CREATE TABLE accounts (
    id varchar(255), 
    client_id varchar(255), 
    balance decimal(10,2) DEFAULT 1000, 
    created_at date
);

INSERT INTO accounts (id, client_id, balance) VALUES ('A1_1', 'C1', 900.00);
INSERT INTO accounts (id, client_id, balance) VALUES ('A1_2', 'C1', 1100.00);
INSERT INTO accounts (id, client_id, balance) VALUES ('A2_1', 'C2', 1000.00);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id varchar(255), 
    account_id_from varchar(255), 
    account_id_to varchar(255), 
    amount decimal(10,2), 
    created_at date
);

INSERT INTO transactions (id, account_id_from, account_id_to, amount) VALUES ('T1', 'A1_1', 'A1_2', 100.00);

DROP TABLE IF EXISTS account_balance;
CREATE TABLE account_balance (
    id varchar(255), 
    balance decimal(10,2)
);

INSERT INTO account_balance (id, balance) VALUES ('A1_1', 900.00);
INSERT INTO account_balance (id, balance) VALUES ('A1_2', 1100.00);
