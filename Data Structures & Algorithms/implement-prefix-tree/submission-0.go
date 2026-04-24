type PrefixTree struct {
	prefixHash map[string]bool 
	wordHash map[string]bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		prefixHash: make(map[string]bool),
		wordHash: make(map[string]bool),
	}
}

func (this *PrefixTree) Insert(word string) {
	this.wordHash[word] = true
	var w string 
	for i := 1 ; i <= len(word); i++ {
		w = word[:i]
		this.prefixHash[w] = true
	}
}

func (this *PrefixTree) Search(word string) bool {
	return this.wordHash[word]
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	return this.prefixHash[prefix]
}
