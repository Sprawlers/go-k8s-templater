package main

func (s *Server) routes() {
    s.router.HandlerFunc("GET", "/health", s.handleHealtCheck())
    s.router.HandlerFunc("GET", "/test", s.handleTest())

    s.router.POST("/webhook", s.handleWebhook())
}
