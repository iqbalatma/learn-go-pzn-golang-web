Golang already has their own web server, so we do not need additional web server like apache or nginx
Golang has package called net/http, so we do not need any third party package to create golang web

To create web we must create server as well
when create a server we need to define host and port

server only receive request
to execute request, we need handler

handler function does not support multiple endpoint
we can use alternative for handler with ServerMux

Longer url pattern will be priority to execute

Header is not case sensitive