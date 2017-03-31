package spawning

type Pool interface {
	Add(string) Pool
	Run() []*Result
}
