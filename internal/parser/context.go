package parser

type Context struct {
	Functions map[string]Function
	Variables map[string]Variable
}

func MakeContext() Context {
	return Context{
		Functions: make(map[string]Function),
		Variables: make(map[string]Variable),
	}
}

func (c *Context) AddFunction(f Function) {
	c.Functions[f.Name] = f
}

func (c *Context) GetFunction(name string) (Function, bool) {
	f, ok := c.Functions[name]
	return f, ok
}

func (c *Context) AddVariable(t Variable) {
	c.Variables[t.Name] = t
}

func (c *Context) GetVariable(name string) (Variable, bool) {
	v, ok := c.Variables[name]
	return v, ok
}
