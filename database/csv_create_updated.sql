CREATE TABLE IF NOT EXISTS "Food_Group" (
  "FoodGroupID" INT PRIMARY KEY,
  "FoodGroupCode" INT,
  "FoodGroupName" TEXT,
  "FoodGroupNameF" TEXT
);

CREATE TABLE IF NOT EXISTS "Food_Source" (
  "FoodSourceID" INT PRIMARY KEY,
  "FoodSourceCode" INT,
  "FoodSourceDescription" TEXT,
  "FoodSourceDescriptionF" TEXT
);

CREATE TABLE IF NOT EXISTS "Food_Name" (
  "FoodID" INT PRIMARY KEY,
  "FoodCode" INT,
  "FoodGroupID" INT,
  "FoodSourceID" INT,
  "FoodDescription" TEXT,
  "FoodDescriptionF" TEXT,
  "FoodDateOfEntry" TEXT,
  "FoodDateOfPublication" TEXT,
  "CountryCode" TEXT,
  "ScientificName" TEXT,
  FOREIGN KEY (FoodGroupID) REFERENCES Food_Group(FoodGroupID), 
  FOREIGN KEY (FoodSourceID) REFERENCES Food_Source(FoodSourceID) 
);

CREATE TABLE IF NOT EXISTS Yield_Name (
  "YieldID" INT PRIMARY KEY,
  "YieldDescription" TEXT,
  "YieldDescriptionF" TEXT
);

CREATE TABLE IF NOT EXISTS Yield_Amount (
  "FoodID" INT,
  "YieldID" INT,
  "YieldAmount" INT,
  "YieldDateofEntry" TEXT,
  PRIMARY KEY (FoodID, YieldID),
  FOREIGN KEY (FoodID) REFERENCES Food_Name(FoodID),
  FOREIGN KEY (YieldID) REFERENCES Yield_Name(YieldID)
);

CREATE TABLE IF NOT EXISTS Refuse_Name (
  "RefuseID" INT PRIMARY KEY,
  "RefuseDescription" TEXT,
  "RefuseDescriptionF" TEXT
);

CREATE TABLE IF NOT EXISTS Refuse_Amount (
  "FoodID" INT,
  "RefuseID" INT,
  "RefuseAmount" INT,
  "RefuseDateOfEntry" TEXT,
  PRIMARY KEY (FoodID, RefuseID),
  FOREIGN KEY (FoodID) REFERENCES Food_Name(FoodID),
  FOREIGN KEY (RefuseID) REFERENCES Refuse_Name(RefuseID)
);

CREATE TABLE IF NOT EXISTS Nutrient_Source (
  "NutrientSourceID" INT PRIMARY KEY,
  "NutrientSourceCode" INT,
  "NutrientSourceDescription" TEXT,
  "NutrientSourceDescriptionF" TEXT
);

CREATE TABLE IF NOT EXISTS Nutrient_Name (
  "NutrientID" INT PRIMARY KEY,
  "NutrientCode" INT,
  "NutrientSymbol" TEXT,
  "NutrientUnit" TEXT,
  "NutrientName" TEXT,
  "NutrientNameF" TEXT,
  "Tagname" TEXT,
  "NutrientDecimals" INT 
);

CREATE TABLE IF NOT EXISTS Nutrient_Amount (
  "FoodID" INT,
  "NutrientID" INT,
  "NutrientValue" REAL,
  "StandardError" INT,
  "NumberofObservations" INT,
  "NutrientSourceID" INT,
  "NutrientDateOfEntry" TEXT,
  PRIMARY KEY (FoodID, NutrientID),
  FOREIGN KEY (FoodID) REFERENCES Food_Name(FoodID),
  FOREIGN KEY (NutrientID) REFERENCES Nutrient_Name(NutrientID)
);

CREATE TABLE IF NOT EXISTS "Measure_Name" (
  "MeasureID" INT PRIMARY KEY,
  "MeasureDescription" TEXT,
  "MeasureDescriptionF" TEXT
);

CREATE TABLE IF NOT EXISTS "Conversion_Factor" (
  "FoodID" INT,
  "MeasureID" INT,
  "ConversionFactorValue" REAL,
  "ConvFactorDateOfEntry" TEXT,
  PRIMARY KEY (FoodID, MeasureID),
  FOREIGN KEY (FoodID) REFERENCES Food_Name(FoodID),
  FOREIGN KEY (MeasureID) REFERENCES Measure_Name(MeasureID)
);
