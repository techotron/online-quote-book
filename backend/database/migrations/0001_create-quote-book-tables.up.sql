CREATE TABLE IF NOT EXISTS quote_books(
    quote_book_id SERIAL PRIMARY KEY,
    quote_book_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS witnesses(
    witness_id SERIAL PRIMARY KEY,
    witness_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS quotees(
    quotee_id SERIAL PRIMARY KEY,
    quotee_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS quotes(
    quote_id SERIAL,
    quote_book_id INT,
    quote_text TEXT NOT NULL,
    quotee_id INT NOT NULL,
    witness_id INT NOT NULL,
    is_deleted BOOLEAN,
    quote_date timestamp,
    PRIMARY KEY(quote_id, quote_book_id),
    CONSTRAINT fk_quotee_on_quotes FOREIGN KEY(quotee_id) REFERENCES quotees(quotee_id),
    CONSTRAINT fk_witness_on_quotes FOREIGN KEY(witness_id) REFERENCES witnesses(witness_id),
    CONSTRAINT fk_quote_book_id_on_quotes FOREIGN KEY(quote_book_id) REFERENCES quote_books(quote_book_id)
);
