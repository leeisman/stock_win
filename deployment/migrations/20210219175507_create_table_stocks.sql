-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS stock_win.stocks(
    id int NOT NULL AUTO_INCREMENT,
    symbol_id varchar(1024),
    symbol_name varchar(1024),
    price FLOAT NOT NULL DEFAULT '0.00',
	PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
