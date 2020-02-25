package rest

import "net/http"
import "github.com/coupa/foundation-go/persistence"

type BaseCrudHandler struct {
	PersistenceManager persistence.PersistenceManager
}

//equivalent to a rails "show" controller action
func (self *BaseCrudHandler) FindOne(c *gin.Context) {
	id := c.Param("id")
	obj, err := self.PersistenceManager.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "...")
		return
	}

	if obj == nil {
		c.JSON(http.StatusNotFound, "...")
		return
	}

	c.JSON(http.StatusOK, obj)
}

//equivalent to a rails "index" controller action
func (self *BaseCrudHandler) FindMany(c *gin.Context) {
	params := ConvertHttpQueryParamsToPersistenceParams(c)
	objs, err := self.PersistenceManager.FindMany(params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "...")
		return
	}

	c.JSON(http.StatusOK, objs)
}

//equivalent to a rails "update" controller action
func (self *BaseCrudHandler) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	obj, err := self.PersistenceManager.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "...")
		return
	}

	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, "...")
		return
	}

	validationErrors := self.PersistenceManager.Validate(obj)
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, "...")
		return
	}

	err = self.PersistenceManager.UpdateOne(obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "...")
		return
	}

	c.JSON(http.StatusOK, obj)
}

//equivalent to a rails "create" controller action
func (self *BaseCrudHandler) CreateOne(c *gin.Context) {
	//...
}

//equivalent to a rails "destroy" controller action
func (self *BaseCrudHandler) DeleteOne(c *gin.Context) {
	//...
}
