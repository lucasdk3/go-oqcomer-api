package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {

	// ctx := context.Background()
	// l, err := ngrok.Listen(ctx,
	// 	config.HTTPEndpoint(),
	// 	ngrok.WithAuthtokenFromEnv(),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("ngrok ingress url: ", l.Addr())
	// http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello from your ngrok-delivered Go app!")
	// }))

	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))

}
