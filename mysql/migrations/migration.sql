
-- create table
CREATE TABLE erajaya.products (
	id BIGINT NOT NULL auto_increment,
	name varchar(255) NOT NULL,
	price FLOAT NOT NULL,
	description varchar(255) NOT NULL,
	quantity INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
	PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;