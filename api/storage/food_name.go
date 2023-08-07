package storage

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

func AllFoodNames() ([]FoodName, error) {
	rows, err := db.Query(`
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
		FROM Food_Name;
	`)
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
