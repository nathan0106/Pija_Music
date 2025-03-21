package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoCulturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TipoLugaresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Pija_Music/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
