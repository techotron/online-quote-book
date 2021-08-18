INSERT INTO quote_books(quote_book_collection, quote_book_name, created_on, last_updated) 
VALUES ('public', 'testquotebook', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO quotees(quotee_name, quote_book_collection, quote_book_name) VALUES ('test quotee', 'public', 'testquotebook');

INSERT INTO witnesses(witness_name, quote_book_collection, quote_book_name) VALUES ('test witness', 'public', 'testquotebook');

INSERT INTO quotes(quote_book_collection, quote_book_name, quote_text, quotee_name, witness_name, is_deleted, quote_date, inserted_date) VALUES 
    ('public', 'testquotebook', 'Balls are round, like windows', 'test quotee', 'test witness', FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    ('public', 'testquotebook', 'Once I trapped my elbow in a keyhole', 'test quotee', 'test witness', FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    ('public', 'testquotebook', 'Mice make the best grass feed', 'test quotee', 'test witness', FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    ('public', 'testquotebook', 'Bricks are often found relaxing in Kent during the summer', 'test quotee', 'test witness', FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    ('public', 'testquotebook', 'Trees get loney in Autumn when the leaves leave', 'test quotee', 'test witness', FALSE, '2015-09-10', CURRENT_TIMESTAMP);
