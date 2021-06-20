// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnVillages(w http.ResponseWriter, r *http.Request) {
	var villages []Village
	var response Response
	var metaInfo MetaInfo
	var countTotal int = 0
	var ipp = 20
	var name string
	var cp = 1

	keys := r.URL.Query()

	nam, ok := keys["name"]
	if !ok {
		log.Println("Url Param 'name' is missing")
		fmt.Println("Url Param 'name' is missing")
		return
	} else {
		name = strings.ToLower((nam[0]))
	}

	numPage, ok := keys["page"]
	if ok {
		i, err := strconv.Atoi(string(numPage[0]))
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		} else {
			cp = i
		}
	}

	pageSize, ok := keys["page_size"]
	if ok {
		i, err := strconv.Atoi(string(pageSize[0]))
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		} else {
			ipp = i
		}
	}

	query := `
		SELECT DISTINCT villages.id AS id, villages.name AS name, districts.name AS district, regencies.name AS regency, provinces.name AS province
		FROM villages
			INNER JOIN districts
				ON villages.district_id = districts.id
			INNER JOIN regencies
				ON districts.regency_id = regencies.id
			INNER JOIN provinces
				ON regencies.province_id = provinces.id
		WHERE villages.name LIKE '%` + name + `%' LIMIT ` + strconv.Itoa((cp-1)*ipp) + `, ` + strconv.Itoa(ipp) + `
		`

	query_count := `
		SELECT COUNT(DISTINCT(villages.id)) AS count_distinct
		FROM villages
			INNER JOIN districts
				ON villages.district_id = districts.id
			INNER JOIN regencies
				ON districts.regency_id = regencies.id
			INNER JOIN provinces
				ON regencies.province_id = provinces.id
		WHERE LOWER(villages.name) LIKE '%` + name + `%'
		`

	db := connect()
	defer db.Close()

	row, err := db.Query(query_count)
	if err != nil {
		log.Print(err)
		fmt.Print(err)
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&countTotal)
		if err != nil {
			log.Fatal(err.Error())
			fmt.Print(err)
		}
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
		fmt.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		var village Village
		err = rows.Scan(&village.Id, &village.Name, &village.District, &village.Regency, &village.Province)
		if err != nil {
			log.Fatal(err.Error())
			fmt.Print(err)
		} else {
			villages = append(villages, village)
		}
	}

	response.Data = villages
	metaInfo.ItemsPerPage = ipp
	metaInfo.CurrentPage = cp
	metaInfo.TotalItems = countTotal
	metaInfo.TotalPage = int(math.Ceil((float64(countTotal) / float64(ipp))))
	response.Meta = metaInfo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/villages", returnVillages)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Server running...")
	handleRequests()
}
