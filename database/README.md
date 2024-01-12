# `database/`

Where the db file is. And everything else used to make it.

# Source

Data is from the [Canadian Nutrient File (CNF), 2015](https://www.canada.ca/en/health-canada/services/food-nutrition/healthy-eating/nutrient-data/canadian-nutrient-file-2015-download-files.html) from [Health Canada](https://www.canada.ca/en/health-canada.html).

# Set up the SQLite DB

## We got some data!

But unfortunately the csv files we downloaded aren't in the best shape. Some files have thousands of empty rows, some data rows are meant to be quoted text (I guess?) but at least one was missing an opening `"`. Furthermore, some files have extra columns. All of these things will cause errors during SQLite import. To further complicate things, we can't simply import the CSVs into SQLite: All colummns will get automatically created with the TEXT datatype, and there isn't a way we can alter that after the fact. Also, it turns out, the files are in `latin-1` encoding and we need `utf-8`.

## Processing the data

So, before we can create our DB and import the data, we need to sanitize it a bit and better format it for importing into sqlite. `clean_the_csvs.js` will read the CSV files we got and try to rectify the issues outlined above, overwriting the old CSV files with the improved data.

```bash
node clean_the_csvs.js
```

If you want to make adjustments to that script and run it again, you'll either need to start with the original `latin-1` files, or remember to expect `utf-8` encoding instead of `latin-1` when reading them in again.

## Create the DB

Now that the data is in better shape, we can create the DB, tables, and records. To ensure we have better types than just `TEXT`, we need to create the tables with SQL. I don't want to write all that SQL, so let's have sqlite do most of the work for us. There's still a bit of manual work, but we should only have to run through it the one time.

### 1. Generate most of the necessary `CREATE TABLE` SQL

Run `init_pre.sh`. **This will run the JS script to clean up the CSV files** and then run a SQLite script that does a few things:

- load the CSVs into SQLite (resulting in the bad datatypes).
- read out the `CREATE TABLE` SQL for each of the resulting tables into an output file, `csv_create.sql`

Comment out the `node` command if you've cleaned up the CSVs already.

### 2. Make better `CEATE TABLE` SQL from the output

Manually edit `csv_create.sql` so that the datatypes are correct (consult the corresponding CSV files) and add some `PRIMARY KEY` and `FOREIGN KEY` directives. Save this file as `csv_create_updated.sql`

### 3. Create the DB and load data for _real_

First, just delete the original `food.db` file that got created before. Then run `init_db.sh`. This will run our updated `CREATE TABLE` statements and then import the CSV files, skipping the header rows.

### 4. Done!

Now to get back to doing what you _actually_ sat down to do!
