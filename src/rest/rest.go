package rest

import "github.com/gin-gonic/gin"

// RunAPIWithHandler Handler를 인자로 받아 API를 실행함
func RunAPIWithHandler(address string, h HandlerInterface) error {
	// Gin 엔진
	r := gin.Default()
	// 상품 목록
	r.GET("/products", h.GetProducts)
	// 프로모션 목록
	r.GET("/promos", h.GetPromos)
	userGroup := r.Group("/user")
	{
		// 사용자 로그아웃 POST 요청
		userGroup.POST("/user/:id/signout", h.SignOut)
		// 구매 목록 조회
		userGroup.GET("/user/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		// 사용자 로그인
		usersGroup.POST("/users/signin", h.SignIn)
		// 사용자 로그인 POST 요청
		usersGroup.POST("/users/signup", h.SignUp)
		// 사용자 추가 POST 요청
		usersGroup.POST("/users", h.AddUser)
		// 결제 POST 요청
		usersGroup.POST("users/charge", h.Charge)
	}
	return r.Run(address)
}

// RunAPI HandlerInterface의 기본 구현을 나타냄
func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
