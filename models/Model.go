package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
	"time"
)

var db *gorm.DB

func init() {
	//创建一个数据库的连接
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	//迁移the schema
	db.AutoMigrate(&Sale{})
	db.AutoMigrate(&User{})
}
type Sale struct {
	gorm.Model  // ID 创建时间 修改时间 删除时间
	CreatUser string  `gorm :"type:varchar(10)" json:"creat_user"`  // 创建人
    SaleNumber string `gorm:"type:varchar(10)" json:"sale_number"` // 销售单号
	Client string `grom:"type:varchar(30)" json:"client" form:"client"`// 客户名称
	City string `json:"city"`
	BillingTime time.Time `json:"billing_time"`// 开单时间
	ContractPeriod string `json:"contract_period"` // 合同账期
	AccountPeriod int `json:"account_period"`//账期
	Merchandiser string `json:"merchandiser"`// 跟单员
	Salesman string `json:"salesman"`// 业务员
	Currency string `json:"currency"`// 币种
	UnitPrice float64 `json:"unit_price"`// 单价
	Quantity float64 `json:"quantity"`// 数量
	AmountReceivable float64 `json:"amount_receivable"`// 应收金额
	Invoice int `json:"invoice"`// 发票
	PaidAmount float64 `json:"paid_amount"`// 实收金额
	UncollectedAmount float64 `json:"uncollected_amount"`// 未收金额
	DueDate  time.Time `json:"due_date"`// 应收日期
	CollectionDate  time.Time `json:"collection_date"`// 收款日
	CollectionAmount  float64 `json:"collection_amount"`// 收款金额
	TimeOut time.Time `json:"time_out"`// 超时
	Remarks string `json:"remarks"`// 备注
}

type User struct {
	Name string `json:"name"` // 登录用户
	Password string `json:"password"`// 登录密码
	LoginDate time.Time // 登录时间
	Permission int `json:"permission"`// 权限
}

//  返回所有用户列表
func fetchUserList(c *gin.Context){
	var userlist []User
	db.Find(&userlist)
	//fmt.Println(userlist)
	if len(userlist) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"没有用户！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"data":userlist,
	})
}


// 插入一条用户记录
func createUser(c *gin.Context){
	permission, _ := strconv.Atoi(c.PostForm("permission"))
	user := User{Name: c.PostForm("name"), Password: c.PostForm("password"), Permission:permission}
	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!"})

}

// 根据用户名查询
func fetchSingleUser(c *gin.Context){
	var user User
	Uname := c.Param("name")
	db.Where(&User{Name: Uname}).Find(&user)
	fmt.Println(Uname)

	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// 更新匹配到的用户数据
func updateUser(c *gin.Context)  {
	var user User
	uname := c.Param("name")
	db.Where(&User{Name: uname}).Find(&user)
	fmt.Println(user)

	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "未找到!"})
		return
	}
	db.Model(&user).Where("name = ?", uname).Update("password",c.PostForm("password"))
	permission, _ := strconv.Atoi(c.PostForm("permission"))
	db.Model(&user).Where("name = ?", c.Param("name")).Update("permission", permission)
	fmt.Println(c.PostFormArray("password"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "用户信息更新成功!"})
}


// 删除用户
func deleteUser(c *gin.Context) {
	var user User
	Uname := c.Param("name")
	db.Where(&User{Name: Uname}).Find(&user)
	fmt.Println(user)

	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "未找到!"})
		return
	}

	db.Model(&user).Where("name = ?", c.Param("name")).Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

// 返回所有订单 Salesorder
func fetchSaleList(c *gin.Context){
	var salelist []Sale
	db.Find(&salelist)
	if len(salelist) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"没有数据！"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"data":salelist,
	})
}

//// 插入一条订单记录
//func createSalesorder(c *gin.Context){
//	//permission, _ := strconv.Atoi(c.PostForm("permission"))
//	//user := User{Name: c.PostForm("name"), Password: c.PostForm("password"), Permission:permission}
//	salesorder := Sale{e: c.PostForm("name"), Password: c.PostForm("password")}
//	db.Save(&salesorder)
//	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": ""})
//
//}

// 查询订单
func fetchSalesorder(c *gin.Context){
	var sale []Sale
	salenumber := c.Query("sale_number")
	client := strconv.QuoteToASCII(c.Query("client"))
	fmt.Println(salenumber)
	client, _ =strconv.Unquote(client)

	if client == "" {
		db.Where("sale_number = ?", salenumber ).Find(&sale)
		if len(sale) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "没找到该订单!"})
			return
		}
	} else {
		db.Where("client LIKE ?","%"+client+"%").Find(&sale)
		if len(sale) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "没找到该客户订单，请换个描述方式!"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": sale})
}

// 更新订单数据
func updateSales(c *gin.Context)  {
	var sale []Sale
	salenum := c.PostForm("sale_number")
	db.Where(&Sale{SaleNumber: salenum}).Find(&sale)

	if len(sale) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "未找到!"})
		return
	}
	for key, sa := range sale {
		fmt.Println(key, sa.SaleNumber)

		//fmt.Println(c.PostForm("id"))
		//db.Model(&sale).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
		db.Model(&sale[key]).Where("sale_number = ?", sa.SaleNumber).Update("collection_date",
			c.PostFormArray("collection_date")[key])
		//fmt.Println(sa.SaleNumber, sa.CollectionDate)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "更新成功!"})
}


// 删除订单
func deleteSale(c *gin.Context) {
	var sale Sale
	salenumber := c.Param("sale_number")
	db.Where(&Sale{SaleNumber:salenumber}).First(&sale)
	fmt.Println(sale)

	if sale.SaleNumber == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "未找到!"})
		return
	}

	db.Delete(&sale)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "删除成功！"})
}



func main(){
	router := gin.Default()

	v1 := router.Group("/v1/user")
	{
		v1.GET("/", fetchUserList)
		v1.GET("/:name", fetchSingleUser)
		v1.POST("/", createUser)
		v1.PUT(":name",updateUser)
		v1.DELETE("/:name", deleteUser)
	}
	sa := router.Group("/sale")
	{
		sa.GET("/",fetchSaleList)
		sa.GET("/q", fetchSalesorder)
		sa.PUT("/:client", updateSales)
		sa.DELETE("/:sale_number", deleteSale)
	}

	router.Run()
}


