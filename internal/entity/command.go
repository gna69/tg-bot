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
	command         string
	action          Action
	workingObjectId uint
	WorkingObject   *CommandObject
}

type CommandObject struct {
	Purchase *Purchase
	Product  *Product
	Recipe   *Recipe
	Workout  *Workout
}

func NewCommand() *Command {
	return &Command{
		command: Stop,
		action:  Nothing,
		WorkingObject: &CommandObject{
			Purchase: &Purchase{},
			Product:  &Product{},
			Recipe:   &Recipe{},
			Workout:  &Workout{},
		},
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

func (c *Command) GetWorkingObjectId() uint {
	return c.workingObjectId
}

func (c *Command) SetWorkingObjectId(id uint) {
	c.workingObjectId = id
}

func (c *Command) SetObjectValue(step Step, value string) error {
	switch c.command {
	case Shopping:
		return c.WorkingObject.Purchase.SetValue(step, value)
	case Products:
		return c.WorkingObject.Product.SetValue(step, value)
	case Recipes:
		return c.WorkingObject.Recipe.SetValue(step, value)
	case Workouts:
	}
	return nil
}
