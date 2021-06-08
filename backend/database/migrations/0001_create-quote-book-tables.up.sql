CREATE TABLE IF NOT EXISTS quote_book(
    quote_book_id SERIAL PRIMARY KEY,
    quote_book_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS witnesses(
    witness_id SERIAL PRIMARY KEY,
    witness_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS quotes(
    quote_id SERIAL,
    quote_book_id INT,
    quote_text TEXT NOT NULL,
    quotee VARCHAR(50) NOT NULL,
    witness_id INT NOT NULL,
    is_deleted BOOLEAN,
    quote_date timestamp,
    PRIMARY KEY(quote_id, quote_book_id),
    CONSTRAINT fk_witness_on_quotes FOREIGN KEY(witness_id) REFERENCES witnesses(witness_id),
    CONSTRAINT fk_quote_book_id_on_quotes FOREIGN KEY(quote_book_id) REFERENCES quote_book(quote_book_id)
);
