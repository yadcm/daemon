package local

//@todo find local addr
//@todo cidr from config
// import (
// 	"context"
// 	protocol "yadcmd/internal/pb/yadcmd.daemon"
// 	"fmt"
// 	"log"
// 	"net/netip"
// 	"os"
// 	"sync"
// 	"time"
// )

// const (
// 	cidr      string = "10.0.0.0/24"
// 	port      uint16 = 49069
// 	batchSize int    = 80
// )

// func Scan() error {
// 	prefix, err := netip.ParsePrefix(cidr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var data string
// 	var addr netip.Addr
// 	var wg sync.WaitGroup
// 	var addrs []netip.Addr = make([]netip.Addr, 0)
// 	for addr = prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
// 		addrs = append(addrs, addr)
// 		wg.Add(1)
// 		if len(addrs) == batchSize {
// 			var result chan netip.AddrPort = make(chan netip.AddrPort, batchSize)
// 			go scanBatch(addrs, &wg, result)
// 			go func() {
// 				wg.Wait()
// 				wg = sync.WaitGroup{}
// 				addrs = make([]netip.Addr, 0)
// 				close(result)
// 			}()
// 			for found := range result {
// 				data += fmt.Sprintln(found)
// 			}
// 		}
// 	}
// 	return os.WriteFile(".knownhosts", []byte(data), os.ModePerm)
// }

// func scanBatch(addrs []netip.Addr, wg *sync.WaitGroup, result chan netip.AddrPort) {
// 	for _, addr := range addrs {
// 		go func(addr netip.Addr) {
// 			defer wg.Done()
// 			addrPort := netip.AddrPortFrom(addr, port)

// 			cl, errConnect := NewClient(addrPort)
// 			if errConnect != nil {
// 				log.Printf("INFO: protocol mismatch %s: %s", addr, errConnect)
// 				return
// 			}
// 			defer cl.Close()

// 			req := model.NewAgreementRequest(protocol.AgreementMethod_MOD)
// 			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 			defer cancel()
// 			resp, err := cl.AgreementStart(ctx, req)
// 			if err != nil {
// 				log.Printf("INFO: protocol mismatch %s: %s", addr, err)
// 				return
// 			}
// 			if resp == nil || resp.Result == nil {
// 				log.Printf("INFO: protocol mismatch %s: resp is nil", addr)
// 				return
// 			}
// 			got := *resp.Result
// 			want := model.ComputeAgreement(req)
// 			if got != want {
// 				log.Printf("INFO: protocol mismatch %s: got=%f, want=%f", addr, got, want)
// 				return
// 			}
// 			result <- addrPort
// 		}(addr)
// 	}
// }
