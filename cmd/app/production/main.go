package app

// @title gfvs
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// Handle deployment to Vercel Serverless
func func (w http.ResponseWriter, r *http.Request) {
  r.RequestURI = r.URL.String()

  server := server.New()
  server.RegisterFiberRoutes()
  server.Use(swagger.New(configs.SwgCfg))

  adaptor.FiberApp(server.App).ServeHTTP(w, r)
}

