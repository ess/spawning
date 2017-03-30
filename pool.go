package spawning

type Pool interface {
	Add(string)
	Run() []*Result
}
