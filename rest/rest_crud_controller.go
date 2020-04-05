package rest

import (
	"github.com/coupa/foundation-go/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrudController struct {
	persistenceService persistence.PersistenceService
	httpQueryParser    HttpQueryParser
}

func NewCrudController(persistenceManager persistence.PersistenceService, httpQueryParser HttpQueryParser) *CrudController {
	if httpQueryParser == nil {
		httpQueryParser = &HttpQueryParserRailsActiveAdmin{}
	}
	return &CrudController{
		persistenceService: persistenceManager,
		httpQueryParser:    httpQueryParser,
	}
}

func (self *CrudController) FindOne(c *gin.Context) {
	id := c.Param("id")
	obj, err := self.persistenceService.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "...")
		return
	}

	//if obj == nil {
	//	c.JSON(http.StatusNotFound, "...")
	//	return
	//}

	c.JSON(http.StatusOK, obj)
}

//equivalent to a rails "index" controller action
func (self *CrudController) FindMany(c *gin.Context) {
	persistenceQuery, err := self.httpQueryParser.Parse(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	objs, err := self.persistenceService.FindMany(persistenceQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, objs)
}

func (self *CrudController) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	obj := self.persistenceService.NewModelObjPtr()
	err := self.persistenceService.FindOneLoad(id, obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	validationErrors, _ := self.persistenceService.Validate(obj)
	if validationErrors.HasErrors() {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	_, err = self.persistenceService.UpdateOne(id, obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (self *CrudController) CreateOne(c *gin.Context) {
	obj := self.persistenceService.NewModelObjPtr()

	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	validationErrors, _ := self.persistenceService.Validate(obj)
	if validationErrors.HasErrors() {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	err := self.persistenceService.CreateOne(obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (self *CrudController) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	_, err := self.persistenceService.DeleteOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
