INSERT INTO quote_books(quote_book_name, created_on, last_updated) VALUES ('testquotebook', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO quotees(quotee_name, quote_book_id) VALUES ('test quotee', 1);

INSERT INTO witnesses(witness_name, quote_book_id) VALUES ('test witness', 1);

INSERT INTO quotes(quote_book_id, quote_text, quotee_id, witness_id, is_deleted, quote_date, inserted_date) VALUES 
    (1, 'Balls are round, like windows', 1, 1, FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    (1, 'Once I trapped my elbow in a keyhole', 1, 1, FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    (1, 'Mice make the best grass feed', 1, 1, FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    (1, 'Bricks are often found relaxing in Kent during the summer', 1, 1, FALSE, '2015-09-10', CURRENT_TIMESTAMP),
    (1, 'Trees get loney in Autumn when the leaves leave', 1, 1, FALSE, '2015-09-10', CURRENT_TIMESTAMP);
