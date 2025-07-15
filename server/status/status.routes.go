package status

type Routes struct {
	GET          string
	GET_DATABASE string
}

func GetRoutes() *Routes {
	return &Routes{
		GET:          "/",
		GET_DATABASE: "/database/",
	}
}
