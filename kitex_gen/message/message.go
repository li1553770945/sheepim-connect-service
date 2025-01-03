// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package message

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/base"
	"strings"
)

type SendMessageReq struct {
	ClientId string `thrift:"clientId,1,required" frugal:"1,required,string" json:"clientId"`
	Event    string `thrift:"event,2,required" frugal:"2,required,string" json:"event"`
	Type     string `thrift:"type,3,required" frugal:"3,required,string" json:"type"`
	Message  string `thrift:"message,4,required" frugal:"4,required,string" json:"message"`
}

func NewSendMessageReq() *SendMessageReq {
	return &SendMessageReq{}
}

func (p *SendMessageReq) InitDefault() {
	*p = SendMessageReq{}
}

func (p *SendMessageReq) GetClientId() (v string) {
	return p.ClientId
}

func (p *SendMessageReq) GetEvent() (v string) {
	return p.Event
}

func (p *SendMessageReq) GetType() (v string) {
	return p.Type
}

func (p *SendMessageReq) GetMessage() (v string) {
	return p.Message
}
func (p *SendMessageReq) SetClientId(val string) {
	p.ClientId = val
}
func (p *SendMessageReq) SetEvent(val string) {
	p.Event = val
}
func (p *SendMessageReq) SetType(val string) {
	p.Type = val
}
func (p *SendMessageReq) SetMessage(val string) {
	p.Message = val
}

var fieldIDToName_SendMessageReq = map[int16]string{
	1: "clientId",
	2: "event",
	3: "type",
	4: "message",
}

func (p *SendMessageReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetClientId bool = false
	var issetEvent bool = false
	var issetType bool = false
	var issetMessage bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetClientId = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetEvent = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetType = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
				issetMessage = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetClientId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetEvent {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetType {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetMessage {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SendMessageReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_SendMessageReq[fieldId]))
}

func (p *SendMessageReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ClientId = v
	}
	return nil
}

func (p *SendMessageReq) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Event = v
	}
	return nil
}

func (p *SendMessageReq) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Type = v
	}
	return nil
}

func (p *SendMessageReq) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *SendMessageReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SendMessageReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SendMessageReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("clientId", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ClientId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SendMessageReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("event", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Event); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SendMessageReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("type", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Type); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *SendMessageReq) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("message", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *SendMessageReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendMessageReq(%+v)", *p)
}

func (p *SendMessageReq) DeepEqual(ano *SendMessageReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ClientId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Event) {
		return false
	}
	if !p.Field3DeepEqual(ano.Type) {
		return false
	}
	if !p.Field4DeepEqual(ano.Message) {
		return false
	}
	return true
}

func (p *SendMessageReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.ClientId, src) != 0 {
		return false
	}
	return true
}
func (p *SendMessageReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Event, src) != 0 {
		return false
	}
	return true
}
func (p *SendMessageReq) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Type, src) != 0 {
		return false
	}
	return true
}
func (p *SendMessageReq) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Message, src) != 0 {
		return false
	}
	return true
}

type SendMessageResp struct {
	BaseResp *base.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,base.BaseResp" json:"baseResp"`
}

func NewSendMessageResp() *SendMessageResp {
	return &SendMessageResp{}
}

func (p *SendMessageResp) InitDefault() {
	*p = SendMessageResp{}
}

var SendMessageResp_BaseResp_DEFAULT *base.BaseResp

