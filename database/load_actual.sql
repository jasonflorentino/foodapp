.read csv_create_updated.sql 

.mode csv

.import --skip 1 "CONVERSION FACTOR.csv" Conversion_Factor
.import --skip 1 "FOOD GROUP.csv" Food_Group
.import --skip 1 "FOOD NAME.csv" Food_Name
.import --skip 1 "FOOD SOURCE.csv" Food_Source
.import --skip 1 "MEASURE NAME.csv" Measure_Name
.import --skip 1 "NUTRIENT AMOUNT.csv" Nutrient_Amount
.import --skip 1 "NUTRIENT NAME.csv" Nutrient_Name
.import --skip 1 "NUTRIENT SOURCE.csv" Nutrient_Source
.import --skip 1 "REFUSE AMOUNT.csv" Refuse_Amount
.import --skip 1 "REFUSE NAME.csv" Refuse_Name
.import --skip 1 "YIELD AMOUNT.csv" Yield_Amount
.import --skip 1 "YIELD NAME.csv" Yield_Name

