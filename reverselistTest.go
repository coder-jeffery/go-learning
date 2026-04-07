package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverList(head *ListNode) *ListNode{

	var prev *ListNode
	curr := head

	for curr != nil{
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func reverListRecursive(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}


func printList(head *ListNode){
	curr := head
	for curr != nil{
		//print
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println(head)
	printList(head)

	reversed_result := reverList(head)
	printList(reversed_result)

}