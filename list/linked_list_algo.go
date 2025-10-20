// Copyright 2021 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

// Reverse 将会创建一个新的链表，其中的元素顺序与 src 相反。
// 不会修改传入的 src。
func Reverse[T any](src *LinkedList[T]) *LinkedList[T] {
	if src == nil || src.length == 0 {
		return NewLinkedList[T]()
	}
	res := NewLinkedList[T]()
	// 遍历 src，并在 res 的头部插入，时间复杂度 O(n)
	for cur := src.head.next; cur != src.tail; cur = cur.next {
		_ = res.Add(0, cur.val)
	}
	return res
}

// ReverseSelf 会直接在 src 上进行原地反转，时间复杂度 O(n)，空间复杂度 O(1)。
// 注意：该操作会改变链表中节点的链接顺序。
func ReverseSelf[T any](src *LinkedList[T]) {
	if src == nil || src.length <= 1 {
		return
	}
	// 对环上的每个节点（包括哨兵头尾）交换 next/prev 指针
	cur := src.head
	for {
		cur.prev, cur.next = cur.next, cur.prev
		cur = cur.prev // 由于已经交换，此处沿着原来的 next 前进
		if cur == src.head {
			break
		}
	}
	// 交换 head 和 tail 的引用，维持哨兵语义
	src.head, src.tail = src.tail, src.head
}

// SwapPairs 以相邻为对的方式交换链表中的元素。
// 例如 [1,2,3,4,5] 经过 SwapPairs 之后为 [2,1,4,3,5]。
// 该实现仅交换节点中的值，不变更节点链接，时间复杂度 O(n)。
func SwapPairs[T any](src *LinkedList[T]) {
	if src == nil || src.length <= 1 {
		return
	}
	for cur := src.head.next; cur != src.tail && cur.next != src.tail; cur = cur.next.next {
		// cur 和 cur.next 都是数据节点
		cur.val, cur.next.val = cur.next.val, cur.val
	}
}

// RotateLeft 将链表循环左移 k 位。
// 当 k 为 0 或者 k 是长度的整数倍时，该方法不做任何操作。
func RotateLeft[T any](src *LinkedList[T], k int) {
	if src == nil || src.length == 0 {
		return
	}
	s := k % src.length
	if s == 0 {
		return
	}
	// A 为原链表第一个数据节点，B 为第 s 个节点（下标从 0 开始），D 为第 s-1 个节点，C 为最后一个数据节点
	A := src.head.next
	B := src.findNode(s)
	D := B.prev
	C := src.tail.prev
	// 断开 D -> B，将 A..D 移动到末尾：D.next -> tail，tail.prev -> D
	D.next = src.tail
	src.tail.prev = D
	// 将 C 连接到 A：C.next -> A，A.prev -> C
	C.next = A
	A.prev = C
	// 将 head 连接到 B：head.next -> B，B.prev -> head
	src.head.next = B
	B.prev = src.head
}

// RotateRight 将链表循环右移 k 位。
// 当 k 为 0 或者 k 是长度的整数倍时，该方法不做任何操作。
func RotateRight[T any](src *LinkedList[T], k int) {
	if src == nil || src.length == 0 {
		return
	}
	s := k % src.length
	if s == 0 {
		return
	}
	// 右移 s 位等价于左移 len-s 位
	RotateLeft(src, src.length-s)
}
