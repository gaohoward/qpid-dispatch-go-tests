package qdrmanagement

import (
	"encoding/json"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/qdrmanagement/entities"
	"reflect"
	"time"
)

const (
	timeout time.Duration = 60 * time.Second
)

var (
	queryCommand = []string{"qdmanage", "query", "--type"}
)

// QdmanageQuery executes a "qdmanage query" command on the given pod
// to retrieve all entities for the given <entity> type.
func QdmanageQuery(ctxData framework.ContextData, pod string, entity string) (string, error) {
	command := append(queryCommand, entity)
	kubeExec := framework.NewKubectlExecCommand(ctxData, pod, timeout, command...)
	return kubeExec.Exec()
}

// QdmanageQueryConnections use qdmanage to query existing connections on the given pod
func QdmanageQueryConnections(ctxData framework.ContextData, pod string) ([]entities.Connection, error) {
	return QdmanageQueryConnectionsFilter(ctxData, pod, nil)
}

// QdmanageQueryConnectionsFilter use qdmanage to query existing connections on the given pod
// filtering entities using the provided filter function (if one is given)
func QdmanageQueryConnectionsFilter(ctxData framework.ContextData, pod string, filter func(entity interface{}) bool) ([]entities.Connection, error) {
	jsonString, err := QdmanageQuery(ctxData, pod, entities.Connection{}.GetEntityId())
	var connections []entities.Connection
	if err == nil {
		err = json.Unmarshal([]byte(jsonString), &connections)
		filtered := FilterEntities(connections, filter)
		connections = nil
		for _, v := range filtered {
			connections = append(connections, v.(entities.Connection))
		}
	}
	return connections, err
}

// QdmanageQueryNodes use qdmanage to query existing nodes on the given pod
func QdmanageQueryNodes(ctxData framework.ContextData, pod string) ([]entities.Node, error) {
	return QdmanageQueryNodesFilter(ctxData, pod, nil)
}

// QdmanageQueryNodesFilter use qdmanage to query existing nodes on the given pod
// filtering entities using the provided filter function (if one given)
func QdmanageQueryNodesFilter(ctxData framework.ContextData, pod string, filter func(entity interface{}) bool) ([]entities.Node, error) {

	jsonString, err := QdmanageQuery(ctxData, pod, entities.Node{}.GetEntityId())
	var nodes []entities.Node
	if err == nil {
		err = json.Unmarshal([]byte(jsonString), &nodes)
		filtered := FilterEntities(nodes, filter)
		nodes = nil
		for _, v := range filtered {
			nodes = append(nodes, v.(entities.Node))
		}
	}
	return nodes, err
}

// filter is an internal method to be invoked by specific Query<Entity> methods
// so all methods can reuse the same code for filtering entities
func FilterEntities(i interface{}, fn func(i interface{}) bool) []interface{} {
	s := reflect.ValueOf(i)
	if s.Kind() != reflect.Slice {
		panic("Expecting a slice")
	}

	var ret []interface{}
	ri := 0
	for j := 0; j < s.Len(); j++ {
		ii := s.Index(j).Interface()
		if fn == nil || fn(ii) {
			ret = append(ret, ii)
			ri++
		}
	}

	return ret
}
