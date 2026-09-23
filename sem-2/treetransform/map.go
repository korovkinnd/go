//go:build !solution

package treetransform

import "fmt"

func Map(root *Node, transform Transform) (*Node, error) {
	if transform == nil {
		return nil, ErrNilTransform
	}

	return mapNode(root, transform)
}

func mapNode(node *Node, transform Transform) (*Node, error) {
	if node == nil {
		return nil, nil
	}

	value, err := transform(node.Value)
	if err != nil {
		return nil, fmt.Errorf("transform node %d: %w", node.Value, err)
	}

	left, err := mapNode(node.Left, transform)
	if err != nil {
		return nil, err
	}

	right, err := mapNode(node.Right, transform)
	if err != nil {
		return nil, err
	}

	return &Node{Value: value, Left: left, Right: right}, nil
}
