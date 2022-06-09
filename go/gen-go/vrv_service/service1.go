// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package vrv_service

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Name
type Service1HelloReuqest struct {
  Name string `thrift:"name,1" db:"name" json:"name"`
}

func NewService1HelloReuqest() *Service1HelloReuqest {
  return &Service1HelloReuqest{}
}


func (p *Service1HelloReuqest) GetName() string {
  return p.Name
}
func (p *Service1HelloReuqest) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Service1HelloReuqest)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *Service1HelloReuqest) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Service1HelloReuqest"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Service1HelloReuqest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *Service1HelloReuqest) Equals(other *Service1HelloReuqest) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Name != other.Name { return false }
  return true
}

func (p *Service1HelloReuqest) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Service1HelloReuqest(%+v)", *p)
}

// Attributes:
//  - Code
//  - Message
//  - Res
type Service1HelloResponse struct {
  Code string `thrift:"code,1" db:"code" json:"code"`
  Message *string `thrift:"message,2" db:"message" json:"message,omitempty"`
  Res string `thrift:"res,3" db:"res" json:"res"`
}

func NewService1HelloResponse() *Service1HelloResponse {
  return &Service1HelloResponse{}
}


func (p *Service1HelloResponse) GetCode() string {
  return p.Code
}
var Service1HelloResponse_Message_DEFAULT string
func (p *Service1HelloResponse) GetMessage() string {
  if !p.IsSetMessage() {
    return Service1HelloResponse_Message_DEFAULT
  }
return *p.Message
}

func (p *Service1HelloResponse) GetRes() string {
  return p.Res
}
func (p *Service1HelloResponse) IsSetMessage() bool {
  return p.Message != nil
}

func (p *Service1HelloResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Service1HelloResponse)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Code = v
}
  return nil
}

func (p *Service1HelloResponse)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Message = &v
}
  return nil
}

func (p *Service1HelloResponse)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Res = v
}
  return nil
}

func (p *Service1HelloResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Service1HelloResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Service1HelloResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "code", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Code)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err) }
  return err
}

func (p *Service1HelloResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetMessage() {
    if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err) }
    if err := oprot.WriteString(ctx, string(*p.Message)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err) }
  }
  return err
}

func (p *Service1HelloResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "res", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:res: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Res)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.res (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:res: ", p), err) }
  return err
}

func (p *Service1HelloResponse) Equals(other *Service1HelloResponse) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Code != other.Code { return false }
  if p.Message != other.Message {
    if p.Message == nil || other.Message == nil {
      return false
    }
    if (*p.Message) != (*other.Message) { return false }
  }
  if p.Res != other.Res { return false }
  return true
}

func (p *Service1HelloResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Service1HelloResponse(%+v)", *p)
}

type Service1 interface {
  // Parameters:
  //  - Request
  Hello(ctx context.Context, request *Service1HelloReuqest) (_r *Service1HelloResponse, _err error)
}

type Service1Client struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewService1ClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *Service1Client {
  return &Service1Client{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewService1ClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *Service1Client {
  return &Service1Client{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewService1Client(c thrift.TClient) *Service1Client {
  return &Service1Client{
    c: c,
  }
}

func (p *Service1Client) Client_() thrift.TClient {
  return p.c
}

func (p *Service1Client) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *Service1Client) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Request
func (p *Service1Client) Hello(ctx context.Context, request *Service1HelloReuqest) (_r *Service1HelloResponse, _err error) {
  var _args0 Service1HelloArgs
  _args0.Request = request
  var _result2 Service1HelloResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "hello", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

type Service1Processor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Service1
}

func (p *Service1Processor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *Service1Processor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *Service1Processor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewService1Processor(handler Service1) *Service1Processor {

  self3 := &Service1Processor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self3.processorMap["hello"] = &service1ProcessorHello{handler:handler}
return self3
}

func (p *Service1Processor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x4.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x4

}

type service1ProcessorHello struct {
  handler Service1
}

func (p *service1ProcessorHello) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := Service1HelloArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "hello", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := Service1HelloResult{}
  var retval *Service1HelloResponse
  if retval, err2 = p.handler.Hello(ctx, args.Request); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing hello: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "hello", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "hello", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type Service1HelloArgs struct {
  Request *Service1HelloReuqest `thrift:"request,1" db:"request" json:"request"`
}

func NewService1HelloArgs() *Service1HelloArgs {
  return &Service1HelloArgs{}
}

var Service1HelloArgs_Request_DEFAULT *Service1HelloReuqest
func (p *Service1HelloArgs) GetRequest() *Service1HelloReuqest {
  if !p.IsSetRequest() {
    return Service1HelloArgs_Request_DEFAULT
  }
return p.Request
}
func (p *Service1HelloArgs) IsSetRequest() bool {
  return p.Request != nil
}

func (p *Service1HelloArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Service1HelloArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Request = &Service1HelloReuqest{}
  if err := p.Request.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
  }
  return nil
}

func (p *Service1HelloArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "hello_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Service1HelloArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err) }
  if err := p.Request.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err) }
  return err
}

func (p *Service1HelloArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Service1HelloArgs(%+v)", *p)
}

// Attributes:
//  - Success
type Service1HelloResult struct {
  Success *Service1HelloResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewService1HelloResult() *Service1HelloResult {
  return &Service1HelloResult{}
}

var Service1HelloResult_Success_DEFAULT *Service1HelloResponse
func (p *Service1HelloResult) GetSuccess() *Service1HelloResponse {
  if !p.IsSetSuccess() {
    return Service1HelloResult_Success_DEFAULT
  }
return p.Success
}
func (p *Service1HelloResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *Service1HelloResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Service1HelloResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &Service1HelloResponse{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *Service1HelloResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "hello_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Service1HelloResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *Service1HelloResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Service1HelloResult(%+v)", *p)
}


