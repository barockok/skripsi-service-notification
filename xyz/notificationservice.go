// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package xyz

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type NotificationService interface {
	// Parameters:
	//  - URL
	//  - Transaction
	Send(url string, transaction *Transaction) (r string, err error)
}

type NotificationServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewNotificationServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *NotificationServiceClient {
	return &NotificationServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewNotificationServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *NotificationServiceClient {
	return &NotificationServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - URL
//  - Transaction
func (p *NotificationServiceClient) Send(url string, transaction *Transaction) (r string, err error) {
	if err = p.sendSend(url, transaction); err != nil {
		return
	}
	return p.recvSend()
}

func (p *NotificationServiceClient) sendSend(url string, transaction *Transaction) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("send", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := NotificationServiceSendArgs{
		URL:         url,
		Transaction: transaction,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *NotificationServiceClient) recvSend() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "send" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "send failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "send failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error5 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error6 error
		error6, err = error5.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error6
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "send failed: invalid message type")
		return
	}
	result := NotificationServiceSendResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type NotificationServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      NotificationService
}

func (p *NotificationServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *NotificationServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *NotificationServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewNotificationServiceProcessor(handler NotificationService) *NotificationServiceProcessor {

	self7 := &NotificationServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self7.processorMap["send"] = &notificationServiceProcessorSend{handler: handler}
	return self7
}

func (p *NotificationServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x8 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x8.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x8

}

type notificationServiceProcessorSend struct {
	handler NotificationService
}

func (p *notificationServiceProcessorSend) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NotificationServiceSendArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("send", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := NotificationServiceSendResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Send(args.URL, args.Transaction); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing send: "+err2.Error())
		oprot.WriteMessageBegin("send", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("send", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - URL
//  - Transaction
type NotificationServiceSendArgs struct {
	URL         string       `thrift:"url,1" json:"url"`
	Transaction *Transaction `thrift:"transaction,2" json:"transaction"`
}

func NewNotificationServiceSendArgs() *NotificationServiceSendArgs {
	return &NotificationServiceSendArgs{}
}

func (p *NotificationServiceSendArgs) GetURL() string {
	return p.URL
}

var NotificationServiceSendArgs_Transaction_DEFAULT *Transaction

func (p *NotificationServiceSendArgs) GetTransaction() *Transaction {
	if !p.IsSetTransaction() {
		return NotificationServiceSendArgs_Transaction_DEFAULT
	}
	return p.Transaction
}
func (p *NotificationServiceSendArgs) IsSetTransaction() bool {
	return p.Transaction != nil
}

func (p *NotificationServiceSendArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NotificationServiceSendArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.URL = v
	}
	return nil
}

func (p *NotificationServiceSendArgs) readField2(iprot thrift.TProtocol) error {
	p.Transaction = &Transaction{}
	if err := p.Transaction.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Transaction), err)
	}
	return nil
}

func (p *NotificationServiceSendArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("send_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NotificationServiceSendArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("url", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:url: ", p), err)
	}
	if err := oprot.WriteString(string(p.URL)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.url (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:url: ", p), err)
	}
	return err
}

func (p *NotificationServiceSendArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("transaction", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:transaction: ", p), err)
	}
	if err := p.Transaction.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Transaction), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:transaction: ", p), err)
	}
	return err
}

func (p *NotificationServiceSendArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NotificationServiceSendArgs(%+v)", *p)
}

// Attributes:
//  - Success
type NotificationServiceSendResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewNotificationServiceSendResult() *NotificationServiceSendResult {
	return &NotificationServiceSendResult{}
}

var NotificationServiceSendResult_Success_DEFAULT string

func (p *NotificationServiceSendResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return NotificationServiceSendResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *NotificationServiceSendResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NotificationServiceSendResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NotificationServiceSendResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *NotificationServiceSendResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("send_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NotificationServiceSendResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *NotificationServiceSendResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NotificationServiceSendResult(%+v)", *p)
}
