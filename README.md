# greetr
Greetr is a cloud native Hello World API powered by a system of highly-scalable, event-driven microservices that provide socially responsible greetings classified by our semi-supervised deep learning algorithms. DevOps. 
# Running/building Greetr
Greetr just needs Go. Once you have that installed you can either run it directly:

`go run main.go`

Or you can build an executable:

`go build main.go`

You can also override the port/IP Greetr serves requests on with the `-bind-addr` flag like so

`go run main.go -bind-addr 127.0.0.1:8888`

You can now access the Greetr service on port 8888 of your loopback interface as opposed to 8080 on all of them.