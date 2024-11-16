package project

type ConnectService struct {
}

func NewConnectService() IConnectService {
	return &ConnectService{}
}
