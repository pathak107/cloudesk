package graphql

import "github.com/pathak107/cloudesk/pkg/cloud"

func makeGraphQLEuery(query string) {

}

func AddVmInfo(vm *cloud.Instance) {
	makeGraphQLEuery("some query mutation")
}

func ChangeStatusOfVM(vmID string, status string) {
	makeGraphQLEuery("some query mutation")
}
