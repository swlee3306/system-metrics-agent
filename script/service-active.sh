#!/bin/bash

# 환경 설정
GOARCH=${GOARCH:-arm64}  # 기본값을 arm64로 설정 (환경 변수가 없을 경우)
SERVER_USER="root"       # 우분투 서버의 사용자 계정
SERVER_IP="192.168.65.2" # 우분투 서버의 IP 주소
SERVER_PORT="5000"       # 우분투 서버의 SSH 포트
TARGET_DIR="/usr/local/bin"
SERVICE_NAME="system-agent"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"

# 1. 빌드된 파일을 우분투 서버로 전송 (SSH 포트 지정)
echo "📤 서버로 바이너리 업로드..."
scp -P $SERVER_PORT system-agent-linux-$GOARCH $SERVER_USER@$SERVER_IP:$TARGET_DIR/system-agent

# 2. 실행 권한 부여 (SSH 포트 지정)
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "chmod +x $TARGET_DIR/system-agent"

# 3. 서비스 파일 생성 (SSH 포트 지정)
echo "📝 서비스 파일 생성..."
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "echo \"[Unit]
Description=System Agent Service
After=network.target

[Service]
ExecStart=$TARGET_DIR/system-agent
Restart=always
User=root
Group=root

[Install]
WantedBy=multi-user.target
\" | sudo tee $SERVICE_FILE"

# 4. 서비스 등록 및 실행 (SSH 포트 지정)
echo "🚀 서비스 등록 및 실행..."
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl daemon-reload"
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl enable $SERVICE_NAME"
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl restart $SERVICE_NAME"

# 5. 서비스 상태 확인 (SSH 포트 지정)
echo "🔍 서비스 상태 확인..."
ssh -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl status $SERVICE_NAME --no-pager"

echo "✅ 배포 완료!"
