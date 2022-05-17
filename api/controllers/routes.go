package controllers

import "github.com/victorsteven/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/auth/signup", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")

	s.Router.HandleFunc("/Wilayah/get_prov", middlewares.SetMiddlewareAuthentication(s.GetProvs)).Methods("GET")
	s.Router.HandleFunc("/Wilayah/get_prov/{prop_id}", middlewares.SetMiddlewareAuthentication(s.GetProv)).Methods("GET")

	s.Router.HandleFunc("/Wilayah/get_city", middlewares.SetMiddlewareAuthentication(s.GetCities)).Methods("GET")
	s.Router.HandleFunc("/Wilayah/get_city/{kota_id}", middlewares.SetMiddlewareAuthentication(s.GetCity)).Methods("GET")

	s.Router.HandleFunc("/Wilayah/get_kec", middlewares.SetMiddlewareAuthentication(s.GetKecamatans)).Methods("GET")
	s.Router.HandleFunc("/Wilayah/get_kec/{kec_id}", middlewares.SetMiddlewareAuthentication(s.GetKecamatan)).Methods("GET")

	s.Router.HandleFunc("/Wilayah/get_neg", middlewares.SetMiddlewareAuthentication(s.GetNegaras)).Methods("GET")
	s.Router.HandleFunc("/Wilayah/get_neg/{id}", middlewares.SetMiddlewareAuthentication(s.GetNegara)).Methods("GET")

	s.Router.HandleFunc("/Wilayah/get_desa", middlewares.SetMiddlewareAuthentication(s.GetDesas)).Methods("GET")
	s.Router.HandleFunc("/Wilayah/get_desa/{desa_id}", middlewares.SetMiddlewareAuthentication(s.GetDesa)).Methods("GET")
}
