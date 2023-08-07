const fs = require("fs");

const files = [
  "CONVERSION FACTOR.csv",
  "FOOD NAME.csv",
  "FOOD GROUP.csv",
  "FOOD SOURCE.csv",
  "MEASURE NAME.csv",
  "NUTRIENT AMOUNT.csv",
  "NUTRIENT NAME.csv",
  "NUTRIENT SOURCE.csv",
  "REFUSE AMOUNT.csv",
  "REFUSE NAME.csv",
  "YIELD AMOUNT.csv",
  "YIELD NAME.csv",
];

for (let file of files) {
  const data = fs.readFileSync(file, { encoding: "latin1" }).toString();
  let rows = data.split("\n");

  let cols = rows[0]
    .split(",")
    .filter((val) => val !== "" && val !== "\r").length;

  rows = rows
    .map(sanitizeRow)
    .filter((row) => row.split(",").some((val) => !!val))
    .map((row) => row.split(",").slice(0, cols).join(",").replace(/"/g, '\\"'));

  fs.writeFileSync(file, rows.join("\n"), { encoding: "utf8" });
}

function sanitizeRow(str) {
  let newStr = "";
  let isInQuotes = false;
  let idx = 0;

  for (idx; idx < str.length; idx++) {
    const char = str[idx];
    if (char === `"`) {
      if (!isInQuotes) {
        isInQuotes = true;
      } else {
        isInQuotes = false;
      }
    }

    if (isInQuotes) {
      switch (char) {
        case ",":
          newStr += ";";
          break;
        default:
          newStr += char;
      }
    } else {
      switch (char) {
        case "\r":
          break;
        default:
          newStr += char;
      }
    }
  }

  return newStr;
}
