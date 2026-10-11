package t_017_707_design_linked_list

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// MyLinkedList 链表
type MyLinkedList struct {
	Node *ListNode
}

// Constructor 创建单链表
func Constructor() MyLinkedList {
	// 定义单链表头节点
	virHead := ListNode{Next: nil}
	// 创建单链表 MyLinkedList 首次创建链表为空 但具有头节点 返回
	linkedList := MyLinkedList{Node: &virHead}
	return linkedList
}

// Get get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (my *MyLinkedList) Get(index int) int {
	virHead := my.Node
	for i := 0; i <= index; i++ {
		if virHead.Next == nil {
			return -1
		}
		virHead = virHead.Next
	}
	return virHead.Val
}

// AddAtHead addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (my *MyLinkedList) AddAtHead(val int) {
	virHead := my.Node
	newNode := ListNode{Val: val, Next: virHead.Next}
	virHead.Next = &newNode
}

// AddAtTail addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
func (my *MyLinkedList) AddAtTail(val int) {
	virHead := my.Node
	newNode := ListNode{Val: val, Next: nil}
	for virHead.Next != nil {
		virHead = virHead.Next
	}
	virHead.Next = &newNode
}

// AddAtIndex addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (my *MyLinkedList) AddAtIndex(index int, val int) {
	virHead := my.Node
	if index <= 0 {
		my.AddAtHead(val)
		return
	}
	newNode := ListNode{Val: val, Next: nil}
	for i := 0; ; i++ {
		if i == index {
			//my.AddAtTail(val)
			newNode.Next = virHead.Next
			virHead.Next = &newNode
			break
		}
		if virHead.Next == nil {
			break
		}
		virHead = virHead.Next
	}
}

// DeleteAtIndex deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
func (my *MyLinkedList) DeleteAtIndex(index int) {
	virHead := my.Node
	for i := 0; i <= index; i++ {
		if i == index {
			if virHead.Next != nil {
				virHead.Next = virHead.Next.Next
			} else {
				virHead.Next = nil
			}
			break
		}
		if virHead.Next == nil {
			break
		}
		virHead = virHead.Next
	}
}

// GetMyLinkedListLen 获取链表长度
func (my *MyLinkedList) GetMyLinkedListLen() int {
	virHead := my.Node
	listLen := 0
	for ; virHead.Next != nil; listLen++ {
		virHead = virHead.Next
	}
	return listLen
}

// PrintMyLinkedList 将链表转化为数组输出
func (my *MyLinkedList) PrintMyLinkedList() []int {
	virHead := my.Node
	listLen := my.GetMyLinkedListLen()
	myList := make([]int, 0, listLen)
	for virHead.Next != nil {
		virHead = virHead.Next
		//fmt.Println(virHead.Val)
		myList = append(myList, virHead.Val)
	}
	fmt.Println("MyLinkedList:", myList)
	return myList
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
