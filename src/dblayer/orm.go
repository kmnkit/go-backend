import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

// NewORM 새 ORM 커넥션
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

// GetAllProducts 모든 상품을 반환
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

// GetPromos 모든 프로모션을 반환
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// GetCustomerByName 이름이으로 고객을 반환
func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{Firstname: firstname, Lastname: lastname}).Find(&customer).Error
}

// GetCustomerById id가 가리키는 고객을 반환
func (db *DBORM) GetCustomerById(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

// GetProduct ID가 가리키는 상품을 반환
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

// AddUser User 추가
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, err error){
	hashPassword(&customer.Pass)
	customer.LoggedIn = True
	return customer, db.Create(&customer).Error
}

// SignInUser 로그인
func (db *DBORM) SignInUser(email, pass string)(customer models.Customer, err error){
	if !checkPassword(pass){
		return customer, errors.New("Invalid Password")
	}
	// 사용자 행을 나타내는 *gorm.DB 타입
	result := db.Table("Customers").Where(&models.Customer{Email:email})
	// loggedin 필드 업데이트
	err = result.Update("loggedin", 1).Error
	if err != nil{
		return customer, err
	}
	// 사용자 행 반환
	return customer, result.Find(&customer).Error
}

// SignOutUserById ID로 사용자를 로그아웃 처리
func (db *DBORM)SignOutUserById(id int) error{
	// ID에 해당하는 사용자 구조체 생성
	customer := models.Customer{
		Model:gorm.Model{
			ID: uint(id),
		},
	}
	// 사용자의 상태를 로그아웃 상태로 업데이트
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

// GetCustomerOrdersByID ID로 고객 주문을 가져온다.
func (db *DBORM) GetCustomerOrdersByID(id int)(orders []models.Order, err error){
	return orders, db.Table("orders").Select("*")
	.Joins("join customers on customers.id = customer_id")
	.Joins("Join products on products.id = product_id")
	.Where("customer_id=?", id).Scan(&orders).Error
}