// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 973
	`ALTER`: {
		//line sql.y: 974
		Category: hGroup,
		//line sql.y: 975
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 988
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 989
		Category: hDDL,
		//line sql.y: 990
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
		//line sql.y: 1012
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1024
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1025
		Category: hDDL,
		//line sql.y: 1026
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1028
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1035
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1036
		Category: hPriv,
		//line sql.y: 1037
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1039
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1044
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1045
		Category: hDDL,
		//line sql.y: 1046
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1048
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1059
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1060
		Category: hDDL,
		//line sql.y: 1061
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1069
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1309
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1310
		Category: hCCL,
		//line sql.y: 1311
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
		//line sql.y: 1328
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1336
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1337
		Category: hCCL,
		//line sql.y: 1338
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
		//line sql.y: 1354
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1368
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1369
		Category: hCCL,
		//line sql.y: 1370
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
		//line sql.y: 1388
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1483
	`CANCEL`: {
		//line sql.y: 1484
		Category: hGroup,
		//line sql.y: 1485
		Text: `CANCEL JOB, CANCEL QUERY
`,
	},
	//line sql.y: 1491
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1492
		Category: hMisc,
		//line sql.y: 1493
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1494
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1502
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1503
		Category: hMisc,
		//line sql.y: 1504
		Text: `CANCEL QUERY <queryid>
`,
		//line sql.y: 1505
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1513
	`CREATE`: {
		//line sql.y: 1514
		Category: hGroup,
		//line sql.y: 1515
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW
`,
	},
	//line sql.y: 1533
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1534
		Category: hDML,
		//line sql.y: 1535
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1538
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1551
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1552
		Category: hCfg,
		//line sql.y: 1553
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1565
	`DROP`: {
		//line sql.y: 1566
		Category: hGroup,
		//line sql.y: 1567
		Text: `DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP USER
`,
	},
	//line sql.y: 1579
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1580
		Category: hDDL,
		//line sql.y: 1581
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1582
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1594
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1595
		Category: hDDL,
		//line sql.y: 1596
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1597
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 1609
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1610
		Category: hDDL,
		//line sql.y: 1611
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1612
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1632
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1633
		Category: hDDL,
		//line sql.y: 1634
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1635
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 1655
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1656
		Category: hPriv,
		//line sql.y: 1657
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1658
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1700
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1701
		Category: hMisc,
		//line sql.y: 1702
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1713
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 1774
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 1775
		Category: hMisc,
		//line sql.y: 1776
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 1777
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1799
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 1800
		Category: hMisc,
		//line sql.y: 1801
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 1802
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1825
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 1826
		Category: hMisc,
		//line sql.y: 1827
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 1828
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 1848
	`GRANT`: {
		ShortDescription: `define access privileges`,
		//line sql.y: 1849
		Category: hPriv,
		//line sql.y: 1850
		Text: `
GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1860
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 1868
	`REVOKE`: {
		ShortDescription: `remove access privileges`,
		//line sql.y: 1869
		Category: hPriv,
		//line sql.y: 1870
		Text: `
REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1880
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 1967
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 1968
		Category: hCfg,
		//line sql.y: 1969
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 1970
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 1982
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 1983
		Category: hCfg,
		//line sql.y: 1984
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 1985
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2015
	`SCRUB TABLE`: {
		ShortDescription: `run a scrub check on a table`,
		//line sql.y: 2016
		Category: hMisc,
		//line sql.y: 2017
		Text: `
SCRUB TABLE <tablename> [WITH <option> [, ...]]

Options:
  SCRUB TABLE ... WITH OPTIONS INDEX ALL
  SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)

`,
	},
	//line sql.y: 2054
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2055
		Category: hCfg,
		//line sql.y: 2056
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2057
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2078
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2079
		Category: hCfg,
		//line sql.y: 2080
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 2085
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2102
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2103
		Category: hTxn,
		//line sql.y: 2104
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2111
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2250
	`SHOW`: {
		//line sql.y: 2251
		Category: hGroup,
		//line sql.y: 2252
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW USERS, SHOW TRANSACTION, SHOW BACKUP,
SHOW JOBS, SHOW QUERIES, SHOW SESSIONS, SHOW TRACE
`,
	},
	//line sql.y: 2278
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2279
		Category: hCfg,
		//line sql.y: 2280
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2281
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2302
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2303
		Category: hCCL,
		//line sql.y: 2304
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2305
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2313
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2314
		Category: hCfg,
		//line sql.y: 2315
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2318
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2335
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2336
		Category: hDDL,
		//line sql.y: 2337
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2338
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2346
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2347
		Category: hDDL,
		//line sql.y: 2348
		Text: `SHOW DATABASES
`,
		//line sql.y: 2349
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2357
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2358
		Category: hPriv,
		//line sql.y: 2359
		Text: `SHOW GRANTS [ON <targets...>] [FOR <users...>]
`,
		//line sql.y: 2360
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2368
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2369
		Category: hDDL,
		//line sql.y: 2370
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2371
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2389
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2390
		Category: hDDL,
		//line sql.y: 2391
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2392
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2405
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2406
		Category: hMisc,
		//line sql.y: 2407
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2408
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2424
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2425
		Category: hMisc,
		//line sql.y: 2426
		Text: `SHOW JOBS
`,
		//line sql.y: 2427
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2435
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2436
		Category: hMisc,
		//line sql.y: 2437
		Text: `
SHOW [KV] TRACE FOR SESSION
SHOW [KV] TRACE FOR <statement>
`,
		//line sql.y: 2440
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2461
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2462
		Category: hMisc,
		//line sql.y: 2463
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
	},
	//line sql.y: 2479
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2480
		Category: hDDL,
		//line sql.y: 2481
		Text: `SHOW TABLES [FROM <databasename>]
`,
		//line sql.y: 2482
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 2494
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2495
		Category: hCfg,
		//line sql.y: 2496
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2497
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 2516
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2517
		Category: hDDL,
		//line sql.y: 2518
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2519
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 2527
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2528
		Category: hDDL,
		//line sql.y: 2529
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2530
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 2538
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 2539
		Category: hPriv,
		//line sql.y: 2540
		Text: `SHOW USERS
`,
		//line sql.y: 2541
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 2614
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 2615
		Category: hMisc,
		//line sql.y: 2616
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 2617
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 2625
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 2626
		Category: hDDL,
		//line sql.y: 2627
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
		//line sql.y: 2653
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 3133
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 3134
		Category: hDML,
		//line sql.y: 3135
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 3136
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 3144
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 3145
		Category: hPriv,
		//line sql.y: 3146
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 3147
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 3169
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 3170
		Category: hDDL,
		//line sql.y: 3171
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 3172
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 3186
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 3187
		Category: hDDL,
		//line sql.y: 3188
		Text: `
CREATE [UNIQUE] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3196
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 3346
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3347
		Category: hTxn,
		//line sql.y: 3348
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3349
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3357
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 3358
		Category: hMisc,
		//line sql.y: 3359
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 3360
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 3368
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 3369
		Category: hTxn,
		//line sql.y: 3370
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 3371
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3385
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 3386
		Category: hTxn,
		//line sql.y: 3387
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3395
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 3408
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 3409
		Category: hTxn,
		//line sql.y: 3410
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 3413
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 3426
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 3427
		Category: hTxn,
		//line sql.y: 3428
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 3429
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 3542
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 3543
		Category: hDDL,
		//line sql.y: 3544
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3545
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 3614
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 3615
		Category: hDML,
		//line sql.y: 3616
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 3621
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 3638
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 3639
		Category: hDML,
		//line sql.y: 3640
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 3644
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 3720
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 3721
		Category: hDML,
		//line sql.y: 3722
		Text: `UPDATE <tablename> [[AS] <name>] SET ... [WHERE <expr>] [RETURNING <exprs...>]
`,
		//line sql.y: 3723
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 3889
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 3890
		Category: hDML,
		//line sql.y: 3891
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 3902
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 3903
		Category: hDML,
		//line sql.y: 3904
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
		//line sql.y: 3917
		SeeAlso: `WEBDOCS/select.html
`,
	},
	//line sql.y: 3977
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 3978
		Category: hDML,
		//line sql.y: 3979
		Text: `TABLE <tablename>
`,
		//line sql.y: 3980
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4243
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 4244
		Category: hDML,
		//line sql.y: 4245
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 4246
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4351
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 4352
		Category: hDML,
		//line sql.y: 4353
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
		//line sql.y: 4371
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}