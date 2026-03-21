# System-Info-Collector

## 한 줄 소개
시스템 메트릭을 수집해 API 형태로 제공하는 경량 에이전트입니다.

## 저장소 성격
- 분류: 인프라 / 에이전트
- 목적: CPU, 메모리, 디스크, 네트워크 정보 수집
- 핵심 기술: Go, 시스템 메트릭, API, 모니터링

시스템 정보를 수집하고 외부에서 활용할 수 있도록 정리하는 실험 저장소입니다.

시스템의 CPU, 메모리, 디스크, 네트워크 정보를 수집하여 API 형식으로 제공하는 경량 에이전트입니다. 시각화, 알림, 데이터 저장, 확장성 같은 요소를 함께 실험하고 정리하고 있습니다.

## 🚀 주요 기능

### 📊 시각화 및 대시보드
- **Grafana 대시보드**: 실시간 모니터링 차트 및 위젯
- **커스텀 위젯**: 시스템별 맞춤형 모니터링 위젯
- **대시보드 템플릿**: 다양한 환경용 대시보드 템플릿
- **모바일 반응형**: 모바일 기기에서도 최적화된 대시보드

### 🚨 알림 및 경고
- **임계값 기반 알림**: CPU, 메모리, 디스크 사용률 임계값 설정
- **다중 알림 채널**: 이메일, Slack, PagerDuty 연동
- **알림 규칙 엔진**: 복잡한 조건 기반 알림 규칙
- **알림 히스토리**: 알림 발생 이력 및 해결 상태 추적

### 💾 데이터 저장 및 분석
- **시계열 데이터베이스**: InfluxDB, TimescaleDB 지원
- **데이터 보존 정책**: 자동 데이터 정리 및 압축
- **성능 분석**: 장기간 성능 트렌드 분석
- **데이터 내보내기**: CSV, JSON 형식 데이터 내보내기

### 🔒 보안 및 인증
- **JWT 기반 인증**: 안전한 API 접근 제어
- **역할 기반 권한**: Admin, User, Guest 권한 관리
- **API 키 관리**: 서비스 간 인증을 위한 API 키
- **감사 로그**: 모든 접근 및 작업 로그 기록

### 📈 확장성 및 성능
- **중앙화된 에이전트 관리**: 다중 서버 에이전트 통합 관리
- **수평적 확장**: 로드 밸런싱 및 클러스터링 지원
- **성능 최적화**: 효율적인 데이터 수집 및 처리
- **리소스 모니터링**: 에이전트 자체 성능 모니터링

### 🛠️ 개발 및 배포
- **크로스 플랫폼 빌드**: ARM64, AMD64 지원
- **자동화된 배포**: 스크립트 기반 서비스 관리
- **SSH 키 관리**: 원격 서버 자동 설정
- **서비스 관리**: systemd 기반 서비스 제어
- **데이터베이스 모델**: 자동 생성된 GORM 모델

## 📦 설치 및 실행

### 1. 저장소 클론
```bash
git clone https://github.com/swlee3306/system-Info-collector.git
cd system-Info-collector
```

### 2. 의존성 설치
```bash
go mod tidy
```

### 3. 환경 변수 설정
```bash
# .env 파일 생성
cp .env.example .env

# 데이터베이스 설정
export INFLUXDB_URL=http://localhost:8086
export INFLUXDB_TOKEN=your-token
export INFLUXDB_ORG=your-org
export INFLUXDB_BUCKET=system-metrics

# Grafana 설정
export GRAFANA_URL=http://localhost:3000
export GRAFANA_API_KEY=your-api-key

# 알림 설정
export ALERT_EMAIL_SMTP=smtp.gmail.com:587
export ALERT_EMAIL_USER=your-email@gmail.com
export ALERT_EMAIL_PASS=your-password
```

### 4. 서비스 실행
```bash
# 개발 모드
go run main.go

# 프로덕션 모드
go build -o system-collector main.go
./system-collector
```

### 5. 크로스 플랫폼 빌드
```bash
# ARM64 빌드 (Linux)
cd script/arm64
GOARCH=arm64 ./build.sh

# AMD64 빌드 (Linux)
cd script/arm64
GOARCH=amd64 ./build.sh

# 빌드된 바이너리 실행
./system-agent-linux-arm64
```

### 6. 서비스 관리
```bash
# 서비스 활성화
./service-active.sh

# 서비스 비활성화
./service-deactive.sh

# SSH 키 설정 (원격 서버용)
./ssh-key-setup.sh
```

## 🏗️ 프로젝트 구조

