CREATE TABLE birds (
                       id SERIAL8 PRIMARY KEY ,
                       specie varchar(255),
                       uuid uuid
);

INSERT INTO birds
VALUES
    (3, ('we2'), (gen_random_uuid()));

INSERT INTO birds(specie, uuid)
SELECT specie, uuid FROM birds;

EXPLAIN SELECT *
        FROM birds b1
                 JOIN birds b2 on b1.uuid = b2.uuid
                 JOIN birds b3 on b2.uuid = b3.uuid
        WHERE b1.specie = '';

SELECT gen_random_uuid();

EXPLAIN (VERBOSE, costs off, format json ) INSERT INTO birds(specie, uuid) VALUES ('we', gen_random_uuid()), ('we2', uuid_in('4cb7974c-424c-4d5e-9394-727c66769dc2'));

EXPLAIN ANALYSE VERBOSE INSERT INTO birds(specie, uuid) VALUES ('we', gen_random_uuid());

EXPLAIN VERBOSE UPDATE birds SET specie = '22' WHERE specie = '23';


EXPLAIN ANALYSE VERBOSE INSERT INTO birds(specie, uuid) (SELECT specie, uuid FROM birds);

EXPLAIN (VERBOSE, costs off, format json )
SELECT b1.specie, (SELECT count(b2.id) FROM birds b2)
FROM birds b1
         JOIN birds b2 ON (b1.uuid = b2.uuid)