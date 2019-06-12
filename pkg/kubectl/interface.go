package kubectl

type Pod struct {
	PodName string
	ContainerName string
	Namespace     string
}

type Node struct {
	NodeName string
}



type describe interface {
	Describe()(string ,error)
}


type get interface {
	Get()(error)
}


type exec interface {
	Exec(cmd []string)(error)
}