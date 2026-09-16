# Important Information to use the golang-migration tool

# Commands
## Executing Migrations
```bash
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```

## Check Migration Version 
```bash
migrate -path=./migrations -database=$DSN version
```

## Migrate to Specific Version
```bash
#migrate to version 1
migrate -path=./migrations -database=$DSN goto 1
```

## Executing Down Migrations
```bash
migrate -path=./migrations -database=$DSN down 1
```

## Executing All Down Migrations
```bash
migrate -path=./migrations -database=$DSN down 
```

# Tips

## How to Handle Migration Failures
During the execution of a migration, all SQL statements leading up to an error execute and will be applied to the DB. Therefore, a failure may leave the databse in an unknown state -- as far as the migration tool is concerned.

Accordingly, the `version` field in the `schema_migrations` table will contain the number for the "failed migration" and the `dirty` field will be set to `true`

At this point, even if you run a `down` migration, you will get the following error:
```bash
Dirty database version {X}. Fix and force version.
```

**To fix** : investigate the original error and figure out if the migration file which failed, was partially applied. If so, then need to manually roll back the partially applied migration

Then must force the `version` number in the `schema_migrations` table to the correct value:

```bash
migrate -path=./migrations -database=$DSN force 1
````
