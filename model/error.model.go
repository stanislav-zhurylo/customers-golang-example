package model

type ErrorMsg struct {
	Field   string `schema:"field"`
	Message string `schema:"message"`
}
