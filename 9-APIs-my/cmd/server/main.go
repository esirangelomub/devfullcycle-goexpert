package main

import (
	"github.com/esirangelomub/devfullcycle-goexpert/9-APIs-my/configs"
	_ "github.com/esirangelomub/devfullcycle-goexpert/9-APIs-my/docs"
	"github.com/esirangelomub/devfullcycle-goexpert/9-APIs-my/internal/entity"
	"github.com/esirangelomub/devfullcycle-goexpert/9-APIs-my/internal/infra/database"
	"github.com/esirangelomub/devfullcycle-goexpert/9-APIs-my/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title My Go Expert API Example
// @version 1.0
// @description Product API with authentication and authorization using JWT
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email eduardo.sirangelo@gmail.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(MyLogMiddleware)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))
	// Aula APIs -> Criando e trabalhando com middlewares
	// Check for more chi middlewares

	r.Route("/products", func(r chi.Router) {
		// middleware to verify and authenticate JWT
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func MyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("my-log-middleware -> request: %s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
