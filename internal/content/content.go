package content

type Content struct{}
type ContentManager interface {
	Get()
	GetStar()
	Put()
	PutStar()
}

func (c Content) Get()      {}
func (c *Content) GetStar() {}
func (c Content) Put()      {}
func (c *Content) PutStar() {}

type DataLayer struct {
	Manager ContentManager
}
