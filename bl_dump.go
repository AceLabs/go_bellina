package bl

import "fmt"

func Dump() {
	dump(root, "    ")
}

func dump(n *Node, tab string) {
	fmt.Println(tab, n.Name)

	for kid := n.Kids.Front(); kid != nil; kid = kid.Next() {
		dump(kid.Value.(*Node), tab + "    " )
	}
}



