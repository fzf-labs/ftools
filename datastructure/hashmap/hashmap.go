package hashmap

var defaultMapCapacity uint64 = 16

// 扩容因子
var expandFactor float64 = 0.75

type mapNode struct {
	key   string
	value any
	next  *mapNode
}

type HashMap struct {
	cap   uint64
	size  uint64
	table []*mapNode
}

func NewHashMap() *HashMap {
	return &HashMap{
		cap:   defaultMapCapacity,
		table: make([]*mapNode, defaultMapCapacity),
	}
}

// NewHashMapWithCapacity 返回具有给定大小和容量的 HashMap 实例
func NewHashMapWithCapacity(size, cap uint64) *HashMap {
	return &HashMap{
		size:  size,
		cap:   cap,
		table: make([]*mapNode, cap),
	}
}

func (hm *HashMap) Get(key string) any {
	hash := hm.Hash(key)
	node := hm.table[hash]
	if node != nil {
		for {
			if node.key == key {
				return node.value
			}
			if node.next == nil {
				return nil
			}
			node = node.next
		}
	}
	return nil
}

func (hm *HashMap) Hash(key string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(key); i++ {
		hash = (hash * seed) + uint64(key[i])
	}
	return hash % hm.cap
}

func (hm *HashMap) Put(key string, value any) {
	if hm.cap == 0 {
		hm.cap = defaultMapCapacity
		hm.table = make([]*mapNode, defaultMapCapacity)
	}
	hash := hm.Hash(key)
	node := hm.table[hash]
	if node == nil {
		hm.table[hash] = newMapNode(key, value)
		hm.size++
	} else {
		tmp := node
		for tmp.next != nil {
			if tmp.key == key {
				tmp.value = value
				return
			}
			tmp = tmp.next
		}
		tmp.next = newMapNode(key, value)
		hm.size++
	}
	if float64(hm.size)/float64(hm.cap) > expandFactor {
		hm.expand()
	}
}

// 扩容
func (hm *HashMap) expand() {
	c := hm.cap
	size := hm.size
	newMap := NewHashMapWithCapacity(size, 2*c)
	for _, pairs := range hm.table {
		for pairs != nil {
			newMap.Put(pairs.key, pairs.value)
		}
	}
	hm.table = newMap.table
}

func newMapNode(key string, value any) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
	}
}
