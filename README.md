[![Golang-Version](https://img.shields.io/badge/Golang-1.12-brightgreen)](Golang-V)
[![GoModules](https://img.shields.io/badge/GoModules-enable-brightgreen)](GoModules)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
![](./logo.png)

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

## Make Configure
```
// require
$ mkdir migrations

$ vim migrations/db.yaml
```

example yaml file.
```yaml
database: sample
host: 127.0.0.1
port: 33068
user: sample
pass: sample
```

## Create Migration

Run the following command. Create Raw SQL Files.

```
$ europa create:migrate --name "migration name"

ex)
$ europa create:migrate -n "sample_migrate"
Create Migrate: ./migrations/migrate/20191028120501_sample_migrate.up.sql [Success]
Create Migrate: ./migrations/migrate/20191028120501_sample_migrate.down.sql [Success]
```

## Run Migration

```
$ europa run:migrate
or 
$ europa run:migrate --name 20191028120501_sample_migrate
```

## Rollback Migration
```
$ europa rollback:migrate
or
$ europa rollback:migrate --name 20191028120501_sample_migrate
```

## Create Seed
```
$ europa create:seed --name "seed name"

ex)
$ europa create:seed -n "sample_seed"
Create Migrate: ./migrations/seed/20191028120501_sample_seed.up.sql [Success]
Create Migrate: ./migrations/seed/20191028120501_sample_seed.down.sql [Success]
```

## Run Seed

```
$ europa run:seed
or 
$ europa run:seed --name 20191028120501_sample_seed
```

## Rollback Seed
```
$ europa rollback:seed
or
$ europa rollback:seed --name 20191028120501_sample_seed
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

# Seed Internal

```
> desc seed_schema;
+-------+------------+------+-----+---------+-------+
| Field | Type       | Null | Key | Default | Extra |
+-------+------------+------+-----+---------+-------+
| seed  | text       | YES  |     | NULL    |       |
| flag  | tinyint(1) | YES  |     | 0       |       |
+-------+------------+------+-----+---------+-------+
2 rows in set (0.00 sec)

```

seed : Seed Name  
flag = 1 : Seed Completed  
flag = 0 : Seed UnCompleted

# License

MIT License