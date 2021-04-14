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
	Cmd      string
	Children []PID
}

func NewPid(pid int) (*PID, error) {
	tree, err := pstree.New()
	if err != nil {
		return nil, err
	}

	p := &PID{Id: pid}
	if pid > 0 {
		p.Cmd = tree.Procs[pid].Name
		err = getChildren(p, tree)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

func newPid(pid int, tree *pstree.Tree) (*PID, error) {
	p := &PID{Id: pid}
	if pid > 0 {
		p.Cmd = tree.Procs[pid].Name
		err := getChildren(p, tree)
		if err != nil {
			return p, err
		}
	}
	return p, nil
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
	if p.Id == 0 {
		return "process is not running"
	}
	var buf bytes.Buffer
	getString(p, 0, &buf)
	return buf.String()
}

func getString(p *PID, indent int, buf io.Writer) {
	str := strings.Repeat("  ", indent)
	_, _ = buf.Write([]byte(fmt.Sprintf("%s%d %s\n", str, p.Id, p.Cmd)))
	for _, c := range p.Children {
		getString(&c, indent+1, buf)
	}
}
