# gRPC
- RPC : (원격 프로시저 호출) 원격으로 method call을 가능하게 하여 별도읜 원격 제어를 위한 코딩 없이 다른 주소 공간에서 함수나 프로시저를 실행할 수 있게 하는 프로세스(서버) 간 통신 기술
- gRPC: **Protocol buffer** 및 **HTTP/2** 방식으로 구현된 Google에서 시작된 오픈소스. RPC 프로토콜을 위한 프레임워크이며, 분산 클라이언트-서버 아키텍처에서 통신을 매우 효율적으로 처리하는 최신 고성능 RPC 시스템

**Protocol buffer**

Protocol buffers는 구조화된 데이터를 직렬화하기 위한 유연하고 효율적인 자동화 메커니즘으로, XML보다 더 작고 빠르며 간단한 기법이다. 여러분의 데이터가 어떻게 구조화될지를 한번 정의하면, 다양한 언어를 가지고 발생된 특별한 소스 코드를 사용하여 다양한 데이터 스트림으로 구조화된 데이터를 쉽게 읽고 쓸 수 있다. 여러분은 또한 "옛날" 형식으로 컴파일된 배포 프로그램의 중단 없이 데이터 구조를 업데이트할 수 있다. 
(출처: https://m.blog.naver.com/oidoman/220773055827)

**HTTP/2**

HOL(Head Of Line) Blocking-특정응답지연, RTT(Round Trip TIme) 증가, 헤비한 Header구조라는 문제점들을 가지고 있습니다. 이러한 문제점들을 해결하기 위해, UI 개발자/프론트엔드개발자는 이미지 스프라이트, 도메인샤딩, CSS/JavaScript 압축, Data URI 등을 업무에 사용하였습니다.
그렇게 고군분투 하던 중, HTTP2가 세상에 소개되었습니다. HTTP2는 성능 뿐만 아니라 속도면에서도 월등한 녀석입니다. Multiplexed Streams(한 커넥션에 여러개의 메세지를 동시에 주고 받을 수 있음), Stream Prioritization(요청 리소스간 의존관계를 설정), Server Push(HTML문서상에 필요한 리소스를 클라이언트 요청없이 보내줄 수 있음), Header Compression(Header 정보를 HPACK압충방식을 이용하여 압축전송)을 사용하여 선을을 획기적으로 향상 시켰습니다. 
(출처: https://medium.com/@shlee1353/http1-1-vs-http2-0-%EC%B0%A8%EC%9D%B4%EC%A0%90-%EA%B0%84%EB%8B%A8%ED%9E%88-%EC%82%B4%ED%8E%B4%EB%B3%B4%EA%B8%B0-5727b7499b78)

## Fetures

- 낮은 데이터 payload: gRPC는 binary 기반으로 통신하기 때문에 효율성이 뛰어남
- Protocol buffer의 IDL로 REST API 대비 강한 타입 체크가 가능
- HTTP/2 지원 
- Server stream, Client stream, 양방향 stream 통신 방식을 지원 

## Structure

gRPC Server < gRPC Stub...(clients) 
            <- (Proto Request)
(Proto Response) ->
client와 server에 stub를 생성하고 정의된 IDL을 토대로 request/response를 생성하여 통신합니다.

## Protocol buffer IDL 정의 (.proto 파일 작성)

직렬화할 정보를 구조화하기 위해 .proto 파일에 protocol buffer message 타입을 정의한다. 각 protocol buffer message는 작은 논리적 정보 단위로서, 이름-값 쌍들을 포함하고 있다. 메시지 타입은 하나 이상의 유일한 숫자 필드를 가지고 있고, 각 필드는 이름과 값 타입을 가지고 있는데, 여기서 값 타입은 기본적인 타입 뿐만 아니라 다른 protocol buffer message 타입이 될 수 있다. 그리고 데이터를 계층적으로 구조화할 수도 있다. 

Protocol Buffers (이하 ProtoBuf) IDL 예제 코드

```
    // 헤더 부분 생략..
    // Greeter 서비스 정의
    service Greeter {
        // 서비스의 RPC 정의
        rpc SayHello (HelloRequest) returns (HelloReply) {}
    }
    // 요청 메시지 정의
    message HelloRequest {
        string name = 1;
    }
    // 응답 메시지 정의
    message HelloReply {
        string message = 1;
    }
```

ProtoBuf IDL 정의만으로 server 와 client(Stub) 코드가 자동으로 생성됩니다.
위 예제 코드에서 sayHello()라는 서비스의 비즈니스 로직과 클라이언트에서 Stub를 사용하여 호줄해주는 부분만 구현해주면 됩니다.



## Use it

- 단일 인스턴스로 돌아가는 CRUD 웹 애플리케이션에서부터 MSA 환경 등 거의 모든 서버 시스템 개발에 적합
- Authentication, Tracing, Load Balancing, Health Checking, API Gateway 등의 생태계가 있음

## Advantages

- ProtoBuf가 지원하는 IDL 활용한 서비스 및 메시지 정의는 MSA의 다양한 기술 스택의 공존으로인한 중복 발생의 단점을 보완하고, 수많은 서비스간의 API 호출로 인한 성능 저하를 개선합니다.
- ProtoBuf에 의한 높은 메시지 압축률은 시스템 전체의 네트워크 트래픽을 획기적으로 줄여줍니다. 이것은 동일한 자원 제약에서 더 많은 서비스 인스턴스를 띄울 수 있다는 것을 의미합니다.


(출처: https://medium.com/@goinhacker/microservices-with-grpc-d504133d191d)

## Setting
ProtoBuf 컴파일러 설치
go get -a github.com/golang/protobuf/protoc-gen-go




* reference
*https://malgogi-developer.tistory.com/39*