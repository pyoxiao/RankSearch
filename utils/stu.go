package utils

type Stu struct {
	Name 		string
	UserName    string
	HduRank 	int
	HduSol 		int
}

type Stus []Stu

func(stus Stus) Len() int {
	return len(stus)
}

func(stus Stus) Swap (i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}

func(stus Stus) Less (i, j int) bool {
	return stus[j].HduRank < stus[i].HduRank
}