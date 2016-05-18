type listCloner interface {
	List(user string) []github.Repo
	Clone(repo github.Repo)
}

type ReposHandlers struct {
	Manager listCloner
}
