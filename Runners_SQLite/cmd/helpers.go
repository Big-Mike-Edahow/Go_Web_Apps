/* helpers.go */

package main

import (
	"log"
)

func getOneRunner(id int) (Runner, error) {
	row, err := db.Query("SELECT id, name, age, country, season_best, personal_best  FROM runners WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	var runner Runner
	for row.Next() {
		err = row.Scan(&runner.Id, &runner.Name, &runner.Age, &runner.Country, &runner.SeasonBest, &runner.PersonalBest)
		if err != nil {
			log.Println(err)
		}
	}
	return runner, err
}

func getAllRunners() ([]Runner, error) {
	rows, err := db.Query("SELECT * FROM runners")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var runners []Runner
	for rows.Next() {
		var runner Runner
		
		err := rows.Scan(&runner.Id, &runner.Name, &runner.Age, &runner.Country, &runner.SeasonBest, &runner.PersonalBest)
		if err != nil {
			log.Println(err)
		}
		runners = append(runners, runner)
	}
	return runners, nil
}
