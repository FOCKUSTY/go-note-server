package hello

type Routes struct {
	GET string
}

func CreateRoutes() Routes {
	return Routes{
		GET: "/",
	}
}
