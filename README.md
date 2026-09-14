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

의존성은 `go.mod`/`go.sum`을 기준으로 사용합니다. 로거의 레벨·메시지·호출 위치·동시 호출·의존 경계는 `go test -race -count=1 ./pkg/logger`로 검사합니다. 수집기 전체 기능 테스트를 의미하지 않습니다. 공개 설정·스크립트 입력 검사는 `python3 -m unittest discover -s tests -v`로 실행할 수 있습니다.
빌드 성공과 기능·운영 검증은 별개입니다. `go run .`은 실제 호스트 메트릭 수집과 서버 실행을 시작하므로 실행 경로와 설정을 먼저 확인하세요.

## 공개 자료의 경계

Grafana 설정은 [공개 예시·운영 설정 구분 안내](grafana/README.md)를 먼저 확인하세요.
`grafana/datasource.yaml`은 `secureJsonData` 아래 명시적인 placeholder를 사용하는 구조 예시입니다. 실제 비밀값은 별도 비공개 설정에서 관리해야 합니다. 격리 Grafana 13.2.1에서 설정 등록과 비밀 필드 처리를 검증했으며 실제 데이터소스 연동은 미검증입니다. ARM64 배포 스크립트도 서버 주소·SSH 키 경로를 명시적으로 입력해야 실행됩니다.

현재 소스에서 레거시 생성 DB 모델과 로거의 DB 의존성을 제거했습니다. 로그는 Go 표준 로거(기본 stderr)로 즉시 출력하며, 호출 위치에는 파일명만 포함합니다. 메시지 자체의 비밀값은 자동으로 정제하지 않습니다.

이전 `LogSetup`·`LogEntries` API와 DB 저장·CSV 보관·자동 디렉터리 삭제는 더 이상 제공하지 않습니다. 외부에서 해당 API를 사용하는 경우 호환되지 않습니다. `Debug/Info/Warn/Error/Fatal` 호출 형식은 유지하며, `Fatal`은 기존처럼 프로세스를 종료하지 않습니다.

최신 트리에서는 기존 추적 실행 바이너리를 제외했습니다. 실행이 필요하면 검토한 소스에서 직접 빌드하세요. 과거 이력·다른 브랜치·PR·이전 바이너리에는 옛 코드가 남을 수 있으므로 이 변경을 전체 기록 정리 완료로 해석하지 마세요.
실제 운영 DB, 호스트 정보, 자격증명은 예제나 이슈에 올리지 마세요.

현재 프로젝트는 학습·실험 상태입니다. 플랫폼·관측성 프로젝트는 [infra-orch-studio](https://github.com/swlee3306/infra-orch-studio)와 [network-collector](https://github.com/swlee3306/network-collector)를 참고할 수 있습니다.
