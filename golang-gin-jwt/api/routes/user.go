package routes

import "log"

type UserRoutes struct{}

func NewUserRoutes() UserRoutes {
	return UserRoutes{}
}

// Setup activityParticipation routes
func (c UserRoutes) Setup() {
	log.Printf("hello user route")
}