```
system-Info-collector/
├── main.go                 # 메인 실행 파일
├── cmd/                    # 서버 시작 로직
│   └── server.go          # 시스템 에이전트 시작
├── internal/               # 내부 패키지
│   └── collector/         # 데이터 수집기
│       ├── cpu/           # CPU 정보 수집
│       │   └── cpu.go
│       ├── memory/        # 메모리 정보 수집
│       │   └── memory.go
│       ├── disk/          # 디스크 정보 수집
│       │   └── disk.go
│       ├── network/       # 네트워크 정보 수집
│       │   └── network.go
│       └── exporter/      # 메트릭 익스포터
│           └── metric_exporter.go
├── pkg/                    # 공통 패키지
│   ├── config/           # 설정 관리
│   │   ├── config_cpu.go
│   │   ├── config_disk.go
│   │   ├── config_memory.go
│   │   └── config_network.go
│   ├── db/               # 데이터베이스 모델
│   │   └── model/        # 생성된 모델 파일들
│   ├── logger/           # 로깅 시스템
│   │   └── logger.go
│   ├── prometheus/       # Prometheus 메트릭
│   │   └── prometheus.go
│   └── web/              # 웹 서버
│       └── router.go
├── storage/               # 데이터 저장
│   └── timeseries.go      # 시계열 데이터베이스 연동
├── alerting/              # 알림 시스템
│   └── alerts.go          # 알림 규칙 및 처리
├── auth/                  # 인증 시스템
│   └── auth.go            # JWT 인증 및 권한 관리
├── grafana/               # Grafana 연동
│   ├── dashboard.json     # 대시보드 설정
│   └── datasource.yaml    # 데이터소스 설정
├── script/                # 빌드 및 배포 스크립트
│   └── arm64/            # ARM64 빌드 스크립트
│       ├── build.sh
│       ├── service-active.sh
│       ├── service-deactive.sh
│       ├── ssh-key-setup.sh
│       └── system-agent-linux-arm64
├── go.mod                # Go 모듈 파일
├── go.sum                # 의존성 체크섬
└── README.md             # 프로젝트 문서
```

## 🔧 API 엔드포인트

### 시스템 정보 조회
```http
# CPU 정보
GET /api/v1/cpu
Authorization: Bearer <token>

# 메모리 정보
GET /api/v1/memory
Authorization: Bearer <token>

# 디스크 정보
GET /api/v1/disk
Authorization: Bearer <token>

# 네트워크 정보
GET /api/v1/network
Authorization: Bearer <token>

# 전체 시스템 정보
GET /api/v1/system
Authorization: Bearer <token>
```

### 메트릭 및 헬스체크
```http
# Prometheus 메트릭
GET /metrics

# 헬스체크
GET /health

# 시스템 상태
GET /api/v1/status
Authorization: Bearer <token>
```

### 알림 관리
```http
# 알림 규칙 생성
POST /api/v1/alerts/rules
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "High CPU Usage",
  "condition": "cpu_usage > 80",
  "duration": "5m",
  "severity": "warning",
  "channels": ["email", "slack"]
}

# 알림 규칙 목록
GET /api/v1/alerts/rules
Authorization: Bearer <token>

# 알림 히스토리
GET /api/v1/alerts/history
Authorization: Bearer <token>
```

## 📊 Grafana 대시보드

### 대시보드 설정
```json
{
  "dashboard": {
    "title": "System Info Collector Dashboard",
    "panels": [
      {
        "title": "CPU Usage",
        "type": "timeseries",
        "targets": [
          {
            "expr": "node_cpu_usage_percent",
            "legendFormat": "CPU Usage %"
          }
        ]
      },
      {
        "title": "Memory Usage",
        "type": "timeseries",
        "targets": [
          {
            "expr": "node_memory_usage_percent",
            "legendFormat": "Memory Usage %"
          }
        ]
      }
    ]
  }
}
```

### 데이터소스 설정
```yaml
apiVersion: 1

datasources:
  - name: Prometheus
    type: prometheus
    uid: prometheus
    access: proxy
    url: http://localhost:9090
    isDefault: true
```

## 🚨 알림 시스템

### 알림 규칙 설정
```go
// CPU 사용률 알림
alert := alerting.AlertCondition{
    Metric:    "cpu_usage",
    Threshold: 80.0,
    Operator:  ">",
    Duration:  5 * time.Minute,
    Severity:  "warning",
    Message:   "High CPU usage detected!",
}

// 메모리 사용률 알림
alert := alerting.AlertCondition{
    Metric:    "memory_usage",
    Threshold: 90.0,
    Operator:  ">",
    Duration:  2 * time.Minute,
    Severity:  "critical",
    Message:   "Memory usage approaching limit!",
}
```

### 알림 채널 설정
```go
// 이메일 알림
emailChannel := alerting.NotificationChannel{
    Type: "email",
    Config: map[string]string{
        "smtp_host": "smtp.gmail.com",
        "smtp_port": "587",
        "username":  "your-email@gmail.com",
        "password":  "your-password",
        "to":        "admin@example.com",
    },
}

// Slack 알림
slackChannel := alerting.NotificationChannel{
    Type: "slack",
    Config: map[string]string{
        "webhook_url": "https://hooks.slack.com/services/...",
        "channel":     "#alerts",
    },
}
```

## 💾 데이터 저장

### InfluxDB 연동
```go
// InfluxDB 클라이언트 설정
client, err := storage.NewInfluxDBClient(
    "http://localhost:8086",
    "your-token",
    "your-org",
    "system-metrics",
)
if err != nil {
    log.Fatal("Failed to create InfluxDB client:", err)
}
defer client.Close()

// 데이터 저장
data := storage.TimeSeriesData{
    Measurement: "cpu_usage",
    Tags: map[string]string{
        "host": "server01",
        "region": "us-west",
    },
    Fields: map[string]interface{}{
        "value": 75.5,
    },
    Timestamp: time.Now(),
}

err = client.Write(data)
if err != nil {
    log.Printf("Failed to write data: %v", err)
}
```

