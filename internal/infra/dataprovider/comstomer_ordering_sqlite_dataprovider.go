package dataprovider
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
)

type sqlite_datasource struct{
	DB *gorm.DB
}

func (datasource* sqlite_datasource) GetAllFoodMenu() *[]entity.FoodMenu {
	var foodMenus []entity.FoodMenu
	datasource.DB.Find(&foodMenus)
	return &foodMenus
}

func (datasource* sqlite_datasource) GetFoodMenuWithName(searchName string) *entity.FoodMenu {
	var foodMenu *entity.FoodMenu
	if err := datasource.DB.Where("id = ?", searchName).First(foodMenu).Error; err != nil {
		//err
	}
	return foodMenu
}

func NewCustomerOrderingSqliteDataStorage() (DataProvider){
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&entity.FoodMenu{})

	return &sqlite_datasource{
		DB:database,
	}
}
