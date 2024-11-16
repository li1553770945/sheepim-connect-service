package main

import (
	"context"
	project "github.com/li1553770945/sheepim-connect-service/kitex_gen/project"
)

// ProjectServiceImpl implements the last service interface defined in the IDL.
type ProjectServiceImpl struct{}

// SendMessage implements the ProjectServiceImpl interface.
func (s *ProjectServiceImpl) SendMessage(ctx context.Context, req *project.SendMessageReq) (resp *project.SendMessageResp, err error) {
	// TODO: Your code here...
	return
}
