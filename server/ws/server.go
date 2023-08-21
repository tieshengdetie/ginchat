package ws

import (
	"encoding/json"
	"fmt"
	"ginchat/models"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[uint64]*Node = make(map[uint64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(wirter http.ResponseWriter, request *http.Request) {
	//校验token
	query := request.URL.Query()
	//token := query.Get("token")
	id := query.Get("userId")

	userId, _ := strconv.ParseUint(id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//content := query.Get("content")
	var checkOrigin bool = true
	//_, err := utils.ParseToken(token)
	//if err != nil {
	//	checkOrigin = false
	//}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return checkOrigin
		},
	}
	ws, err := upgrader.Upgrade(wirter, request, nil)
	if err != nil {
		fmt.Println("websocket连接错误")
		return
	}

	//当前连接
	var node = &Node{
		Conn:      ws,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//绑定关系（当前用户和当前连接）
	//加锁
	rwLocker.Lock()
	clientMap[userId] = node
	//解锁
	rwLocker.Unlock()
	sendMsg(userId, []byte("欢迎连接"))
	//起协程完成发送消息逻辑
	go sendProc(node)
	//完成接收逻辑
	go recvProc(node)

}

// 像管道中推送数据
func sendMsg(targetId uint64, message []byte) {
	//读取连接池中的数据，加读锁
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- message
	}
}

// 监听管道中的数据并发送数据
func sendProc(node *Node) {

	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				fmt.Println("发送消息失败")
				return
			}
		}

	}
}

// 接收websocket中的连接发送过来的消息，并转发出去
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			fmt.Println("消息接收失败")
			return
		}
		//将数据分发
		//broadMsg(data)
		dispath(data)
	}
}
func broadMsg(data []byte) {
	udpMsgChan <- data
}

// 消息分发
func dispath(data []byte) {

	fmt.Println("接收到的数据 :", string(data))
	//解析数据
	message := &models.Message{}

	err := json.Unmarshal(data, message)
	if err != nil {
		fmt.Println("数据解析失败")
		fmt.Println(err)
		return
	}
	fmt.Println("解析的数据为", message)
	//存入管道
	sendMsg(message.TargetId, data)
}

// 定义udp 消息存放管道
var udpMsgChan chan []byte = make(chan []byte, 1024)

// udp 方式数据发送
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		select {
		case data := <-udpMsgChan:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// udp 方式数据接收
func udpRecvProc() {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		dispath(buf[0:n])

	}
}

func init() {
	go udpRecvProc()
	go udpRecvProc()
}
