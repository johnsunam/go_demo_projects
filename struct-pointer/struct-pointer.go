package main

import (
	"fmt"
)

type Node interface {
	SetValue(v int)
	GetValue() int
}
type SLLNode struct {
	next  *SLLNode
	value int
}

type SLLNode1 struct {
	next  *SLLNode1
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode SLLNode1) SetValue1(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func (sNode SLLNode1) GetValue1() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

func NewSLLNode1() SLLNode1 {
	return SLLNode1{}
}

//type power node

type PowerNode struct {
	next  *PowerNode
	value int
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value: ", node.GetValue())
	node1 := NewSLLNode1()
	node1.SetValue1(4)
	fmt.Println("Node1 is of value", node1.GetValue1())

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Printf("node is of value ", node.GetValue())
	if n, ok := node.(*PowerNode); ok {
		fmt.Println("this is a power node of value", n.value)
	}
}
