package engine

import (
	"fmt"

	api "github.com/awe76/saga/api/sagatransactionapis/v1"
)

type route = map[string][]*api.Operation
type getOperationKey = func(op *api.Operation) string

func getFrom(op *api.Operation) string {
	return op.From
}

func getTo(op *api.Operation) string {
	return op.To
}

func getKey(op *api.Operation, isRollback bool) string {
	return fmt.Sprintf("%s:%s:%s:%v", op.Name, op.From, op.To, isRollback)
}

func hasOp(m map[string]*api.Operation, op *api.Operation, isRollback bool) bool {
	key := getKey(op, isRollback)
	_, found := m[key]
	return found
}

func addOp(m map[string]*api.Operation, op *api.Operation, isRollback bool) {
	key := getKey(op, isRollback)
	m[key] = op
}

func removeOp(m map[string]*api.Operation, op *api.Operation, isRollback bool) {
	key := getKey(op, isRollback)
	delete(m, key)
}

func allMatched(current string, relation route, isMatched func(op *api.Operation) bool) bool {
	// get all related operations
	if ops, found := relation[current]; found {
		// for each operation
		for _, op := range ops {
			if !isMatched(op) {
				return false
			}
		}
	}

	return true
}

func createRoute(operations []*api.Operation, getKey getOperationKey) route {
	result := make(map[string][]*api.Operation)
	for _, op := range operations {
		key := getKey(op)
		if next, found := result[key]; found {
			result[key] = append(next, op)
		} else {
			result[key] = append([]*api.Operation{}, op)
		}
	}

	return result
}

type tracer struct {
	isReady        func(current string) bool
	isFinished     func(current string) bool
	getNext        func(current string) (ops []*api.Operation, found bool)
	isProcessed    func(op *api.Operation) bool
	canBeSpawned   func(op *api.Operation) bool
	getNextVertex  func(op *api.Operation) string
	endWorkflow    func() error
	spawnOperation func(op *api.Operation) error
}

func CreateDirectTracer(
	w *api.Workflow,
	s *api.State,
	endWorkflow func() error,
	spawnOperation func(op *api.Operation) error,
) *tracer {
	from := createRoute(w.Operations, getFrom)
	to := createRoute(w.Operations, getTo)

	isMatched := func(op *api.Operation) bool {
		return hasOp(s.Done, op, false)
	}

	return &tracer{
		isReady: func(current string) bool {
			return allMatched(current, to, isMatched)
		},
		isFinished: func(current string) bool {
			return current == w.End
		},
		getNext: func(current string) ([]*api.Operation, bool) {
			ops, found := from[current]
			return ops, found
		},
		isProcessed: func(op *api.Operation) bool {
			return hasOp(s.Done, op, false)
		},
		canBeSpawned: func(op *api.Operation) bool {
			return !hasOp(s.InProgress, op, false)
		},
		getNextVertex: func(op *api.Operation) string {
			return op.To
		},
		endWorkflow: endWorkflow,
		spawnOperation: func(op *api.Operation) error {
			return spawnOperation(op)
		},
	}
}

func CreateReverseTracer(
	w *api.Workflow,
	s *api.State,
	endWorkflow func() error,
	spawnOperation func(op *api.Operation) error,
) *tracer {
	from := createRoute(w.Operations, getFrom)
	to := createRoute(w.Operations, getTo)

	isMatched := func(op *api.Operation) bool {
		done := hasOp(s.Done, op, false)
		return !hasOp(s.InProgress, op, false) && !hasOp(s.InProgress, op, true) && (!done || (done && hasOp(s.Done, op, true)))
	}

	return &tracer{
		isReady: func(current string) bool {
			return allMatched(current, from, isMatched)
		},
		isFinished: func(current string) bool {
			return current == w.Start
		},
		getNext: func(current string) ([]*api.Operation, bool) {
			ops, found := to[current]
			return ops, found
		},
		isProcessed: func(op *api.Operation) bool {
			done := hasOp(s.Done, op, false)
			return !done || (done && hasOp(s.Done, op, true))
		},
		canBeSpawned: func(op *api.Operation) bool {
			return !hasOp(s.InProgress, op, true)
		},
		getNextVertex: func(op *api.Operation) string {
			return op.From
		},
		endWorkflow: endWorkflow,
		spawnOperation: func(op *api.Operation) error {
			return spawnOperation(op)
		},
	}
}

func (t *tracer) ResolveWorkflow(current string) error {
	// current vertex is ready for resolution
	if t.isReady(current) {
		if t.isFinished(current) {
			err := t.endWorkflow()
			if err != nil {
				return err
			}
		} else {
			// get next operations
			if ops, found := t.getNext(current); found {
				// for each next operation
				for _, op := range ops {
					if t.isProcessed(op) {
						// if operation has been already processed continue resolution of the next vertex
						nextVertex := t.getNextVertex(op)
						err := t.ResolveWorkflow(nextVertex)
						if err != nil {
							return err
						}
					} else if t.canBeSpawned(op) {
						// if operation can be spawned do it
						err := t.spawnOperation(op)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
