package ws

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type WebSocket struct {
	conn    net.Conn
	handler http.HandlerFunc
}

var pool = map[int64]net.Conn{}

func NewWS() *WebSocket {
	return &WebSocket{}
}

func (websocket *WebSocket) RegisterHandlers(handlerFunc http.HandlerFunc) {
	websocket.handler = handlerFunc
}

func (websocket *WebSocket) Run(port string) error {
	return http.ListenAndServe(":"+port, websocket.handler)
}

type myCalculator interface {
	Calculate(string) (float64, error)
}

type store interface {
	Store(ctx context.Context, s string) error
	List(ctx context.Context) ([]string, error)
}

type CalcHandler struct {
	calc    myCalculator
	storage store
	out     chan string
}

func New(storage store, calc myCalculator) CalcHandler {
	h := CalcHandler{
		calc:    calc,
		storage: storage,
		out:     make(chan string, 100),
	}

	go h.broadcastWriter()

	return h
}

func (h CalcHandler) Handle() http.HandlerFunc {
	return func(resp http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, resp)
		if err != nil {
			return
		}

		pool[timestamp()] = conn
		h.updateConn(r.Context(), conn)

		go func() {
			defer func() {
				conn.Close()
			}()

			for {
				msg, err := wsutil.ReadClientText(conn)
				if err != nil {
					fmt.Printf("can't read data from clinet with error: %+v\n", err)
					return
				}

				result, err := h.calc.Calculate(string(msg))
				if err != nil {
					//TODO: send error message to client also
					fmt.Printf("can't calculate expression: %+v\n", err)
					continue
				}

				h.storage.Store(r.Context(), fmt.Sprintf("%s = %.5f", msg, result))
				h.update(r.Context())
			}
		}()
	}
}

func (h CalcHandler) update(ctx context.Context) {
	list, err := h.storage.List(ctx)
	if err != nil {
		fmt.Printf("can't receive list from storage %+v\n", err)
		return
	}
	if len(list) == 0 {
		return
	}

	h.out <- strings.Join(list, ",")
}

func (h CalcHandler) updateConn(ctx context.Context, conn net.Conn) {
	list, err := h.storage.List(ctx)
	if err != nil {
		fmt.Printf("can't receive list from storage %+v\n", err)
		return
	}
	if len(list) == 0 {
		return
	}
	h.writer(conn, strings.Join(list, ","))
}

func (h CalcHandler) broadcastWriter() {
	for msg := range h.out {
		for i, u := range pool {
			err := h.writer(u, msg)
			if err != nil {
				fmt.Printf("can't write data to client with error %+v\n", err)
				delete(pool, i)
			}
		}
	}
}

func (h CalcHandler) writer(c net.Conn, msg string) error {
	return wsutil.WriteServerText(c, []byte(msg))
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
