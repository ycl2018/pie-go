# Pie-go
 
The go implementation of Pie language interpreter by antlr4.

Go语言版Pie语言解释器，基于antlr4实现。

```go
package compile

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/pie-go/interpreter/vm"
)

func TestPieInterpreter_Interp(t *testing.T) {
	tests := []struct {
		name    string
		program string
	}{
		{
			name: "apple.pie",
			program: `
i = 0
while i<10:
	print i*3.2
	i = i + 1
	if i<5 print i + " is less than 5"
	else print "foo"
.
	`},
		{
			name: "cheery.pie",
			program: `
def f(x) return 2*x
print f(4)
`,
		},
		{
			name: "factorials.pie",
			program: `
def fact(n):
	if n < 2 return 1
	return n * fact(n-1)
.

print fact(10)
`,
		},
		{
			name: "forward.pie",
			program: `
print f(4)           # references definition on next line
def f(x) return 2*x
print new User       # references definition on next line
struct User { name, password }
`,
		},
		{
			name: "localstruct.pie",
			program: `
struct User { name, password } # define global struct
def f():                       # define f
    struct User { x, y }       # hides global User def
    u = new User               # create new User instance, put in local u
    print u                    # prints "{x=null, y=null}"
.                              # end body of f
print new User                 # prints "{name=null, password=null}"
f()                            # call f
`,
		},
		{
			name: "lookup.pie",
			program: `
x = 1           # create global variable
def f(x):       # define f in global space
    print x     # access parameter; prints 10
    y = 2       # create local variable
.
def g():        # define g in global space
    x = 3       # set global variable
.
f(10)
g()
print x         # prints 3 (g alters global value)
`,
		},
		{
			name: "loop.pie",
			program: `
n = 100
i = 0
while i<n:
        i = i + 1
.
print "looped "+n+" times."
`,
		},
		{
			name: "struct.pie",
			program: `
struct User { name, password }
u = new User
u.name = "parrt"
print "Login: "+u.name
print u
`,
		},
		{
			name: "structerr.pie",
			program: `
struct User { name, password }
u = new User
u.name = "parrt"    # make u.name a string
u.name.y = "parrt"  # u.name is a string not a struct
u.x = 3             # x isn't a field of User; can't write to it
print u.x           # check for unknown field in expr as well
`,
		},
		{
			name: "nested.pie",
			program: `
struct User { name, addr }
struct Address { street, city, state, zip }
u = new User
u.name = "parrt"
addr = new Address
addr.street = "123 Main St"
u.addr =  addr
u.addr.city = "Chicago"
print u.addr.street
print u.addr.city
print u.addr
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPieCompiler()
			code, err := p.Compile(antlr.NewInputStream(tt.program))
			if err != nil {
				t.Fatal(err)
			}
			t.Log(p.Dump())
			interpreter := vm.NewInterpreter()
			err = interpreter.Run(code)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
```