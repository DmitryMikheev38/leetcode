package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val: 0,
															Next: &ListNode{
																Val: 0,
																Next: &ListNode{
																	Val: 0,
																	Next: &ListNode{
																		Val: 0,
																		Next: &ListNode{
																			Val: 0,
																			Next: &ListNode{
																				Val: 0,
																				Next: &ListNode{
																					Val: 0,
																					Next: &ListNode{
																						Val: 0,
																						Next: &ListNode{
																							Val: 0,
																							Next: &ListNode{
																								Val: 0,
																								Next: &ListNode{
																									Val: 0,
																									Next: &ListNode{
																										Val: 0,
																										Next: &ListNode{
																											Val: 0,
																											Next: &ListNode{
																												Val: 0,
																												Next: &ListNode{
																													Val: 0,
																													Next: &ListNode{
																														Val: 0,
																														Next: &ListNode{
																															Val: 0,
																															Next: &ListNode{
																																Val: 0,
																																Next: &ListNode{
																																	Val: 0,
																																	Next: &ListNode{
																																		Val:  1,
																																		Next: nil,
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := &ListNode{}
	cur := resultList

	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		res := v1 + v2 + carry
		carry = res / 10
		cur.Next = &ListNode{
			Val: res % 10,
		}

		cur = cur.Next

	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return resultList.Next
}