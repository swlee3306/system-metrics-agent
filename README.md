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

의존성은 `go.mod`/`go.sum`을 기준으로 사용합니다. Go 기능 테스트는 아직 없으며, 공개 설정·스크립트 입력 검사는 `python3 -m unittest discover -s tests -v`로 실행할 수 있습니다.
빌드 성공과 기능·운영 검증은 별개입니다. `go run .`은 실제 호스트 메트릭 수집과 서버 실행을 시작하므로 실행 경로와 설정을 먼저 확인하세요.

## 공개 자료의 경계

Grafana 설정은 [공개 예시·운영 설정 구분 안내](grafana/README.md)를 먼저 확인하세요.
`grafana/datasource.yaml`의 비밀번호는 명시적인 placeholder로 교체했습니다. 실제 비밀값은 별도 비공개 설정에서 관리해야 하며, Grafana 연동은 미검증입니다. ARM64 배포 스크립트도 서버 주소·SSH 키 경로를 명시적으로 입력해야 실행됩니다.

`pkg/db/model/`에는 레거시 생성 모델이 포함되어 있습니다. 이를 새로운 제품 기능이나 본인 기여의 근거로 확대하지 않습니다.
실제 운영 DB, 호스트 정보, 자격증명은 예제나 이슈에 올리지 마세요.

현재 프로젝트는 학습·실험 상태입니다. 플랫폼·관측성 프로젝트는 [infra-orch-studio](https://github.com/swlee3306/infra-orch-studio)와 [network-collector](https://github.com/swlee3306/network-collector)를 참고할 수 있습니다.
