type MyHashSet struct {
	set map[int]any
}

func Constructor() MyHashSet {
	return MyHashSet{
		set: make(map[int]any),
	}
}

func (this *MyHashSet) Add(key int) {
    this.set[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
    delete(this.set,key)
}

func (this *MyHashSet) Contains(key int) bool {
    _ , ok := this.set[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 