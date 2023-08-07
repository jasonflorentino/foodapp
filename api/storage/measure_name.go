package storage

type MeasureName struct {
	Id           string `json:"id"`
	Description  string `json:"description"`
	DescriptionF string `json:"description_f"`
}

func AllMeasureNames() ([]MeasureName, error) {
	rows, err := db.Query(`
		SELECT
			Id,
			Description,
			DescriptionF
		FROM Measure_Name
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var measureNames []MeasureName
	for rows.Next() {
		var mn MeasureName
		err := rows.Scan(
			&mn.Id,
			&mn.Description,
			&mn.DescriptionF,
		)
		if err != nil {
			return nil, err
		}

		measureNames = append(measureNames, mn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return measureNames, nil
}
