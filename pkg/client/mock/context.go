package mock

type Context map[interface{}]interface{}

func (c Context) Set(key, value interface{}) {
	if c == nil {
		c = make(map[interface{}]interface{})
	}
	c[key] = value
}

func (c Context) Get(key interface{}) interface{} {
	if c == nil {
		return nil
	}
	return c[key]
}
