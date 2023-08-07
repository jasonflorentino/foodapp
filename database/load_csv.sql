.mode csv

.import "CONVERSION FACTOR.csv" Conversion_Factor
.import "FOOD GROUP.csv" Food_Group
.import "FOOD NAME.csv" Food_Name
.import "FOOD SOURCE.csv" Food_Source
.import "MEASURE NAME.csv" Measure_Name
.import "NUTRIENT AMOUNT.csv" Nutrient_Amount
.import "NUTRIENT NAME.csv" Nutrient_Name
.import "NUTRIENT SOURCE.csv" Nutrient_Source
.import "REFUSE AMOUNT.csv" Refuse_Amount
.import "REFUSE NAME.csv" Refuse_Name
.import "YIELD AMOUNT.csv" Yield_Amount
.import "YIELD NAME.csv" Yield_Name

.output "csv_create.sql"

.schema Yield_Name
.schema Yield_Amount
.schema Refuse_Name
.schema Refuse_Amount
.schema Nutrient_Source
.schema Nutrient_Name
.schema Nutrient_Amount
.schema Measure_Name
.schema Food_Source
.schema Food_Name
.schema Food_Group
.schema Conversion_Factor

.output
.quit
