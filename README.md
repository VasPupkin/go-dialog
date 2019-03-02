go-dialog
=========

go-dialog is a Go wrapper for the dialog utility originally written by Savio Lam, and later rewritten by Thomas E. Dickey.

Usage
=========
```go
package main

import (
	"fmt"
	"github.com/VasPupkin/go-dialog"	
)
func main() {
   d := dialog.New(dialog.AUTO, 0)
   d.Msgbox("Hello world!")
}
```

Installation
=========
```bash
 go get "github.com/VasPupkin/go-dialog"
```

Contributors
=========
* [Valeriy Soloviov](http://github.com/weldpua2008/) weldpua2008 @gmail.com
* [Dmitry Orzhehovsky](http://github.com/dorzheh/) dorzheh @gmail.com
* Pavel Vershinin master-dev @inbox.ru
