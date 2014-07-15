ginAuth Example
========

This folder includes a simple example RESTful server with ginAuth already plugged in.  It requires a 
Postgres SQL database with a table simply called `users` with three fields:

1. `id` - primary key int 
2. `email` - character varying(255)
3. `password` - character varying(60) 

Of course you could easily plug in another SQL database really easily.  See 
[the Beego ORM documentation](http://beego.me/docs/mvc/model/overview.md) for more information.