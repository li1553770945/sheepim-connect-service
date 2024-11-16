namespace go connect

struct ConnectResp{
    1: required i32 Code
    2: required string message
}

service ConnectService {
    ConnectResp Connect()(api.get="/connect")
}
