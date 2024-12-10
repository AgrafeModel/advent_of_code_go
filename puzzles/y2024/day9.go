package y2024

import (
	"fmt"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type part struct {
	value int
	empty bool
}

type file struct {
	size  int
	empty bool
	value int
}

func Day9Part1() int {
	res := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (9)), func(line string) {
		id := 0
		var rs []part

		for i, n := range line {
			number := utils.ParseInt(n)
			if i%2 == 0 {
				for range number {
					rs = append(rs, part{
						value: id,
						empty: false,
					})
				}
				id++
			} else {
				for _ = range number {
					rs = append(rs, part{
						value: 0,
						empty: true,
					})
				}
			}
		}

		for i := 0; i < len(rs); i++ {
			n := rs[i]
			if n.empty {
				//get the last element
				last := rs[len(rs)-1]
				rs = rs[:len(rs)-1]
				rs[i] = last

				// remove last empty part
				for rs[len(rs)-1].empty {
					rs = rs[:len(rs)-1]
				}
			}
		}

		for i, n := range rs {
			res += i * n.value
		}

		fmt.Println(rs)
	})

	return res
}

func dbug(rs []file) {
	//debug
	for _, f := range rs {
		if f.empty {
			fmt.Printf("%s", strings.Repeat(".", f.size))
		} else {
			fmt.Printf("%s", strings.Repeat(fmt.Sprintf("%d", f.value), f.size))
		}
	}
	fmt.Println("")
}

func Day9Part2() int {
	res := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (9)), func(line string) {
		id := 0
		var rs []file
		fmt.Println(line)
		for i, n := range line {
			number := utils.ParseInt(n)
			if i%2 == 0 { // file part
				rs = append(rs, file{
					value: id,
					empty: false,
					size:  number,
				})
				id++
			} else { // empty part
				rs = append(rs, file{
					value: 0,
					empty: true,
					size:  number,
				})
			}
		}
		for x := 0; x < len(rs); x++ {
			v := rs[x]
			if v.empty {
				//switch with an elemenet that fit from the end
				for i := len(rs) - 1; i > x; i-- {

					if !rs[i].empty && rs[i].size <= v.size {

						size := rs[i].size
						diff := v.size - size

						rs[x] = rs[i]
						rs[i].empty = true
						if diff > 0 {
							rs = utils.InsertAfter(rs, x, file{
								value: 0,
								empty: true,
								size:  diff,
							})
						}

						break
					}

					//remove the last element if empty
					// for rs[len(rs)-1].empty {
					// 	rs = rs[:len(rs)-1]
					// }
				}
			}
		}

		//checksum
		vi := 0
		for _, v := range rs {
			if v.empty {
				vi += v.size
				continue
			}
			for range v.size {
				res += v.value * vi
				vi++
			}
		}

	})

	return res
}
