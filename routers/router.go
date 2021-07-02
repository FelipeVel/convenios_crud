// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/FelipeVel/convenios_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/convenio",
			beego.NSInclude(
				&controllers.ConvenioController{},
			),
		),

		beego.NSNamespace("/pais_categoria",
			beego.NSInclude(
				&controllers.PaisCategoriaController{},
			),
		),

		beego.NSNamespace("/tipo_categoria",
			beego.NSInclude(
				&controllers.TipoCategoriaController{},
			),
		),

		beego.NSNamespace("/tipo_movilidad",
			beego.NSInclude(
				&controllers.TipoMovilidadController{},
			),
		),

		beego.NSNamespace("/movilidad",
			beego.NSInclude(
				&controllers.MovilidadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
