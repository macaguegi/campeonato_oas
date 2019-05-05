package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:EquiposController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:EquiposController"],
        beego.ControllerComments{
            Method: "InsertarEquipo",
            Router: `/InsertarEquipo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:EquiposController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:EquiposController"],
        beego.ControllerComments{
            Method: "ConsultarEquipos",
            Router: `/consultarEquipos`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "ConsultarCiudades",
            Router: `/consultarCiudades`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "ConsultarPaises",
            Router: `/consultarPaises`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "InsertarCiudad",
            Router: `/insertarCiudad`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "InsertarPais",
            Router: `/insertarPais`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:PartidosController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:PartidosController"],
        beego.ControllerComments{
            Method: "InsertarPartido",
            Router: `/InsertarPartido`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:PartidosController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:PartidosController"],
        beego.ControllerComments{
            Method: "ConsultarPartidos",
            Router: `/consultarPartidos`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/macaguegi/sudamericana_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
