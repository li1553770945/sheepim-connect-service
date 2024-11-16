namespace go project
include "base.thrift"

struct SendMessageReq{
    1: required string userId
    2: required string event
    3: required string type
    4: required string message
}
struct SendMessageResp {
    1: required base.BaseResp baseResp
}

service ProjectService {
   SendMessageResp SendMessage(SendMessageReq req)
}
