# CS 정리 

## CS part

### 네트워크

#### OSI 7 layer

![OSI 7 layer]https://blog.kakaocdn.net/dn/rOOKR/btrpUHuv9WV/0d09FTgzKuWIhrizkCWoR0/img.gif

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

### OSI 계층별 통신 프로토콜

#### IP(Internet Protocol): TCP/IP 네트워크에서 출발지, 목적지 주소를 지정하는 프로토콜 *네트워크계층*
- Packet의 목적지 주소를 보고 최적의 경로를 찾아 패킷을 전송해주는게 주된 역할
- 신뢰성이 없고 비연결 지향적임(신뢰성보다는 효율성에 중점)


#### TCP *전송계층*
- Transport Layer의 프로토콜로써, 신뢰성과 연결 지향적 특징을 가짐
- 혼잡 제어와 흐름 제어 기능을 제공(Sliding Window) / 에러 제어도 가능(Checksum)
- 재전송 방법 : GO-Back-N 방법 (되돌아온 ACK 번호 이후의 것 전부 재전송)
- ICMP 프로토콜로 주기적으로 송수신 가능 여부 체크
- 클라이언트와 서버의 데이터 교환을 위해 TCP 3-Way Handshake를 통해 TCP Session을 확립해야 함

#### UDP *전송계층*
- Transport Layer의 프로토콜로써, TCP와 달리 비신뢰성과 비연결지향적인 특징을 가짐
- Sequence Number와 Ack Number가 없어서 순서 제어와 흐름 제어가 불가능
- 따로 연결하는 과정이 없어서 빠른처리와 실시간성을 요구하는 서비스에 적합

#### FTP *응용계층*
- 원격 파일 전송 프로토콜 20/21 port
#### HTTP *응용계층*
- 웹 전송 통신 프로토콜 80 port
- HTTP 1.1 -  Keep Alive Connection : 연결을 종료하지 않고 유지하기 때문에, 서비스 요청 이후에 바로 연결을 종료하지 않고 연결 대기 이후 종료
- HTTP 2.0 -  더 빠른 데이터 처리 가능 / 멀티플렉싱(Multiplexing) : 동시 다발적 양방향 통신 / 헤더압축(Header Compression) : 헤더 정보를 1/3 수준으로 압축 / 서버푸시(Server Push) : 웹 서버가 웹 브라우저에게 필요한 데이터를 알아서 미리 전송
#### SSH *응용계층*
- 원격 접속 프로토콜 Telnet -> 보안성을 위해 사용 22 port

### API + RESTful
- Application Interface
- 

### 쓰레드와 프로세스 (+메모리 스택 구조)

### 데이터베이스

- 페이징 관련

## Golang part

### 장단점 

### 경량 쓰레드 Go routine (go)

### 쓰레드 통신 Go Channel (chan)

