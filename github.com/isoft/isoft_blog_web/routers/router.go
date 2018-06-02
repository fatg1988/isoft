package routers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_blog_web/controllers"
)

func init() {
	beego.Router("catalog/edit", &controllers.CatalogController{}, "get:Edit;post:PostEdit")
	beego.Router("catalog/list", &controllers.CatalogController{}, "get:List;post:PostList")
	beego.Router("catalog/delete", &controllers.CatalogController{}, "post:PostDelete")
	beego.Router("/catalog/more", &controllers.CatalogController{}, "get:More")

	beego.Router("blog/edit", &controllers.BlogController{}, "get:Edit;post:PostEdit")
	beego.Router("blog/list", &controllers.BlogController{}, "get:List;post:PostList")
	beego.Router("blog/delete", &controllers.BlogController{}, "post:PostDelete")
	beego.Router("blog/search", &controllers.BlogController{}, "get:Search")
	beego.Router("blog/publish", &controllers.BlogController{}, "post:PostPublish")
	beego.Router("blog/detail", &controllers.BlogController{}, "get:Detail")
}
