//Package binarytree binary tree common struct and functions
package binarytree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	type args struct {
		nodes []*Node
	}
	a := &Node{
		Val: "a",
	}
	b := &Node{
		Val: "b",
	}
	c := &Node{
		Val: "c",
	}
	d := &Node{
		Val: "d",
	}
	tests := []struct {
		name        string
		args        args
		wantRoot    *Node
		wantErr     bool
		detailCheck bool
	}{
		{name: "empty tree", args: args{}, wantRoot: nil},
		{name: "simple one node", args: args{nodes: []*Node{a}}, wantRoot: a},
		{name: "abcd tree", args: args{[]*Node{a, b, nil, c, nil, d}}, wantRoot: a},
		{name: "invalid tree nodes", args: args{[]*Node{a, nil, nil, b}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoot, err := BinaryTree(tt.args.nodes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("BinaryTree() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
		cleanNodesConnection(a, b, c, d)
	}
}
