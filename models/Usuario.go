package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Usuario struct {
	Id                 int       `orm:"column(Id_Usuario);pk;auto"`
	Nombres            string    `orm:"column(Nombres)"`
	Apellido           string    `orm:"column(Apellido)"`
	IdCredenciales    *Credenciales  `orm:"column(Id_Credenciales);rel(fk)"`
	Email              string    `orm:"column(Email)"`
	Activo             bool      `orm:"column(Activo)"`
	FechaCreacion      time.Time `orm:"column(Fecha_Creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(Feccha_Modificasion);type(timestamp with time zone);auto_now"`
}

func (t *Usuario) TableName() string {
	return "Usuario"
}

func init() {
	orm.RegisterModel(new(Usuario))
}

// AddUsuario insert a new Usuario into database and returns
// last inserted Id on success.
func AddUsuario(m *Usuario) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	id, err = o.Insert(m)
	return
}

// GetUsuarioById retrieves Usuario by Id. Returns error if
// Id doesn't exist
func GetUsuarioById(id int) (v *Usuario, err error) {
	o := orm.NewOrm()
	v = &Usuario{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsuario retrieves all Usuario matches certain condition. Returns empty list if
// no records exist
func GetAllUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Usuario)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []Usuario
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUsuario updates Usuario by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsuarioById(m *Usuario) (err error) {
	o := orm.NewOrm()
	m.Activo = true
	v := Usuario{Id: m.Id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsuario deletes Usuario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsuario(id int) (err error) {
	o := orm.NewOrm()
	v := Usuario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Usuario{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
