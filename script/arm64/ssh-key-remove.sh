#!/bin/bash

# 서버 계정명과 IP를 입력받음
SERVER_USER=${1:-"root"}
SERVER_IP=${2}

# SSH 공개 키 경로
SSH_PUB_KEY="../ssh-key/id_rsa.pub"

# 서버 IP가 입력되지 않은 경우 오류 메시지 출력
if [[ -z "$SERVER_IP" ]]; then
    echo "❌ 사용법: ./ssh-key-remove.sh [계정명] [서버 IP]"
    echo "예시: ./ssh-key-setup.sh abc x.x.x.x"
    exit 1
fi

# 공개 키 내용 가져오기
PUB_KEY_CONTENT=$(cat "$SSH_PUB_KEY")

# 서버에서 공개 키 제거
echo "🗑 서버 ($SERVER_USER@$SERVER_IP)에서 공개 키 제거 중..."
ssh -p 22 "$SERVER_USER@$SERVER_IP" "sed -i '/$(echo $PUB_KEY_CONTENT | sed 's/[\/&]/\\&/g')/d' ~/.ssh/authorized_keys"

# 완료 메시지
echo "✅ 공개 키 삭제 완료! 이제 SSH 키 기반 인증이 불가능합니다."
