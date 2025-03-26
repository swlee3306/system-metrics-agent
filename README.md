# System Info Collector

**System Info Collector**는 시스템의 CPU, 메모리, 디스크, 네트워크 정보를 수집하여 API 형식으로 제공하는 경량 에이전트입니다.  
이 프로그램은 서버 자원 상태를 모니터링하고, 데이터 분석 또는 성능 튜닝에 필요한 정보를 효율적으로 제공하도록 설계되었습니다.

---

## 주요 기능

- **CPU 정보**: 프로세서 이름, 코어 수, 클럭 속도, 캐시 크기 등을 수집.
- **메모리 정보**: 총 메모리, 사용 메모리, 사용 가능한 메모리, 스왑 메모리 등.
- **디스크 정보**: 디스크 사용량, 가용 공간, 마운트 정보, 읽기/쓰기 속도 등.
- **네트워크 정보**: 네트워크 인터페이스, 송수신 데이터, 연결 상태 등.
- **REST API 제공**: 수집한 데이터를 RESTful API를 통해 제공.
- **SSH key 등록 스크립트 제공**: 접속할 서버의 SSH public key data 등록하는 기능 제공.

---

## 설치 및 실행

### 1. **사전 요구 사항**
- **빌드 운영체제**: mac (ARM 기반) (x86 도 가능)
- **운영 체제**: Ubuntu 20.04 이상
- **언어 런타임**: Go 1.19 이상
- **필수 권한**: 시스템 자원 접근을 위한 관리자 권한

### 2. **설치**

1. **프로젝트 클론**
   ```bash
   git clone https://github.com/<your-username>/system-Info-collector.git
   cd system-Info-collector
   ```

2. **빌드**
   ```bash
   cd script/arm/
   ./build.sh (default: arm64)
    -- 만약 x86 으로 빌드를 원하면
      ㄴ GOARCH=amd64 ./build.sh 으로 실행

   ./service-active.sh (바이너리 파일 배포후, 서비스 등록)
   변경 해야 할 값
   GOARCH=${GOARCH:-arm64}  # 기본값을 arm64로 설정 (환경 변수가 없을 경우)
   SERVER_USER="{{설치할 서버 계정}}"       # 우분투 서버의 사용자 계정
   SERVER_IP="{{설치할 서버 IP}}" # 우분투 서버의 IP 주소
   SERVER_PORT="{{설치할 서버 Port}}"       # 우분투 서버의 SSH 포트
   TARGET_DIR="/usr/local/bin"
   SERVICE_NAME="system-agent"
   SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"

   ./service-deactive.sh (서비스 삭제)
   변경 해야 할 값
   SERVER_USER="{{설치할 서버 계정}}"       # 우분투 서버의 사용자 계정
   SERVER_IP="{{설치할 서버 IP}}" # 우분투 서버의 IP 주소
   SERVER_PORT="{{설치할 서버 Port}}"       # 우분투 서버의 SSH 포트
   TARGET_DIR="/usr/local/bin"
   SERVICE_NAME="system-agent"
   SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"   
   ```

3. **실행 후 API**
   - **http://{{server ip}}:8080/api/v1.0/cpuinfo**
