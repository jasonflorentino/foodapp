package storage

import (
	"fmt"
	"net/url"
	"strconv"
)

type FoodName struct {
	FoodId            string     `json:"id"`
	Code              NullInt64  `json:"code"`
	FoodGroupId       NullInt64  `json:"food_group_id"`
	FoodSourceId      NullInt64  `json:"food_source_id"`
	Description       NullString `json:"description"`
	DescriptionF      NullString `json:"description_f"`
	DateOfEntry       NullString `json:"date_of_entry"`
	DateOfPublication NullString `json:"date_of_publication"`
	CountryCode       NullString `json:"country_code"`
	ScientificName    NullString `json:"scientific_name"`
}

func AllFoodNames(params url.Values) ([]FoodName, error) {
	limitSql := "LIMIT 100"
	offsetSql := "OFFSET 0"

	if limit, err := strconv.Atoi(params.Get("limit")); err != nil {
		// Dont care, could be empty
	} else {
		if limit > 0 && limit <= 200 {
			limitSql = fmt.Sprintf("LIMIT %d", limit)
		}
	}

	if offset, err := strconv.Atoi(params.Get("offset")); err != nil {
		// Dont care, could be empty
	} else {
		if offset > 0 {
			offsetSql = fmt.Sprintf("OFFSET %d", offset)
		}
	}

	rows, err := db.Query(fmt.Sprintf(`
		SELECT 
			FoodId, 
			FoodCode, 
			FoodGroupId, 
			FoodSourceId, 
			FoodDescription, 
			FoodDescriptionF, 
			FoodDateOfEntry, 
			FoodDateOfPublication, 
			CountryCode, 
			ScientificName  
		FROM Food_Name %s %s;
	`, limitSql, offsetSql))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var foodNames []FoodName
	for rows.Next() {
		var fn FoodName
		err := rows.Scan(
			&fn.FoodId,
			&fn.Code,
			&fn.FoodGroupId,
			&fn.FoodSourceId,
			&fn.Description,
			&fn.DescriptionF,
			&fn.DateOfEntry,
			&fn.DateOfPublication,
			&fn.CountryCode,
			&fn.ScientificName,
		)
		if err != nil {
			return nil, err
		}

		foodNames = append(foodNames, fn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foodNames, nil
}
