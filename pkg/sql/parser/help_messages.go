// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 977
	`ALTER`: {
		//line sql.y: 978
		Category: hGroup,
		//line sql.y: 979
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 992
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 993
		Category: hDDL,
		//line sql.y: 994
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1016
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1028
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1029
		Category: hDDL,
		//line sql.y: 1030
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1032
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1039
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1040
		Category: hPriv,
		//line sql.y: 1041
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1043
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1048
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1049
		Category: hDDL,
		//line sql.y: 1050
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1052
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1063
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1064
		Category: hDDL,
		//line sql.y: 1065
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1073
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1313
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1314
		Category: hCCL,
		//line sql.y: 1315
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1332
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1340
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1341
		Category: hCCL,
		//line sql.y: 1342
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1358
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1372
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1373
		Category: hCCL,
		//line sql.y: 1374
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   comma = '...'          [CSV-specific]
   comment = '...'        [CSV-specific]
   nullif = '...'         [CSV-specific]

`,
		//line sql.y: 1392
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1487
	`CANCEL`: {
		//line sql.y: 1488
		Category: hGroup,
		//line sql.y: 1489
		Text: `CANCEL JOB, CANCEL QUERY
`,
	},
	//line sql.y: 1495
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1496
		Category: hMisc,
		//line sql.y: 1497
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1498
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1506
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1507
		Category: hMisc,
		//line sql.y: 1508
		Text: `CANCEL QUERY <queryid>
`,
		//line sql.y: 1509
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1517
	`CREATE`: {
		//line sql.y: 1518
		Category: hGroup,
		//line sql.y: 1519
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW
`,
	},
	//line sql.y: 1537
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1538
		Category: hDML,
		//line sql.y: 1539
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1542
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1555
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1556
		Category: hCfg,
		//line sql.y: 1557
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1569
	`DROP`: {
		//line sql.y: 1570
		Category: hGroup,
		//line sql.y: 1571
		Text: `DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP USER
`,
	},
	//line sql.y: 1583
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1584
		Category: hDDL,
		//line sql.y: 1585
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1586
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1598
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1599
		Category: hDDL,
		//line sql.y: 1600
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1601
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 1613
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1614
		Category: hDDL,
		//line sql.y: 1615
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1616
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1636
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1637
		Category: hDDL,
		//line sql.y: 1638
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1639
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 1659
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1660
		Category: hPriv,
		//line sql.y: 1661
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1662
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1704
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1705
		Category: hMisc,
		//line sql.y: 1706
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1717
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 1778
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 1779
		Category: hMisc,
		//line sql.y: 1780
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 1781
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1803
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 1804
		Category: hMisc,
		//line sql.y: 1805
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 1806
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1829
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 1830
		Category: hMisc,
		//line sql.y: 1831
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 1832
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 1852
	`GRANT`: {
		ShortDescription: `define access privileges`,
		//line sql.y: 1853
		Category: hPriv,
		//line sql.y: 1854
		Text: `
GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1864
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 1872
	`REVOKE`: {
		ShortDescription: `remove access privileges`,
		//line sql.y: 1873
		Category: hPriv,
		//line sql.y: 1874
		Text: `
REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1884
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 1971
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 1972
		Category: hCfg,
		//line sql.y: 1973
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 1974
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 1986
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 1987
		Category: hCfg,
		//line sql.y: 1988
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 1989
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2019
	`SCRUB TABLE`: {
		ShortDescription: `run a scrub check on a table`,
		//line sql.y: 2020
		Category: hMisc,
		//line sql.y: 2021
		Text: `
SCRUB TABLE <tablename> [WITH <option> [, ...]]

Options:
  SCRUB TABLE ... WITH OPTIONS INDEX ALL
  SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  SCRUB TABLE ... WITH OPTIONS PHYSICAL

`,
	},
	//line sql.y: 2063
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2064
		Category: hCfg,
		//line sql.y: 2065
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2066
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2087
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2088
		Category: hCfg,
		//line sql.y: 2089
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 2094
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2111
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2112
		Category: hTxn,
		//line sql.y: 2113
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2120
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2259
	`SHOW`: {
		//line sql.y: 2260
		Category: hGroup,
		//line sql.y: 2261
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW USERS, SHOW TRANSACTION, SHOW BACKUP,
SHOW JOBS, SHOW QUERIES, SHOW SESSIONS, SHOW TRACE
`,
	},
	//line sql.y: 2287
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2288
		Category: hCfg,
		//line sql.y: 2289
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2290
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2311
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2312
		Category: hCCL,
		//line sql.y: 2313
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2314
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2322
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2323
		Category: hCfg,
		//line sql.y: 2324
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2327
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2344
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2345
		Category: hDDL,
		//line sql.y: 2346
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2347
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2355
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2356
		Category: hDDL,
		//line sql.y: 2357
		Text: `SHOW DATABASES
`,
		//line sql.y: 2358
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2366
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2367
		Category: hPriv,
		//line sql.y: 2368
		Text: `SHOW GRANTS [ON <targets...>] [FOR <users...>]
`,
		//line sql.y: 2369
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2377
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2378
		Category: hDDL,
		//line sql.y: 2379
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2380
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2398
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2399
		Category: hDDL,
		//line sql.y: 2400
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2401
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2414
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2415
		Category: hMisc,
		//line sql.y: 2416
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2417
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2433
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2434
		Category: hMisc,
		//line sql.y: 2435
		Text: `SHOW JOBS
`,
		//line sql.y: 2436
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2444
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2445
		Category: hMisc,
		//line sql.y: 2446
		Text: `
SHOW [KV] TRACE FOR SESSION
SHOW [KV] TRACE FOR <statement>
`,
		//line sql.y: 2449
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2470
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2471
		Category: hMisc,
		//line sql.y: 2472
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
	},
	//line sql.y: 2488
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2489
		Category: hDDL,
		//line sql.y: 2490
		Text: `SHOW TABLES [FROM <databasename>]
`,
		//line sql.y: 2491
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 2503
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2504
		Category: hCfg,
		//line sql.y: 2505
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2506
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 2525
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2526
		Category: hDDL,
		//line sql.y: 2527
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2528
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 2536
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2537
		Category: hDDL,
		//line sql.y: 2538
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2539
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 2547
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 2548
		Category: hPriv,
		//line sql.y: 2549
		Text: `SHOW USERS
`,
		//line sql.y: 2550
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 2623
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 2624
		Category: hMisc,
		//line sql.y: 2625
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 2626
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 2634
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 2635
		Category: hDDL,
		//line sql.y: 2636
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 2662
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 3142
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 3143
		Category: hDML,
		//line sql.y: 3144
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 3145
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 3153
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 3154
		Category: hPriv,
		//line sql.y: 3155
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 3156
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 3178
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 3179
		Category: hDDL,
		//line sql.y: 3180
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 3181
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 3195
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 3196
		Category: hDDL,
		//line sql.y: 3197
		Text: `
CREATE [UNIQUE] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3205
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 3355
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3356
		Category: hTxn,
		//line sql.y: 3357
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3358
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3366
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 3367
		Category: hMisc,
		//line sql.y: 3368
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 3369
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 3377
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 3378
		Category: hTxn,
		//line sql.y: 3379
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 3380
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3394
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 3395
		Category: hTxn,
		//line sql.y: 3396
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3404
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 3417
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 3418
		Category: hTxn,
		//line sql.y: 3419
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 3422
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 3435
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 3436
		Category: hTxn,
		//line sql.y: 3437
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 3438
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 3551
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 3552
		Category: hDDL,
		//line sql.y: 3553
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3554
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 3623
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 3624
		Category: hDML,
		//line sql.y: 3625
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 3630
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 3647
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 3648
		Category: hDML,
		//line sql.y: 3649
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 3653
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 3729
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 3730
		Category: hDML,
		//line sql.y: 3731
		Text: `UPDATE <tablename> [[AS] <name>] SET ... [WHERE <expr>] [RETURNING <exprs...>]
`,
		//line sql.y: 3732
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 3898
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 3899
		Category: hDML,
		//line sql.y: 3900
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 3911
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 3912
		Category: hDML,
		//line sql.y: 3913
		Text: `
SELECT [DISTINCT]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
       [ FOR UPDATE ]
`,
		//line sql.y: 3926
		SeeAlso: `WEBDOCS/select.html
`,
	},
	//line sql.y: 3986
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 3987
		Category: hDML,
		//line sql.y: 3988
		Text: `TABLE <tablename>
`,
		//line sql.y: 3989
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4252
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 4253
		Category: hDML,
		//line sql.y: 4254
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 4255
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4360
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 4361
		Category: hDML,
		//line sql.y: 4362
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 4380
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
