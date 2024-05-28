Handle : a function

func Handle(pattern string, handler Handler)

HandleFunc :

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

Handler : type interface. Simply need a method called ServeHTTP that has params of response and *request

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

HandlerFunc: a type, because it has a ServeHTTP method, HandlerFunc is a Handler type!

type HandlerFunc func(ResponseWriter, *Request)

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

func (HandlerFunc) ServeHTTP : the method for HandlerFunc types.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
ServeHTTP calls f(w, r).

