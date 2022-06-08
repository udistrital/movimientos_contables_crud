package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetExistByCuenta",
            Router: "/exist/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:SaldoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:SaldoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:SaldoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:SaldoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_contables_crud/controllers:TransaccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
