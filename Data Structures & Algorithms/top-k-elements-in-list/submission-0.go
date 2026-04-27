
type Item struct {
	e int
	c int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]Item)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			mt := m[num]
			mt.c++
			m[num] = mt
			continue
		}
		m[num] = Item{
			e: num,
			c: 1,
		}
	}
	var items []Item
	for _, v := range m {
		items = append(items, v)
	}
	sort.Slice(items, func(i, j int) bool {
		it := items[i]
		jt := items[j]
		return it.c > jt.c
	})
	res := []int{}
	for i := 0; i < k; i++ {
		it := items[i]
		res = append(res, it.e)
	}
	return res
}