package tg_bot

const (
	Start    = "/start"
	Shopping = "/shopping"
	Products = "/products"
	Recipes  = "/recipes"
	Workouts = "/workouts"
	Stop     = "/stop"
)

type operation uint8

const (
	Nothing operation = iota
	ShowAll
	Add
	Change
	Delete
)

const (
	DefaultOffset  = 0
	DefaultTimeout = 60
	WelcomeMessage = `Привет, я твой бот помощник по дому.
Я помогу тебе вести учет продуктов, список необходимых покупок, рецептов и тренировок.`
	FarewellMessage = `Всего доброго, жду тебя снова!`
)

const (
	ShoppingBanner = `
Это режим управления покупками!
Здесь ты можешь вести список вещей, которые тебе необходимо купить.
В данном режиме можно:
1. Посмотреть список.
2. Добавить покупку.
3. Исправить покупку.
4. Удалить покупку.
Чтобы выбрать действие, просто пришли мне его номер! 😊
`
	ProductsBanner = `
Это режим управления продуктами! 
Здесь ты можешь вести список продуктов и их количество, которые остались у тебя дома.
В данном режиме можно:
1. Посмотреть список.
2. Добавить продукт.
3. Изменить данные продукта.
4. Удалить продукт.
Чтобы выбрать действие, просто пришли мне его номер! 😊
`
	RecipesBanner = `
Это режим управления рецептами!
Здесь ты можешь вести список своих самых лучших рецептов.
В данном режиме можно:
1. Посмотреть рецепты.
2. Добавить.
3. Изменить.
4. Удалить рецепт.
Чтобы выбрать действие, просто пришли мне его номер! 😊
`
	WorkoutsBanner = `
Это режим управления тренировками!
Если ты занимаешься спортом и оплачиваешь тренировки, то я помогу тебе с расписанием тренировок, а так же подскажу, когда их необходимо продлевать.
В данном режиме можно:
1. Посмотреть тренировку.
2. Добавить тренировку.
3. Изменить.
4. Удалить.
Чтобы выбрать действие, просто пришли мне его номер! 😊
`
	AboutDisable = `Я выключен, чтобы воспользоваться этой функцией сначала запусти меня!`
)
