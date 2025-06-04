package response

type Response struct {
	Status bool 
	Message string 
	Data interface{} 
	Error interface{}
}