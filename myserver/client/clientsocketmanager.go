package client

var gClientMgr *ClientManager

func init() {
	gClientMgr = &ClientManager{}
}

func GetClientMgr() *ClientManager {
	return gClientMgr
}

type ClientManager struct {
}

func (this *ClientManager) AddSocket(pSock *ClientSocket) {

}
