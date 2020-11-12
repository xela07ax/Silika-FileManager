package r7uter

type Router struct {
	name string
	loger chan <- [4]string
}
func NewRouterManager(loger chan <- [4]string) *Router {
	return &Router{
		name: 		"RouteManager",
		loger: 		loger,
	}
}