# system-metrics-agent

Go 기반 시스템 정보 수집·API 실험 프로젝트입니다. 기본 실행 경로와 추가 실험 패키지를 구분해 살펴보세요.

## 실행 경로와 범위

`main.go` → `cmd/StartSystemAgent` → metric exporter + web server가 기본 경로입니다.
CPU·메모리·디스크·네트워크 수집 코드는 `internal/collector/`에 있습니다.

`auth/`, `alerting/`, `storage/`에는 인증·알림·시계열 저장 관련 구현이 있으나, 폴더가 존재한다는 이유만으로 기본 실행 경로에 모두 연결됐거나 운영 검증됐다고 보지 않습니다.
Grafana, 이메일/Slack/PagerDuty, 다중 서버 관리, 고가용성·처리량 보장은 이 저장소의 검증된 성과로 주장하지 않습니다.

## 코드 확인

```sh
go build -mod=readonly ./...
```

의존성은 `go.mod`/`go.sum`을 기준으로 사용합니다. 자동화된 테스트 파일은 현재 없습니다.
빌드 성공과 기능·운영 검증은 별개입니다. `go run .`은 실제 호스트 메트릭 수집과 서버 실행을 시작하므로 실행 경로와 설정을 먼저 확인하세요.

## 공개 자료의 경계

Grafana 설정은 [공개 예시·운영 설정 구분 안내](grafana/README.md)를 먼저 확인하세요.
기존 `grafana/datasource.yaml`은 자격 증명 후보가 남아 있고 운영 사용 여부가 미확인이므로 그대로 적용할 수 있는 검증된 예시가 아닙니다.

`pkg/db/model/`에는 레거시 생성 모델이 포함되어 있습니다. 이를 새로운 제품 기능이나 본인 기여의 근거로 확대하지 않습니다.
실제 운영 DB, 호스트 정보, 자격증명은 예제나 이슈에 올리지 마세요.

현재 프로젝트는 학습·실험 상태입니다. 플랫폼·관측성 프로젝트는 [infra-orch-studio](https://github.com/swlee3306/infra-orch-studio)와 [network-collector](https://github.com/swlee3306/network-collector)를 참고할 수 있습니다.
