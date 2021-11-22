package msgtypes

type SetNamespaceMsg string
type GetNamespacesMsg []string
type SetKindMsg struct {
	Kind  string
	Items []string
}
type GetKindInstanceDescribeMsg string

type KindInstanceYaml string
type DumpToYamlMsg bool
