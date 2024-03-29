package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type Movimiento struct {
	Id                int          `orm:"column(id);pk;auto"`
	TerceroId         int          `orm:"column(tercero_id);null"`
	CuentaId          string       `orm:"column(cuenta_id)"`
	NombreCuenta      string       `orm:"column(nombre_cuenta)"`
	TipoMovimientoId  int          `orm:"column(tipo_movimiento_id)"`
	Valor             float64      `orm:"column(valor);digits(20);decimals(7)"`
	Descripcion       string       `orm:"column(descripcion);null"`
	Activo            bool         `orm:"column(activo)"`
	FechaCreacion     string       `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string       `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	TransaccionId     *Transaccion `orm:"column(transaccion_id);rel(fk)"`
	SaldoAnterior     float64      `orm:"column(saldo_anterior);digits(20);decimals(7);null"`
	NuevoSaldo        float64      `orm:"column(nuevo_saldo);digits(20);decimals(7);null"`
}

func (t *Movimiento) TableName() string {
	return "movimiento"
}

func init() {
	orm.RegisterModel(new(Movimiento))
}

// AddMovimiento insert a new Movimiento into database and returns
// last inserted Id on success.
func AddMovimiento(m *Movimiento) (id int64, err error) {
	s := Saldo{
		CuentaId:          m.CuentaId,
		Debito:            0,
		Credito:           0,
		Saldo:             0,
		FechaCreacion:     time_bogota.TiempoBogotaFormato(),
		FechaModificacion: time_bogota.TiempoBogotaFormato(),
	}
	valor := m.Valor
	debito := m.TipoMovimientoId == 344
	o := orm.NewOrm()
	if err = o.Begin(); err == nil {

		if _, _, err = o.ReadOrCreate(&s, "cuenta_id"); err != nil {
			o.Rollback()
			return
		}
		if debito {
			s.Debito += valor
		} else {
			s.Credito += valor
		}
		m.SaldoAnterior = s.Saldo
		s.Saldo = s.Debito - s.Credito
		m.NuevoSaldo = s.Saldo

		s.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if _, err = o.Update(&s, "debito", "credito", "fecha_modificacion", "saldo"); err != nil {
			o.Rollback()
			return
		}
		if id, err = o.Insert(m); err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}
	return
}

// GetMovimientoById retrieves Movimiento by Id. Returns error if
// Id doesn't exist
func GetMovimientoById(id int) (v *Movimiento, err error) {
	o := orm.NewOrm()
	v = &Movimiento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovimiento retrieves all Movimiento matches certain condition. Returns empty list if
// no records exist
func GetAllMovimiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Movimiento)).RelatedSel()
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

	var l []Movimiento
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

// UpdateMovimiento updates Movimiento by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovimientoById(m *Movimiento) (err error) {
	o := orm.NewOrm()
	v := Movimiento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovimiento deletes Movimiento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovimiento(id int) (err error) {
	o := orm.NewOrm()
	v := Movimiento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Movimiento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetExistMovimientosByCuentaId ...
func GetExistMovimientosByCuentaId(cuentaId string) (exist bool, err error) {
	o := orm.NewOrm()
	var movimientos []Movimiento
	_, err = o.QueryTable("movimiento").Filter("cuenta_id", cuentaId).Limit(1, 0).All(&movimientos)
	if err != nil {
		return
	}
	if len(movimientos) > 0 {
		exist = true
	}
	return
}
