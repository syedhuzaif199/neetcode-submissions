type Node struct {
    key int
	val int
	next *Node
	prev *Node
}

type LRUCache struct {
    m map[int]*Node
    head *Node
    tail *Node
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
        m        : map[int]*Node{},
        capacity : capacity,
    }
}

func (this *LRUCache) AddAtHead(node *Node) {
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	node.next = this.head
    node.prev = nil
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) AddAtTail(node *Node) {
	if this.tail == nil {
		this.head = node
		this.tail = node
		return 
	}

	node.prev = this.tail
    node.next = nil
	this.tail.next = node
	this.tail = node
}

func (this *LRUCache) PopHead() (*Node, bool) {
    if this.head == nil {
        return nil, false
    }

    out := this.head
    this.head = this.head.next
    if this.head == nil {
        this.tail = nil
    } else {
        this.head.prev = nil
    }
    return out, true
}

func (this *LRUCache) PopTail() (*Node, bool) {
    if this.tail == nil {
        return nil, false
    }

    out := this.tail
    this.tail = this.tail.prev
    if this.tail == nil {
        this.head = nil
    } else {
        this.tail.next = nil
    }
    return out, true
}

func (this *LRUCache) Remove(node *Node) {
    if node == nil {
        return
    }
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        this.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        this.tail = node.prev
    }
}

func (this *LRUCache) Get(key int) int {
    node, exists := this.m[key]
    if !exists {
        return -1
    }
    this.Remove(node)
    this.AddAtHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, exists := this.m[key]
    if exists {
        node.val = value
        this.Remove(node)
        this.AddAtHead(node)
        return
    }
    if len(this.m) == this.capacity {
        removed, _ := this.PopTail()
        // no success checking
        delete(this.m, removed.key)
        node = removed
        node.prev = nil
        node.next = nil
    } else {
        node = new(Node)
    }

    node.key = key
    node.val = value
    this.AddAtHead(node)
    this.m[key] = node
}