SET debug_print_parse = FALSE;
SET debug_print_rewritten = FALSE;
SET debug_print_plan = TRUE;

INSERT INTO users (id, age, name)
VALUES (1, 20, 'Bob'),
       (2, 22, 'Alice');

INSERT INTO users (SELECT u.id + 10, u.age, u.name, u.surname FROM users u);

DELETE
FROM users
WHERE 1 = 1

