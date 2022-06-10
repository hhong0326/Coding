# CS 정리 

- [Part 1.CS](#part-1cs)
- [Part 2.Language](#part-2language)


## Part 1.CS

### 네트워크

#### OSI 7 layer

![OSI 7 layer](https://blog.kakaocdn.net/dn/rOOKR/btrpUHuv9WV/0d09FTgzKuWIhrizkCWoR0/img.gif)

* 물리 계층(Physical Layer)
    * PDU : Bit Stream(0과 1의 연속 / 전기적 신호)
    * Protocol : Ethernet RS-232C
    * Equipment : Repeater / Hub
* 데이터링크 계층(Datalink Layer)
    * PDU : Frame
    * Protocol : Ethernet / HDLC / PPP / ...
    * Equipment : Bridge / L2 Switch
* 네트워크 계층(Network Layer)
    * PDU : Packet
    * Protocol : IP / ARP / ICMP
    * Equipment : Router / L3 Switch
* 전송 계층(Transport Layer)
    * PDU : Segment
    * Protocol : TCP / UDP
    * Equipment : L4 Switch, Gateway
* 세션 계층(Session Layer)
* 표현 계층(Presentation Layer)
    * 입출력 데이터를 하나의 표현 형태로 변환 (이해할 수 있는 포맷으로 변환)
    * Protocol : JPEG, MPEG, SMG, AFP
* 응용 계층(Application Layer)
    * PDU : Data or Message
    * Protocol : Telnet / HTTP / FTP / SSH  ...
    * Equipment : PC / Server / ...

### OSI 계층별 통신 프로토콜

#### *네트워크계층*
##### IP(Internet Protocol): TCP/IP 네트워크에서 출발지, 목적지 주소를 지정하는 프로토콜 
- Packet의 목적지 주소를 보고 최적의 경로를 찾아 패킷을 전송해주는게 주된 역할
- 신뢰성이 없고 비연결 지향적임(신뢰성보다는 효율성에 중점)

#### *전송계층*
##### TCP 
- Transport Layer의 프로토콜로써, 신뢰성과 연결 지향적 특징을 가짐
- 혼잡 제어와 흐름 제어 기능을 제공(Sliding Window) / 에러 제어도 가능(Checksum)
- 재전송 방법 : GO-Back-N 방법 (되돌아온 ACK 번호 이후의 것 전부 재전송)
- ICMP 프로토콜로 주기적으로 송수신 가능 여부 체크
- 클라이언트와 서버의 데이터 교환을 위해 TCP 3-Way Handshake를 통해 TCP Session을 확립해야 함

##### UDP 
- Transport Layer의 프로토콜로써, TCP와 달리 비신뢰성과 비연결지향적인 특징을 가짐
- Sequence Number와 Ack Number가 없어서 순서 제어와 흐름 제어가 불가능
- 따로 연결하는 과정이 없어서 빠른처리와 실시간성을 요구하는 서비스에 적합

#### *응용계층*
##### FTP
- 원격 파일 전송 프로토콜 20/21 port
##### HTTP
- 웹 전송 통신 프로토콜 80 port
- HTTP 1.1 -  Keep Alive Connection : 연결을 종료하지 않고 유지하기 때문에, 서비스 요청 이후에 바로 연결을 종료하지 않고 연결 대기 이후 종료
- HTTP 2.0 
    - 더 빠른 데이터 처리 가능 
    - 멀티플렉싱(Multiplexing) : 동시 다발적 양방향 통신
    - 헤더압축(Header Compression) : 헤더 정보를 1/3 수준으로 압축
    - 서버푸시(Server Push) : 웹 서버가 웹 브라우저에게 필요한 데이터를 알아서 미리 전송
##### SSH
- 원격 접속 프로토콜 Telnet -> 보안성을 위해 사용 22 port

### API + RESTful
- Application Interface
- 
#### RESTful
`REST` 란, REpresentational State Transfer 의 약자이다. 여기에 ~ful 이라는 형용사형 어미를 붙여 ~한 API 라는 표현으로 사용된다. 즉, REST 의 기본 원칙을 성실히 지킨 서비스 디자인은 'RESTful'하다라고 표현할 수 있다.

REST가 디자인 패턴이다, 아키텍처다 많은 이야기가 존재하는데, 하나의 `아키텍처`로 볼 수 있다. 좀 더 정확한 표현으로 말하자면, REST 는 Resource Oriented Architecture 이다. API 설계의 중심에 `자원(Resource)이 있고 HTTP Method 를 통해 자원을 처리하도록 설계`하는 것이다.

### 쓰레드와 프로세스 (+메모리 스택 구조)

### 데이터베이스

#### 트랜잭션
데이터 상태를 변환시키는 `하나의 논리적 기능을 수행하기 위한 작업의 단위` 또는 `한꺼번에 모두 수행되어야 할 일련의 연산들`을 의미

- 데이터베이스 시스템에서 병행 제어 및 회복 작업 시 처리되는 작업의 논리적 단위
- 사용자가 시스템에 대한 서비스 요구 시 시스템이 응답하기 위한 상태 변환 과정의 작업 단위
- 하나의 트랜잭션은 `Commmi`t 또는 `Rollback` 된다

##### 특징
- **Atomicity(원자성)**
    - 트랜잭션 연산은 DB에 모두 반영되든지 전혀 반영되지 않아야 한다
    - 트랜잭션 내 모든 명령 또한 완벽하게 수행되어야 하며, 오류 발생 시 전부 취소되어야 한다
- **Consistency(일관성)**
    - 트랜잭션이 실행을 완료하면 언제나 일관성있는 데이터베이스 상태로 변환한다
    - 시스템이 가지고 있는 고정요소는 트랜잭션 수행 전과 수행 이후의 상태가 같아야 한다
- **Isolation(독립성)**
    - 둘 이상의 트랜잭션이 동시에 병렬 실행되는 경우 어느 하나의 트랜잭션 연산 실행중에 다른 트랜잭션의 연산이 끼어들 수 없다
    - 트랜잭션이 완료될 때까지 다른 트랜잭션 수행 결과를 참조할 수 없다
- **Durability(지속성)**
    - 트랜잭션 결과는 시스템이 고장나도 영구적으로 반영되어야 한다

##### Commit
한 개의 논리적 단위인 트랜잭션의 작업이 성공적으로 끝나고 DB가 다시 일관된 상태에 있을 때, 완료된 것을 트랜잭션 관리자에게 알리는 연산이다

##### Rollback
- 트랜잭셕 처리가 비정상적으로 종료되어 DB의 일관성을 깨뜨렸을때, 일부가 정상처리되도 원자성 구현을 위해 해당 트랜잭션이 행한 모든 연산을 취소(Undo) 하는 연산이다
- Rollback 연산 시 해당 트랜잭션을 재시작하거나 폐기한다

#### Index
인덱스는 데이터베이스 테이블에 대한 `검색 성능의 속도를 높여주는` 자료 구조입니다. 특정 컬럼에 인덱스를 생성하면, 해당 컬럼의 데이터들을 정렬하여 `별도의 메모리 공간`에 데이터의 `물리적 주소`와 함께 저장됩니다.

Where 절에 지정한 인덱스 컬럼이 있을 때, 옵티마이저에서 판단하여 생성된 인덱스를 탈 수가 있습니다. 
인덱스에 저장되어 있는 데이터의 물리적 주소로 가서 데이터를 가져오는 동작을 하여 검색 속도의 향상을 가져올 수 있습니다.

![Index Table](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FcQi8RP%2Fbtq8BkRrRfb%2Fa5C0jH5pfSA2KKz7C9fB7k%2Fimg.png)

##### Index를 사용하는 이유
- 데이터들이 정렬되어 있기 때문에, 테이블을 스캔할 때 해당 Where 조건에 맞는 데이터들을 더 빠르게 찾을 수 있다
- 또한 Order by에 의한 정렬 과정은 굉장히 부하가 많이 걸리는 작업인데 인덱스를 사용하면 이를 피할 수 있다
    - Order by: 정렬과 동시에 1차적으로 메모리 정렬, 메모리보다 큰 작업은 디스크 I/O 발생으로 자원 소모
- MIN, MAX 효율적인 처리 가능


##### Index를 선택하는 기준
- 카디널리티(= 해당 컬럼에 대한 중복된 수치에 대한 상태)가 높은 순서로 고른다 (카디널리티란 컬럼의 중복도에 대한 것)
- 중복도가 낮으면 카디널리티가 높다 (PK FK) / 중복도가 높으면 카디널리티가 낮다 (일반 컬럼)
- 인덱스를 1개의 컬럼에만 걸어야 한다면 카디널리티가 가장 높은 것을 잡아야 한다
- 되도록 업데이트가 빈번하지 않은 컬럼으로 인덱스를 구성한다
- where 절에서 자주 사용하고 "=" 으로 비교되는 컬럼에는 인덱스 추가를 고려한다
- join시 자주 사용하는 컬럼은 인덱스로 등록한다. 단일 인덱스 여러개보다 다중컬럼 인덱스의 생성을 고려한다
- ORDER BY 절에서 자주 사용하는 컬럼은 인덱스로 등록한다
- 인덱스 테이블 생성시 DB의 약 10%에 해당하는 저장공간이 필요하므로 전체 데이터 중에 10~15% 이하의 데이터를 처리하는 경우에 효율적임을 고려한다

#### Pagination
서버와 클라이언트의 통신 상황에서 모든 데이터를 한꺼번에 가져오지 않고 `필요한 갯수`를 지정하고 상황에 맞춰 `정렬 기준`이 추가된다
이러한 조건에 맞춰 데이터를 가져오는 것을 Pagination, 페이징이라고 한다

##### Offset-based Pagination
SQL에서는 LIMIT을 활용하며, 절대값으로 갯수를 지정하기 때문에 중복된 데이터값이 출력하게 되어 성능이슈가 있을 수 있다

##### Cursor-based Pagination
이를 보안하기 위해 커서 기반 페이징은 포인터와 비슷한 cursor를 만든 다음, cursor가 가리키는 것부터 다음 n개의 데이터를 요구하는 방식이다.

커서가 매우 매력적으로 보이지만 주의해야할 사항은 커서 기반 페이징을 위해 `정렬 기준`에 포함되는 필드 중 하나 이상은 반드시 `고유값`을 가져야 한다는 것이다.
고유값은 WHERE 절의 조건들을 이용하여 새로운 고유값, cursor를 만들어서 해결할 수 있다

*참조*
> [[DB] 데이터베이스 인덱스(Index) 란 무엇인가?](https://coding-factory.tistory.com/746)
> [기술 면접에 자주 나오는 질문들 - 인덱스](https://fors.tistory.com/641?category=882553)
> [Pagenation(페이징) | Offset-based, Cursor-based](https://daeuungcode.tistory.com/128)


- 페이징 관련

## Part 2.Language

### 장단점 

### 경량 쓰레드 Go routine (go)

### 쓰레드 통신 Go Channel (chan)

> [](https://github.com/JaeYeopHan/Interview_Question_for_Beginner)