package controllers

import (
	"encoding/json"
	"errors"
	"github.com/sena_2824182/Pija_Music/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// LugaresController operations for Lugares
type LugaresController struct {
	beego.Controller
}

// URLMapping ...
func (c *LugaresController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Lugares
// @Param	body		body 	models.Lugares	true		"body for Lugares content"
// @Success 201 {int} models.Lugares
// @Failure 403 body is empty
// @router / [post]
func (c *LugaresController) Post() {
	var v models.Lugares
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLugares(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Dato registrado", "Data": v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Lugares by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Lugares
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LugaresController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLugaresById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Consulta correctamente", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Lugares
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Lugares
// @Failure 403
// @router / [get]
func (c *LugaresController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllLugares(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "todos los datos", "Data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Lugares
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Lugares	true		"body for Lugares content"
// @Success 200 {object} models.Lugares
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LugaresController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Lugares{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLugaresById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Datos actualizados", "Data": v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Lugares
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LugaresController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteLugares(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Se elimino corretamente", "Dato elminado con ID": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
