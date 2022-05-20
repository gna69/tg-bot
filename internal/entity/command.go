package entity

const (
	Start    = "/start"
	Shopping = "/shopping"
	Products = "/products"
	Recipes  = "/recipes"
	Workouts = "/workouts"
	Stop     = "/stop"
)

type Command struct {
	command  string
	action   Action
	objectId uint
	Purchase *Purchase
	Product  *Product
	Recipe   *Recipe
	Workout  *Workout
}

func NewCommand() *Command {
	return &Command{
		command: Stop,
		action:  Nothing,
	}
}

func (c *Command) GetCommand() string {
	return c.command
}

func (c *Command) SetCommand(command string) {
	c.command = command
}

func (c *Command) GetAction() Action {
	return c.action
}

func (c *Command) SetAction(action Action) {
	c.action = action
}

func (c *Command) GetObjectId() uint {
	return c.objectId
}

func (c *Command) SetObjectId(id uint) {
	c.objectId = id
}
