package entity

const (
	Start    = "/start"
	Groups   = "/groups"
	Shopping = "/shopping"
	Products = "/products"
	Recipes  = "/recipes"
	Workouts = "/workouts"
)

type Command struct {
	command         string
	action          Action
	currentUser     uint
	Object          Object
	workingObjectId uint
}

func NewCommand() *Command {
	return &Command{
		command: Start,
		action:  Nothing,
	}
}

func (c *Command) GetCommand() string {
	return c.command
}

func (c *Command) SetCommand(command string) {
	c.command = command
}

func (c *Command) GetCurrentUser() uint {
	return c.currentUser
}

func (c *Command) SetCurrentUser(id uint) {
	c.currentUser = id
}

func (c *Command) GetAction() Action {
	return c.action
}

func (c *Command) SetAction(action Action) {
	c.action = action
}

func (c *Command) GetWorkingObjectId() uint {
	return c.workingObjectId
}

func (c *Command) SetWorkingObjectId(id uint) {
	c.workingObjectId = id
}
