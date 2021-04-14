package service

import (
	"bytes"
	"fmt"
	"github.com/sbinet/pstree"
	"io"
	"strings"
)

type PID struct {
	Id       int
	Children []PID
}

func NewPid(pid int) (*PID, error) {
	tree, err := pstree.New()
	if err != nil {
		return nil, err
	}
	p := &PID{Id: pid}
	err = getChildren(p, tree)
	if err != nil {
		return p, err
	}
	return p, nil
}

func newPid(pid int, tree *pstree.Tree) (*PID, error) {
	p := &PID{Id: pid}
	err := getChildren(p, tree)
	return nil, err
}

func getChildren(pid *PID, tree *pstree.Tree) error {
	for _, cid := range tree.Procs[pid.Id].Children {
		child, err := newPid(cid, tree)
		if err != nil {
			return err
		}
		pid.Children = append(pid.Children, *child)
	}
	return nil
}

func (p *PID) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("PID: %d\n", p.Id))
	getString(p, 1, &buf)
	return buf.String()
}

func getString(p *PID, indent int, buf io.StringWriter) {
	str := strings.Repeat("  ", indent)
	for _, c := range p.Children {
		_, _ = buf.WriteString(fmt.Sprintf("%s%#v\n", str, c))
		getString(&c, indent+1, buf)
	}
}
