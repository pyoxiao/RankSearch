package utils

const (
	strRank = "<tr><td>Rank</td><td align=center>"
	strSol = "<tr><td>Problems Solved</td><td align=center>"
)

type strType int

const (
	rank strType = 0
	sol  strType = 1
)
var	smap = map[strType]string{
	rank: strRank,
	sol:  strSol,
}

