package router

func InitializeRouter() {
	InitializeDocRoute()
	InitializeTopGroupRouter()
	InitializeMiddleGroupRouter()
	InitializeEndpointRouter()
}
