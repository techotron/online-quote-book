INSERT INTO quote_books(quote_book_name) VALUES ('testquotebook');

INSERT INTO quotees(quotee_name) VALUES ('test quotee');

INSERT INTO witnesses(witness_name) VALUES ('test witness');

INSERT INTO quotes(quote_book_id, quote_text, quotee_id, witness_id, is_deleted, quote_date, inserted_date) VALUES 
    (1, 'Balls are round, like windows', 1, 1, FALSE, '2015-09-10', current_timestamp),
    (1, 'Once I trapped my elbow in a keyhole', 1, 1, FALSE, '2015-09-10', current_timestamp),
    (1, 'Mice make the best grass feed', 1, 1, FALSE, '2015-09-10', current_timestamp),
    (1, 'Bricks are often found relaxing in Kent during the summer', 1, 1, FALSE, '2015-09-10', current_timestamp),
    (1, 'Trees get loney in Autumn when the leaves leave', 1, 1, FALSE, '2015-09-10', current_timestamp);
