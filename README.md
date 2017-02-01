## Go-EZmail: quick and easy email sending with Go

Usage of this small and simple package should be pretty self explanatory. However, feel free to see the [example section below](#example) for usage with `package main`. Also, I'll be working on a more thorough explanation on my site and (possibly) adding this package to GoDoc.

There are two main things you have to know before you use this package:

1. Go. 
2. There are actually two packages. 
	* The most useful of which is the main one, `go-ezmail.go`, which can be imported with `import "github.com/benobro/go-ezmail"` (the package can be called with `ezmail`)
	* The second one adds an option for a code (the variable is called `yourcode`). This is useful if you'll be using the email to send a verification link/code, for example. It can be imported with `import "github.com/benobro/go-ezmail/with-code"` and called with `ezmailcode`. 


For help with `net/smtp`, the driving force behind this package, see [the official Go documentation](https://golang.org/pkg/net/smtp).

## Example:

This example is for the `ezmail` package.

Remember, if you're using Gmail, you'll need to enable [less secure apps access](https://support.google.com/accounts/answer/6010255?hl=en).
```
package main

import (
	"github.com/benobro/go-ezmail"
)

func main(){ 
	ezmail.SendMail("SEND@email.com", "YOUR-SECRET-PASSWORD", "smtp.gmail.com", 587, "RECEIVE@email.com", "test", "test ")
}

```

Here's an example of `ezmailcode`. The code I generate here is a 50 character string, containing a-z lowercase, A-Z uppercase, and 0-9. That string is then sent along with the `message` in `ezmailcode.SendMail`

```
package main

import (
	"github.com/benobro/go-ezmail/with-code"
	"math/rand"
	"time"
)

func main(){ 
	yourcode := RandomString(50)
	ezmailcode.SendMail("SEND@email.com", "YOUR-SECRET-PASSWORD", "smtp.gmail.com", 587, "RECEIVE@email.com", "test", "test ", yourcode)
}

func RandomString(strlen int) string {
        rand.Seed(time.Now().UTC().UnixNano())
        const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
        result := make([]byte, strlen)
        for i := 0; i < strlen; i++ {
                result[i] = chars[rand.Intn(len(chars))]
        }
        return string(result)
}
```

_Go-EZmail is written by [Ben O'Brien](https://benobro.io)._ 

