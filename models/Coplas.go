package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Coplas struct {
	Id                int          `orm:"column(Id_Coplas);pk;auto"`
	NombreCopla       string       `orm:"column(Nombre_Copla)"`
	DescripcionCoplas string       `orm:"column(Descripcion_Coplas)"`
	Autor             *AutorCoplas `orm:"column(Autor);rel(fk)"`
	Origen            string       `orm:"column(Origen)"`
	Tema              string       `orm:"column(Tema)"`
	Ocasion           string       `orm:"column(Ocasion);null"`
	Activo            bool         `orm:"column(Activo)"`
	FechaCreacion     time.Time    `orm:"column(Fecha_Creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion time.Time    `orm:"column(Fecha_Modificacion);type(timestamp with time zone);auto_now"`
}

func (t *Coplas) TableName() string {
	return "Coplas"
}

func init() {
	orm.RegisterModel(new(Coplas))
}

// AddCoplas insert a new Coplas into database and returns
// last inserted Id on success.
func AddCoplas(m *Coplas) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	id, err = o.Insert(m)
	return
}

// GetCoplasById retrieves Coplas by Id. Returns error if
// Id doesn't exist
func GetCoplasById(id int) (v *Coplas, err error) {
	o := orm.NewOrm()
	v = &Coplas{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCoplas retrieves all Coplas matches certain condition. Returns empty list if
// no records exist
func GetAllCoplas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Coplas)).RelatedSel()
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Coplas
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

// UpdateCoplas updates Coplas by Id and returns error if
// the record to be updated doesn't exist
func UpdateCoplasById(m *Coplas) (err error) {
	o := orm.NewOrm()
	m.Activo = true
	v := Coplas{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCoplas deletes Coplas by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCoplas(id int) (err error) {
	o := orm.NewOrm()
	v := Coplas{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Coplas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
