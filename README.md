# simpl-coding-challenge
This repo is to show case simpl coding challenge

StartUp: 

- docker-compose up -d
- docker-compose exec simpl sh
- /bin/bash -c "go build -v -o ./bin ./cmd && ./cmd/cmd"

For any help on supported commands, type help and it shows all supported commands

For Database and table creation. 

- docker-compose exec mysqldb sh
- mysql -u root -prootroot
- use simpl
- create table merchant_details(merchant_name varchar(30), email_id varchar(40), merchant_discount double, CONSTRAINT unique_name UNIQUE (merchant_name));
- create table user_details(user_name varchar(50),email_id varchar(50),credit_limit int(11),due_amount int(11), CONSTRAINT unique_name UNIQUE (user_name));
- create table transactions(user_name varchar(50), merchant_name varchar(50), txn_value int(11), merchant_discount double, FOREIGN KEY (user_name) REFERENCES user_details(user_name),  FOREIGN KEY (merchant_name) REFERENCES merchant_details(merchant_name));
- create table user_payments(user_name varchar(50), amount_paid int(11), paid_on DATE, payment_id varchar(20), FOREIGN KEY (user_name) REFERENCES user_details(user_name));

