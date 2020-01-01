package v1

import (
	"fmt"
	"go-example/internal/models"
	"go-example/internal/services"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.Debug("Init users")
}

// NewUserAPI create userService
func NewUserAPI(db *gorm.DB, autoMigrate bool) UserAPI {
	if autoMigrate {
		fmt.Println("Migrating User model")
		db.AutoMigrate(&models.User{})
		db.Create(models.User{
			Model: models.Model{ID: "1"}, Username: "utain", Email: "utain@gmail.com", Password: "P@55w0rd", Firstname: "Utain", Lastname: "Wongpreaw",
		})
	}
	return &userAPI{userService: services.NewUserService(db)}
}

//UserAPI interface
type UserAPI interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

// userAPI is a service private
type userAPI struct {
	userService services.UserService
}

// GetAllUser return all User
func (p userAPI) GetAllUser(c *gin.Context) {
	users := p.userService.GetAllUser(0, -1, "")
	c.JSON(http.StatusOK, users)
}

// GetUser return only one User
func (p userAPI) GetUser(c *gin.Context) {
	username := c.Param("name")
	if username == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	user := p.userService.GetUser(username)
	if reflect.DeepEqual(*user, models.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
