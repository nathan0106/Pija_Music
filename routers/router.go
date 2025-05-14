// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/Pija_Music/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Artista",
			beego.NSInclude(
				&controllers.ArtistaController{},
			),
		),

		beego.NSNamespace("/Trajes_Tipicos",
			beego.NSInclude(
				&controllers.TrajesTipicosController{},
			),
		),
		beego.NSNamespace("/Administrador",
		beego.NSInclude(
			&controllers.AdministradorController{},
		),
	),

		beego.NSNamespace("/Comida",
			beego.NSInclude(
				&controllers.ComidaController{},
			),
		),

		beego.NSNamespace("/Lugares",
			beego.NSInclude(
				&controllers.LugaresController{},
			),
		),

		beego.NSNamespace("/Cultura",
			beego.NSInclude(
				&controllers.CulturaController{},
			),
		),

		beego.NSNamespace("/Estilo_Musical",
			beego.NSInclude(
				&controllers.EstiloMusicalController{},
			),
		),

		beego.NSNamespace("/Credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/Autor_Coplas",
			beego.NSInclude(
				&controllers.AutorCoplasController{},
			),
		),

		beego.NSNamespace("/Coplas",
			beego.NSInclude(
				&controllers.CoplasController{},
			),
		),

		beego.NSNamespace("/Canciones",
			beego.NSInclude(
				&controllers.CancionesController{},
			),
		),

		beego.NSNamespace("/Tipo_Cultura",
			beego.NSInclude(
				&controllers.TipoCulturaController{},
			),
		),

		beego.NSNamespace("/Tipo_Lugares",
			beego.NSInclude(
				&controllers.TipoLugaresController{},
			),
		),

		beego.NSNamespace("/Favoritos",
			beego.NSInclude(
				&controllers.FavoritosController{},
			),
		),

		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
