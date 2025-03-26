#!/bin/bash

# 서버 계정명과 IP를 입력받음
SERVER_USER=${1:-"root"}
SERVER_IP=${2}

# SSH 공개 키 경로
SSH_PUB_KEY="../ssh-key/id_rsa.pub"

# 서버 IP가 입력되지 않은 경우 오류 메시지 출력
if [[ -z "$SERVER_IP" ]]; then
    echo "❌ 사용법: ./ssh-key-setup.sh [계정명] [서버 IP]"
    echo "예시: ./ssh-key-setup.sh abc x.x.x.x"
    exit 1
fi

# 공개 키 등록 실행
echo "🔑 서버 ($SERVER_USER@$SERVER_IP)에 공개 키 등록 중..."
ssh-copy-id -i "$SSH_PUB_KEY" -p 22 "$SERVER_USER@$SERVER_IP"

# 완료 메시지
echo "✅ 공개 키 등록 완료! 이제 SSH 키 기반 인증을 사용할 수 있습니다."