# HomegrownDB

## Used Sources
### Main DB sources
- base tutorial https://cstack.github.io/db_tutorial/
- base in deept explanation on how postgres works: https://www.interdb.jp/pg/
- how db works long article http://coding-geek.com/how-databases-work/
- SQLite docs https://www.sqlite.org/arch.html
- SQLite vde https://www.sqlite.org/vdbe.html
- article how postgresql works https://patshaughnessy.net/2014/10/13/following-a-select-statement-through-postgres-internals
- postgres docs:
  - internal docs https://www.postgresql.org/docs/14/overview.html
  - executor docs: https://www.postgresql.org/docs/current/executor.html
- postgres pro:
  - MVCC part1 https://postgrespro.com/blog/pgsql/5967856
  - locks part 1 https://postgrespro.com/blog/pgsql/5967999
  - queries in postgres
- buffer strategy https://github.com/postgres/postgres/blob/master/src/backend/storage/buffer/bufmgr.c#L719
- postgres reporisotrium parse nodes: https://doxygen.postgresql.org/primnodes_8h_source.html
- plan types explained: https://www.pgmustard.com/blog/2019/9/17/postgres-execution-plans-field-glossary
- postgres src:
  - nodes:
    - primitive nodes: https://doxygen.postgresql.org/primnodes_8h_source.html
    - parse nodes: https://doxygen.postgresql.org/parsenodes_8h_source.html
    - plan nodes: https://doxygen.postgresql.org/plannodes_8h_source.html
    - nodes: https://doxygen.postgresql.org/nodes_8h_source.html
  - free space map https://github.com/postgres/postgres/tree/master/src/backend/storage/freespace
  - buffer https://github.com/postgres/postgres/blob/master/src/backend/storage/buffer/README
  - executor tuple format: https://doxygen.postgresql.org/tuptable_8h_source.html

### other
- postgres query parser python util: https://github.com/lelit/pglast


### Additional resources sources
- hash map in java http://coding-geek.com/how-does-a-hashmap-work-in-java/
- parse tree https://en.wikipedia.org/wiki/Parse_tree
- b-tree https://www.programiz.com/dsa/b-tree
- b+tree https://www.programiz.com/dsa/b-plus-tree
- postgres data aligment https://www.enterprisedb.com/postgres-tutorials/data-alignment-postgresql
