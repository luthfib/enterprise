package routes

import (
	"enterprise/routes/api"
	authenticatedroutes "enterprise/routes/authenticated"
	"enterprise/routes/middlewares/auth"
	"enterprise/routes/middlewares/tenant"
	publicroutes "enterprise/routes/public"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func Database(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func Validator(valid *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Validator", valid)
		c.Next()
	}
}

func InitRouter(db *mongo.Client) *gin.Engine {
	validate = validator.New()
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		fmt.Print(err)
		// Handle the error. Probably bail out if we can't connect.
	}

	store.Options(sessions.Options{
		//Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	})

	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	r.Use(Database(db))
	r.Use(Validator(validate))
	r.LoadHTMLGlob("build/index.html")
	r.Use(static.Serve("/static", static.LocalFile("./build/static", true)))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	public := r.Group("/")
	publicroutes.Bootstrap(public)

	private := r.Group("/")
	private.Use(auth.Required())
	private.Use(tenant.GetContext())
	authenticatedroutes.Bootstrap(private)

	apiv1 := r.Group("/api")
	apiv1.Use(auth.Required())
	apiv1.Use(tenant.GetContext())
	api.Bootstrap(apiv1)

	return r
}
