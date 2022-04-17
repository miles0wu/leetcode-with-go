package main

type MyCircularDeque struct {
	data []int
	head int
	tail int
	used int
}

func Constructor(k int) MyCircularDeque {
	var data = make([]int, k, k)

	return MyCircularDeque{data, 1, 0, 0}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head--
	this.alignHead()
	this.data[this.head] = value
	this.used++

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail++
	this.alignTail()
	this.data[this.tail] = value
	this.used++

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.used--
	this.head++
	this.alignHead()
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.used--
	this.tail--
	this.alignTail()
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.used == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.used >= len(this.data)
}

func (this *MyCircularDeque) alignHead() {
	this.head = (this.head + len(this.data)) % len(this.data)
}

func (this *MyCircularDeque) alignTail() {
	this.tail = (this.tail + len(this.data)) % len(this.data)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
