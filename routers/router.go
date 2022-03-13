// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/movimientos_contables_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/transaccion",
			beego.NSInclude(
				&controllers.TransaccionController{},
			),
		),

		beego.NSNamespace("/movimiento",
			beego.NSInclude(
				&controllers.MovimientoController{},
			),
		),
		beego.NSNamespace("/saldo",
			beego.NSInclude(
				&controllers.SaldoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
