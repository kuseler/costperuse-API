// db-interaction
package DBInteraction

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Item struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Uses         int       `json:"Uses"`
	DateBought   time.Time `json:"date_bought"`
	Notes        string    `json:"notes"`
	Category     int       `json:"category"`
	CostPerUse   float64   `json:"cost_per_use"`
	CostPerMonth float64   `json:"cost_per_month"`
}

func getMonthDiffFromToday(date time.Time) int64 {
	duration := time.Since(date)
	fmt.Println(duration)
	// an average month is 30.4375 days (365.25 / 12), to make the months more comparable
	averageMonths := int64(duration.Hours() / 30.4375 / 24)
	fmt.Println(averageMonths)
	return averageMonths
}

func getCategoryIDByName(conn *sql.DB, categoryName string) int {
	rows, err := conn.Query("SELECT id FROM categories WHERE category_name=$1", categoryName)
	if err != nil {
		log.Fatal(err)
	}
	categories := new([]string)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(categories)
	}
	fmt.Println(categories)
	tmp := *categories
	categoryID, err := strconv.Atoi(tmp[0])
	return categoryID
}

func getItemsByCategory(conn *sql.DB, category string) json.RawMessage {
	id := GetCategoryItemsAsList(conn, category)
	temp := getQueryResultAsItemList(conn, string(id[0]))
	marshaled, err := json.Marshal(temp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	return marshaled
}

func getQueryResultAsItemList(conn *sql.DB, sqlStatement string) []Item {
	// not truly universal, as it expects the statement "SELECT * FROM items WHERE "
	rows, err := conn.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	fmt.Printf("rows:%v\n", rows)
	defer rows.Close()
	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.ID, &temp.Name, &temp.Price, &temp.Uses, &temp.DateBought, &temp.Notes, &temp.Category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("temp:%v\n", temp)

		averageMonths := getMonthDiffFromToday(temp.DateBought)
		temp.CostPerMonth = temp.Price / float64(averageMonths)
		temp.CostPerUse = temp.Price / float64(temp.Uses)
		items = append(items, temp)

	}
	return items
}

func GetCategoryItemsAsList(conn *sql.DB, categoryIDOrName string) []Item {
	numericID, err := strconv.Atoi(categoryIDOrName)
	if err == nil {
		log.Printf("category supplied by ID %v", numericID)
	} else {
		numericID = getCategoryIDByName(conn, categoryIDOrName)
	}
	query := fmt.Sprintf("SELECT * FROM items WHERE category=%v", numericID)
	items := getQueryResultAsItemList(conn, query)
	return items
}
