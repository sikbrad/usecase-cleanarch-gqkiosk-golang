package dataprovider

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
	"log"
)

type sqlite_datasource struct {
	DB *sql.DB
}

func (datasource *sqlite_datasource) GetAllFoodMenu() *[]entity.FoodMenu {
	var foodMenus []entity.FoodMenu

	rows, _ := datasource.DB.Query("SELECT id, name, description, image_url FROM food_menus")
	for rows.Next() {
		var foodMenu entity.FoodMenu
		rows.Scan(&foodMenu.Id, &foodMenu.Name, &foodMenu.Description, &foodMenu.ImageUrl)
		//fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)

		foodMenus = append(foodMenus, foodMenu)
	}

	return &foodMenus
}

func (datasource *sqlite_datasource) GetFoodMenuWithName(searchName string) *entity.FoodMenu {
	var foodMenu entity.FoodMenu
	row := datasource.DB.QueryRow(`SELECT id, name, description, image_url FROM food_menus WHERE name LIKE $1`, searchName)
	err := row.Scan(&foodMenu.Id, &foodMenu.Name, &foodMenu.Description, &foodMenu.ImageUrl)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("error in GetFoodMenuWithName : ", err.Error())
	}

	return &foodMenu
}

func NewCustomerOrderingSqliteDataStorage() DataProvider {
	database, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect to database")
	}

	//@ref https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS food_menus (id TEXT PRIMARY KEY,name TEXT, description TEXT, image_url TEXT);")
	statement.Exec()

	if err != nil {
		panic("failed to provisioning the database")
	}

	return &sqlite_datasource{
		DB: database,
	}
}
