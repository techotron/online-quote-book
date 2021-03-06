CREATE TABLE IF NOT EXISTS quote_books(
    quote_book_id SERIAL,
    quote_book_collection TEXT NOT NULL,
    quote_book_name TEXT NOT NULL,
    created_on TIMESTAMP,
    last_updated TIMESTAMP,
    PRIMARY KEY (quote_book_collection, quote_book_name)
);

CREATE TABLE IF NOT EXISTS witnesses(
    witness_id SERIAL,
    quote_book_collection TEXT NOT NULL,
    quote_book_name TEXT NOT NULL,
    witness_name VARCHAR(50) NOT NULL,
    PRIMARY KEY(quote_book_collection, quote_book_name, witness_name),
    CONSTRAINT fk_quotebook_on_witnesses FOREIGN KEY(quote_book_collection, quote_book_name) REFERENCES quote_books(quote_book_collection, quote_book_name)
);

CREATE TABLE IF NOT EXISTS quotees(
    quotee_id SERIAL,
    quote_book_collection TEXT NOT NULL,
    quote_book_name TEXT NOT NULL,
    quotee_name VARCHAR(50) NOT NULL,
    PRIMARY KEY(quote_book_collection, quote_book_name, quotee_name),
    CONSTRAINT fk_quotebook_on_quotees FOREIGN KEY(quote_book_collection, quote_book_name) REFERENCES quote_books(quote_book_collection, quote_book_name)
);

CREATE TABLE IF NOT EXISTS quotes(
    quote_id SERIAL,
    quote_book_collection TEXT NOT NULL,
    quote_book_name TEXT NOT NULL,
    quote_text TEXT NOT NULL,
    quotee_name VARCHAR(50) NOT NULL,
    witness_name VARCHAR(50) NOT NULL,
    is_deleted BOOLEAN,
    inserted_date TIMESTAMP,
    quote_date VARCHAR,
    PRIMARY KEY(quote_id, quote_book_collection, quote_book_name),
    CONSTRAINT fk_quotee_on_quotes FOREIGN KEY(quote_book_collection, quote_book_name, quotee_name) REFERENCES quotees(quote_book_collection, quote_book_name, quotee_name),
    CONSTRAINT fk_witness_on_quotes FOREIGN KEY(quote_book_collection, quote_book_name, witness_name) REFERENCES witnesses(quote_book_collection, quote_book_name, witness_name),
    CONSTRAINT fk_quote_book_id_on_quotes FOREIGN KEY(quote_book_collection, quote_book_name) REFERENCES quote_books(quote_book_collection, quote_book_name)
);
