# Europa
Simple Database CLI Migration Tools.

# Support Databases
Only

- MySQL

Reflects the developer's view of religion.

# How To Install

```
$ go get github.com/lrf141/europa
```

# How To Use

## Create Migration

Run the following command. Create Raw SQL Files.

```
$ europa create:migrate -n "migration name"

ex)
$ europa create:migrate -n "sample_migrate"
Create Migrate: ./migrations/migrate/20191028120501_sample_migrate.up.sql [Success]
Create Migrate: ./migrations/migrate/20191028120501_sample_migrate.down.sql [Success]
```

## Run Migration

```
$ europa run:migrate
```

## Rollback Migration
```
$ europa rollback:migrate
```

# Migrate Internal

Manage Migrate 
```
> desc migrate_schema;
+---------+------------+------+-----+---------+-------+
| Field   | Type       | Null | Key | Default | Extra |
+---------+------------+------+-----+---------+-------+
| migrate | text       | YES  |     | NULL    |       |
| flag    | tinyint(1) | YES  |     | 0       |       |
+---------+------------+------+-----+---------+-------+

ex)
> select * from migrate_schema;
+----------------------------------+------+
| migrate                          | flag |
+----------------------------------+------+
| 20191027021931_test_table_create |    0 |
| 20191028120501_sample_migrate    |    0 |
+----------------------------------+------+
2 rows in set (0.00 sec)

```

migrate : Migration Name  
flag = 1 : Migration Completed  
flag = 0 : Migration UnCompleted