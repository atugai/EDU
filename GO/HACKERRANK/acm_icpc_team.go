/*
  
[hackerrank.com] ACM ICPC Team
https://www.hackerrank.com/challenges/acm-icpc-team

*/

package hackerrank

import (
	"sort"
)

// AcmTeam calculates maximum number of teams can cover maximum number of topics.
func AcmTeam(topic []string) []int32 {
	ts := [][]int{}
	for _, v := range topic {
		tempRes := []int{}
		for _, d := range v {
			if d == '0' {
				tempRes = append(tempRes, 0)
			} else {
				tempRes = append(tempRes, 1)
			}
		}
		ts = append(ts, tempRes)
	}

	resMap := map[int32]int32{}
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			var teamScore int32 = 0
			for k := range ts[i] {
				if ts[i][k] == 1 || ts[j][k] == 1 {
					teamScore++
				}
			}

			if _, ok := resMap[teamScore]; !ok {
				resMap[teamScore] = 1
			} else {
				resMap[teamScore]++
			}
		}
	}

	scores := []int{}
	for k := range resMap {
		scores = append(scores, int(k))
	}

	sort.Ints(scores)
	maxScore := int32(scores[len(scores)-1])
	return []int32{maxScore, resMap[maxScore]}
}
