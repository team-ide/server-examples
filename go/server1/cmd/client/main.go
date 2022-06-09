package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/team-ide/server-examples/go/gen-go/vrv_service"
	server_transport "github.com/team-ide/server-examples/go/server1/transport"
)

func main() {
	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an addsvc. This presumption is reflected in
	// the addcli binary and the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings. For an
	// example service with a client built on top of a service discovery system,
	// see profilesvc.
	fs := flag.NewFlagSet("addcli", flag.ExitOnError)
	var (
		thriftAddr     = fs.String("thrift-addr", "", "Thrift address of addsvc")
		thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
		thriftBuffer   = fs.Int("thrift-buffer", 0, "0 for unbuffered")
		thriftFramed   = fs.Bool("thrift-framed", false, "true to enable framing")
		method         = fs.String("method", "hello", "sum, concat")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags] <a> <b>")
	fs.Parse(os.Args[1:])
	if len(fs.Args()) != 2 {
		fs.Usage()
		os.Exit(1)
	}

	// This is a demonstration client, which supports multiple transports.
	// Your clients will probably just define and stick with 1 transport.
	var (
		svc vrv_service.Service1
		err error
	)
	// It's necessary to do all of this construction in the func main,
	// because (among other reasons) we need to control the lifecycle of the
	// Thrift transport, i.e. close it eventually.
	var protocolFactory thrift.TProtocolFactory
	switch *thriftProtocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprintf(os.Stderr, "error: invalid protocol %q\n", *thriftProtocol)
		os.Exit(1)
	}
	var transportFactory thrift.TTransportFactory
	if *thriftBuffer > 0 {
		transportFactory = thrift.NewTBufferedTransportFactory(*thriftBuffer)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if *thriftFramed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}
	transportSocket, err := thrift.NewTSocket(*thriftAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	transport, err := transportFactory.GetTransport(transportSocket)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if err := transport.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer transport.Close()
	client := vrv_service.NewService1ClientFactory(transport, protocolFactory)
	svc = server_transport.NewThriftClient(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch *method {
	case "hello":
		req := &vrv_service.Service1HelloReuqest{}
		var res *vrv_service.Service1HelloResponse
		res, err := svc.Hello(context.Background(), req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%v\n", res)

	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", *method)
		os.Exit(1)
	}
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
