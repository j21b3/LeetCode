package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	一共可以分为3步
		1、在每一个原始节点的后面创建新节点，数值复制原始节点
		2、复制random，一个新节点的random等于前面节点的random.next
		3、把新节点全部提出来串成链表
*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for h1 := head; h1 != nil; h1 = h1.Next.Next {
		h1.Next = &Node{
			Val:  h1.Val,
			Next: h1.Next,
		}
	}

	for h2 := head; h2 != nil; h2 = h2.Next.Next {
		if h2.Random != nil {
			h2.Next.Random = h2.Random.Next
		}
	}

	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew
}
