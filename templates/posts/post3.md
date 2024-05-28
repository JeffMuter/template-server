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

1. What is the purpose of the net/http package in Go?

The purpose of this package in Go is to give the project a set of types, interfaces, methods and functions needed for a wide variety of work on, in my experience, web servers. The best parts, that I have used the most, are in creating handlers, serving html, and creating http status codes. There is also great toosl for dealing with sessions and cookies to get a powerful, and scalable web server running quickly.

2. How do you start a basic HTTP server using the net/http package? Describe the essential functions and their roles.

The core parts of an http server are going to be mux, ListenAndServe, a router, and handlers. ListenAndServe wil create the active port for the server to use, the router and mux work together to take requests, and find the correct handler to generate the correct response, then we create the response and send it back to the user.

3. Explain the use of http.HandleFunc and how it differs from http.Handle.

Handle and HandlerFunc are very interesting. I wrote a blog post about them recently. They taught me a lot as a golang developer about how amazingly powerful it is to create your own types, very flexible, but the way the code was written, the code was short and simple, yet could do so much. Anyways, Handle takes in a string path, and a Handler type. Handler type is any type that has a ServeHTTP method attached. Where HandleFunc, swaps out the Handler for a function that is classified as a Handler because it satisfies the Handler interface. So it's more specific in that it must be a function. Handle in my experience, is more commonly used with middleware. And HandleFunc is what I use when I am passing in a Handler function.

4. How would you set and retrieve cookies in an HTTP handler?

net/http has great tooling for working with cookies. For your typical request type that comes in, you have a field of, let's say, r.Cookie, that you can read from, and on your response you can write to it in the same way. I've mostly had experience with this in writing authentication middleware.

5. What are the key differences between http.Request and http.ResponseWriter? How are they typically used in a handler function?

Very fun questions. http.Request contains information about the request from the user. A lot of information can be found withing, cookies, sessions details, JWT, and other metadata properties. You often use information that exists from the front-end to determine what kind of response you need to generate. Whether that information be in the path of the request, or stored in other structs. Whereas http.ResponseWriter is how we create our response. Again, storing data, serving html files, authentication, and all other informatio we want to return to the user. I guess, to simplify, they are technically the same kind of structure. They are http transfer data structures, however, one comes from the user, the other contains the information and status codes we plan to send back to the user.

6. Describe how you would implement basic request logging for an HTTP server.

Great question, unfortunately, I haven't implemented logging yet, so I won't try to speak on this.

7. How do you handle URL parameters and query strings in a Go HTTP server? Provide an example.

I have spent the most time building my own projects, and since I never built a go web server before 1.21 , I wanted to spend as much time working without abstractions like web frameworks like Gin or Echo. So url parameters exist as wildcards in the Url path. Such as /{userId}/{newFriendId}. Both of these would be int wildcards on a path. net/http provides great tooling to deal with the path variables found, that makes them very easy to work with.

8. Explain the concept of middleware in the context of an HTTP server. How would you implement middleware in Go?

In the past, I used Middleware almost exclusively for authentication. What I do is protect my Handler functions behind a Middleware layer to check the session and cookie of the request, before calling the Handler. If the user is validated, send back the Handler of the path they were intending to travel to. If the user was not validated, send back some other path we want the user to go to if their request failed.

To be a little more technical. Unprotected paths have http.HandlerFunc() in the router. Protected routes look like http.Handle("path", MiddlewareFunc()). The middleware returns a Handler type, and takes in func(http.ResponseWriter, *http.Request). From that HandlerFunc we pass in, we generate the correct Handler to return after we have validated the user in the Middleware function. Personally, I had a rediculously hard time understanding Middleware when I first came to Golang, after spending a lot of time studying the net/http package, and writing a blog article on it, I feel much better about it, although I'd like to explore other kinds of middleware in the future.

9. How can you create a custom HTTP client using the net/http package? What are some scenarios where you might need a custom client?

I am not sure I understand the question, so I will refrain from responding here.

10. What are some common security considerations when building HTTP servers in Go, and how can you address them using the net/http package?

I think most of the security concerns are addressed in terms of golang having great sexurity packages already for hashing. The net/http package have fantastic middleware capabilities, which is everything that I have used. Golang being a statically typed language helps protect you from mailicious data stored in the url. The main security concern I think I have heard where people need to be more protective, is defending your handlers from SQL queries stored in the path. Always worth keeping an eye on.