### 데이터 쿼리
```go
// Flux 쿼리 실행
fluxQuery := `
from(bucket: "system-metrics")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "cpu_usage")
  |> filter(fn: (r) => r.host == "server01")
`

records, err := client.Query(fluxQuery)
if err != nil {
    log.Printf("Failed to query data: %v", err)
} else {
    log.Printf("Query results: %+v", records)
}
```

## 🔐 보안 및 인증

### JWT 토큰 생성
```go
// JWT 토큰 생성
token, err := auth.GenerateToken("user123", []string{"admin"})
if err != nil {
    log.Printf("Failed to generate token: %v", err)
}

// 토큰 검증
claims, err := auth.VerifyToken(token)
if err != nil {
    log.Printf("Invalid token: %v", err)
}
```

### 역할 기반 접근 제어
```go
// 권한 확인
if auth.HasRole(claims, "admin") {
    // 관리자 권한 작업
}

// API 키 인증
apiKey := c.GetHeader("X-API-Key")
if !auth.ValidateAPIKey(apiKey) {
    c.JSON(401, gin.H{"error": "Invalid API key"})
    return
}
```

## 🧪 테스트

### 단위 테스트 실행
```bash
# 모든 테스트 실행
go test ./...

# 특정 패키지 테스트
go test ./collectors
go test ./storage
go test ./alerting

# 테스트 커버리지 확인
go test -cover ./...
```

### 통합 테스트
```bash
# 데이터베이스 통합 테스트
go test ./storage -v

# 알림 시스템 테스트
go test ./alerting -v
```

## 📈 성능 및 모니터링

### 메트릭 수집
```go
// Prometheus 메트릭 설정
metrics.InitMetrics()
metrics.StartMetricsServer("9090")

// 커스텀 메트릭 추가
metrics.SystemInfoCounter.WithLabelValues("cpu", "success").Inc()
metrics.SystemInfoDuration.WithLabelValues("cpu").Observe(duration.Seconds())
```

### 로깅 설정
```go
// 구조화된 로깅
logger.WithFields(logrus.Fields{
    "component": "collector",
    "metric":    "cpu",
    "value":     cpuUsage,
}).Info("System metric collected")
```

## 🚀 배포 및 운영

### Docker 사용
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o system-collector main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/system-collector .
COPY --from=builder /app/.env .
CMD ["./system-collector"]
```

### Docker Compose
```yaml
version: '3.8'
services:
  system-collector:
    build: .
    ports:
      - "8080:8080"
      - "9090:9090"
    environment:
      - INFLUXDB_URL=http://influxdb:8086
      - GRAFANA_URL=http://grafana:3000
    depends_on:
      - influxdb
      - grafana
  
  influxdb:
    image: influxdb:2.0
    ports:
      - "8086:8086"
    environment:
      - DOCKER_INFLUXDB_INIT_MODE=setup
      - DOCKER_INFLUXDB_INIT_USERNAME=admin
      - DOCKER_INFLUXDB_INIT_PASSWORD=password
      - DOCKER_INFLUXDB_INIT_ORG=myorg
      - DOCKER_INFLUXDB_INIT_BUCKET=system-metrics
  
  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
```

## 🔧 환경 변수

### 필수 환경 변수
```bash
# 서버 설정
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# 데이터베이스 설정
INFLUXDB_URL=http://localhost:8086
INFLUXDB_TOKEN=your-token
INFLUXDB_ORG=your-org
INFLUXDB_BUCKET=system-metrics

# Grafana 설정
GRAFANA_URL=http://localhost:3000
GRAFANA_API_KEY=your-api-key
```

### 선택적 환경 변수
```bash
# 알림 설정
ALERT_EMAIL_SMTP=smtp.gmail.com:587
ALERT_EMAIL_USER=your-email@gmail.com
ALERT_EMAIL_PASS=your-password

# 로깅 설정
LOG_LEVEL=info
LOG_FORMAT=json
```

## 🤝 기여하기

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 라이센스

이 프로젝트는 MIT 라이센스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

## 🙏 감사의 말

- [Go](https://golang.org/) - 프로그래밍 언어
- [Gin](https://gin-gonic.com/) - HTTP 웹 프레임워크
- [InfluxDB](https://www.influxdata.com/) - 시계열 데이터베이스
- [Grafana](https://grafana.com/) - 시각화 플랫폼
- [Prometheus](https://prometheus.io/) - 메트릭 수집

## 📞 지원 및 문의

- 이슈 리포트: [GitHub Issues](https://github.com/swlee3306/system-Info-collector/issues)
- 이메일: swlee3306@gmail.com
- 문서: [Wiki](https://github.com/swlee3306/system-Info-collector/wiki)

---

**System-Info-Collector** - 엔터프라이즈급 시스템 모니터링 솔루션 🚀
