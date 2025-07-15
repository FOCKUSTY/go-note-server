package note

type Routes struct {
	GET string
}

func GetRoutes() *Routes {
	return &Routes{
		GET: "/",
	}
}