func (p *SendMessageResp) GetBaseResp() (v *base.BaseResp) {
	if !p.IsSetBaseResp() {
		return SendMessageResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *SendMessageResp) SetBaseResp(val *base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_SendMessageResp = map[int16]string{
	1: "baseResp",
}

func (p *SendMessageResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *SendMessageResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetBaseResp bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetBaseResp = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetBaseResp {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SendMessageResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_SendMessageResp[fieldId]))
}

func (p *SendMessageResp) ReadField1(iprot thrift.TProtocol) error {
	p.BaseResp = base.NewBaseResp()
	if err := p.BaseResp.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SendMessageResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SendMessageResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SendMessageResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("baseResp", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SendMessageResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendMessageResp(%+v)", *p)
}

func (p *SendMessageResp) DeepEqual(ano *SendMessageResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *SendMessageResp) Field1DeepEqual(src *base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

type MessageService interface {
	SendMessage(ctx context.Context, req *SendMessageReq) (r *SendMessageResp, err error)
}

type MessageServiceClient struct {
	c thrift.TClient
}

func NewMessageServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MessageServiceClient {
	return &MessageServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewMessageServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MessageServiceClient {
	return &MessageServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewMessageServiceClient(c thrift.TClient) *MessageServiceClient {
	return &MessageServiceClient{
		c: c,
	}
}

func (p *MessageServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *MessageServiceClient) SendMessage(ctx context.Context, req *SendMessageReq) (r *SendMessageResp, err error) {
	var _args MessageServiceSendMessageArgs
	_args.Req = req
	var _result MessageServiceSendMessageResult
	if err = p.Client_().Call(ctx, "SendMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type MessageServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      MessageService
}

func (p *MessageServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MessageServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MessageServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMessageServiceProcessor(handler MessageService) *MessageServiceProcessor {
	self := &MessageServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("SendMessage", &messageServiceProcessorSendMessage{handler: handler})
	return self
}
func (p *MessageServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type messageServiceProcessorSendMessage struct {
	handler MessageService
}

func (p *messageServiceProcessorSendMessage) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MessageServiceSendMessageArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("SendMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := MessageServiceSendMessageResult{}
	var retval *SendMessageResp
	if retval, err2 = p.handler.SendMessage(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SendMessage: "+err2.Error())
		oprot.WriteMessageBegin("SendMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("SendMessage", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type MessageServiceSendMessageArgs struct {
	Req *SendMessageReq `thrift:"req,1" frugal:"1,default,SendMessageReq" json:"req"`
}

func NewMessageServiceSendMessageArgs() *MessageServiceSendMessageArgs {
	return &MessageServiceSendMessageArgs{}
}

func (p *MessageServiceSendMessageArgs) InitDefault() {
	*p = MessageServiceSendMessageArgs{}
}

var MessageServiceSendMessageArgs_Req_DEFAULT *SendMessageReq

func (p *MessageServiceSendMessageArgs) GetReq() (v *SendMessageReq) {
	if !p.IsSetReq() {
		return MessageServiceSendMessageArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *MessageServiceSendMessageArgs) SetReq(val *SendMessageReq) {
	p.Req = val
}

var fieldIDToName_MessageServiceSendMessageArgs = map[int16]string{
	1: "req",
}

func (p *MessageServiceSendMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MessageServiceSendMessageArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MessageServiceSendMessageArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MessageServiceSendMessageArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewSendMessageReq()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *MessageServiceSendMessageArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SendMessage_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MessageServiceSendMessageArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *MessageServiceSendMessageArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MessageServiceSendMessageArgs(%+v)", *p)
}

func (p *MessageServiceSendMessageArgs) DeepEqual(ano *MessageServiceSendMessageArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *MessageServiceSendMessageArgs) Field1DeepEqual(src *SendMessageReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type MessageServiceSendMessageResult struct {
	Success *SendMessageResp `thrift:"success,0,optional" frugal:"0,optional,SendMessageResp" json:"success,omitempty"`
}

func NewMessageServiceSendMessageResult() *MessageServiceSendMessageResult {
	return &MessageServiceSendMessageResult{}
}

func (p *MessageServiceSendMessageResult) InitDefault() {
	*p = MessageServiceSendMessageResult{}
}

var MessageServiceSendMessageResult_Success_DEFAULT *SendMessageResp

func (p *MessageServiceSendMessageResult) GetSuccess() (v *SendMessageResp) {
	if !p.IsSetSuccess() {
		return MessageServiceSendMessageResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MessageServiceSendMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*SendMessageResp)
}

var fieldIDToName_MessageServiceSendMessageResult = map[int16]string{
	0: "success",
}

func (p *MessageServiceSendMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessageServiceSendMessageResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MessageServiceSendMessageResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MessageServiceSendMessageResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewSendMessageResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *MessageServiceSendMessageResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SendMessage_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MessageServiceSendMessageResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *MessageServiceSendMessageResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MessageServiceSendMessageResult(%+v)", *p)
}

func (p *MessageServiceSendMessageResult) DeepEqual(ano *MessageServiceSendMessageResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *MessageServiceSendMessageResult) Field0DeepEqual(src *SendMessageResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
