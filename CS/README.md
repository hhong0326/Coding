# CS 정리 

## CS part

### 네트워크

- OSI 7 layer

* 물리 계층(Physical Layer)
    PDU : Bit Stream(0과 1의 연속 / 전기적 신호)
    Protocol : Ethernet RS-232C
    Equipment : Repeater / Hub
* 데이터링크 계층(Datalink Layer)
    PDU : Frame
    Protocol : Ethernet / HDLC / PPP / ...
    Equipment : Bridge / L2 Switch
* 네트워크 계층(Network Layer)
    PDU : Packet
    Protocol : IP / ARP / ICMP
    Equipment : Router / L3 Switch
* 전송 계층(Transport Layer)
    PDU : Segment
    Protocol : TCP / UDP
    Equipment : L4 Switch, Gateway
* 세션 계층(Session Layer)
* 표현 계층(Presentation Layer)
    입출력 데이터를 하나의 표현 형태로 변환 (이해할 수 있는 포맷으로 변환)
    Protocol : JPEG, MPEG, SMG, AFP
* 응용 계층(Application Layer)
    PDU : Data or Message
    Protocol : Telnet / HTTP / FTP / SSH  ...
    Equipment : PC / Server / ...

### 통신 프로토콜
#### IP(Internet Protocol): TCP/IP 네트워크에서 출발지, 목적지 주소를 지정하는 프로토콜
- Packet의 목적지 주소를 보고 최적의 경로를 찾아 패킷을 전송해주는게 주된 역할
- 신뢰성이 없고 비연결 지향적임(신뢰성보다는 효율성에 중점)

### API 

### 쓰레드와 프로세스 (+메모리 스택 구조)

### 데이터베이스

- 페이징 관련

## Golang part

### 장단점 

### 경량 쓰레드 Go routine (go)

### 쓰레드 통신 Go Channel (chan)

