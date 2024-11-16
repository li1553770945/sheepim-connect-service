namespace go message
include "base.thrift"

struct SendMessageReq{
    1: required string clientId
    2: required string event
    3: required string type
    4: required string message
}
struct SendMessageResp {
    1: required base.BaseResp baseResp
}

service MessageService {
   SendMessageResp SendMessage(SendMessageReq req)
}
