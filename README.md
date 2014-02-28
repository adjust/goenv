#goenv

This is a nifty little package to help reduce the boilerplate code of our go apps and make them behave more like rails.

##Features
- rails style `config/config.yaml` to configure different databases. (see `example.yaml`)
- `GO_ENV` similar to `RAILS_ENV` to have development, testing and production databases
- default `log/server.log` for log output
- prefab ExitHandler to make your app handle SIGHUP correctly

##How to use
Get it

    go get github.com/adjust/goenv

###General
Import it

    import github.com/adjust/goenv

Use it

    redis_host, redis_port, redis_db := goenv.GetRedis()

###Exit Handler
To make your app quit on SIGHUP and execute some function before closing use a custom exit handler.

    type MyExitHandler struct {
    }

    func (self *MyExitHandler) OnExit() {
        log.Println("running sig handler")
        //do something like cleanup or saving progress...
    }

    func main() {
        goenv.SetExitHandler(&MyExitHandler{})
    }

##How to extend database config functions
Easy: Just fork and create a new file (good place is to start with a `postgres.go` copy) with a function returning the parameters you want.

## License

This Software is licensed under the MIT License.

Copyright (c) 2012 adjust GmbH,
http://www.adjust.com

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
