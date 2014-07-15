ginAuth
========

ginAuth is a simple, sessionless, encrypted cookie based authentication middleware for the 
[Gin Web Framework](https://github.com/gin-gonic/gin)!

## External Packages

ginAuth uses a few external packages to accomplish it's goal:

* [Beego - Configuration](http://beego.me/docs/module/config.md) - 
this package is used to load values from config files.
* [Gorilla Web Toolkit - Secure Cookie](http://www.gorillatoolkit.org/pkg/securecookie) -  we use the 
secure cookie package to encode and decode the cookie data.

## Getting Started
A quick overview of the package.  The package is very small and very documented so don't fear to read over it.

### Global Values
These values can be set:

    	CookieName string                       // the name of the cookie that will be used, default: "token"
    	ConfigPath string                       // path to config file, default: ""
    	ConfigType string                       // type of config file, default: "ini"
    	Prefix string                           // the key in ctx.Keys[] to use, default: ""
    	HashKey []byte                          // hash key for securecookie
    	BlockKey []byte                         // block key for securecookie
    	Expiration int64                        // time until the cookie expires in seconds, default: 604800
    	Unauthorized func(ctx *gin.Context)     // function called if user is not authorized
    	Authorized func(ctx *gin.Context)       // function called if user is authorized
    	SecureCookie *securecookie.SecureCookie // global secure cookie object

### Creating Encryption Keys
The package is mostly ready to go out of the box except for the fact that you must include a HashKey and a BlockKey.

From the Gorilla docs:

>The hashKey is required, used to authenticate the cookie value using HMAC. It is recommended to use a key with 32 or 
64 bytes.
>
>The blockKey is optional, used to encrypt the cookie value -- set it to nil to not use encryption. If set, the length 
must correspond to the block size of the encryption algorithm. For AES, used by default, valid lengths are 16, 24, or 
32 bytes to select AES-128, AES-192, or AES-256.

Luckily Gorilla provides a function to easily deal with this:

```go
HashKey = securecookie.GenerateRandomKey(64)
BlockKey = securecookie.GenerateRandomKey(32)
```

Of course you will not want to random new keys every time you run your server.


### Configuration File Attributes
There are two ways to change the values of the module, one is to to simply set them as you would any other global.
The other is to load values from a config file.  The following fields are recognized:

* cookiename - loads to the CookieName global
* prefix - loads to the Prefix global
* hashkey - loads to the HashKey global - expected to be a hexideciamal representation of a byte array.
* blockkey - loads to the BlockKey global - expected to be a hexideciamal representation of a byte array.
* expiration - loads to the Expiration global

### Methods
There are three methods you need to be aware of in order to get stuff done:

1. `ginAuth.Use()` - this method is used whenever you use the Gin's Use() when setting up routes and groups, and you 
want it to use authentication.
2. `ginAuth.Login(ctx *gin.Context, extra map[string]string)` - this method is used to log in the user.  The first
parameter is the current context, and the second is a map of strings of data to set on the cookie.  There are a few
reserved indexes: ip, hash, and expiration.  It will return an error if you try to use any of those indexes.
3. `ginAuth.Logout(ctx *gin.Context)` - simply pass in the current context and it'll remove the cookie from the user.
