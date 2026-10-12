// lru 算法
package main

import "fmt"

// 双向链表 Node 节点
type Node struct {
	Key   string
	Value string
	pre   *Node
	next  *Node
}

// 初始化
func (n *Node) Init(key string, value string) {
	n.Key = key
	n.Value = value
}

// 头节点位置
var head *Node

// 尾节点位置
var end *Node

// 内存大小
var limit int

// 内存结构体 大小
type LRUCache struct {
	limit   int
	HashMap map[string]*Node
}

// 获取一个内存
func GetLRUCache(limit int) *LRUCache {
	lruCache := LRUCache{limit: limit}
	lruCache.HashMap = make(map[string]*Node, limit)
	return &lruCache
}

// 获取内存值
func (l *LRUCache) get(key string) string {
	if v, ok := l.HashMap[key]; ok {
		// 刷新链表节点顺序
		l.refreshNode(v)
		return v.Value
	} else {
		return ""
	}
}

// 添加一个内存值
func (l *LRUCache) put(key, value string) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

// 刷新链表 将链表中该 node 节点删除 删除之后添加到链表尾部
func (l *LRUCache) refreshNode(node *Node) {
	if node == end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

// 删除链表节点
func (l *LRUCache) removeNode(node *Node) string {
	if node == end {
		end = end.pre
	} else if node == head {
		head = head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

// 添加链表节点
func (l *LRUCache) addNode(node *Node) {
	if end != nil {
		end.next = node
		node.pre = end
		node.next = nil
	}
	end = node
	if head == nil {
		head = node
	}
}

func main() {
	lruCache := GetLRUCache(5)
	lruCache.put("001", "用户信息1")
	lruCache.put("002", "用户信息2")
	lruCache.put("003", "用户信息3")
	lruCache.put("004", "用户信息4")
	lruCache.put("005", "用户信息5")
	lruCache.get("002")
	lruCache.put("004", "更新用户信息4")
	lruCache.put("006", "更新用户信息6")
	fmt.Println(lruCache.get("001"))
	fmt.Println(lruCache.get("006"))
	fmt.Println(lruCache.HashMap)
}
