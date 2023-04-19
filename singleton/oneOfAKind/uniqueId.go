package oneOfAKind

type idServant struct {
	nextId int
}

var idServantInstance *idServant

func init() {
	idServantInstance = &idServant{}
}

func GetId() int {
	defer func() {
		idServantInstance.nextId++
	}()
	return idServantInstance.nextId
}
