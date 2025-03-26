#!/usr/bin/expect -f

set timeout 10

# 서버 목록 파일 경로
set fp [open "../server-list/server-list.txt" r]

# SSH 공개 키 경로
set ssh_pub_key "../ssh-key/id_rsa.pub"

# 각 서버에 대해 반복 실행
while {[gets $fp line] != -1} {
    if {[regexp {(\S+)\s+(\S+)\s+(\S+)\s+(\S+)} $line -> ip user password port]} {

        puts "\n🔑 서버 ($user@$ip:$port)에 공개 키 등록 중..."

        # 1. 원격 서버에 .ssh 디렉토리 생성
        spawn ssh -p $port $user@$ip "mkdir -p ~/.ssh && chmod 700 ~/.ssh"
        expect {
            "yes/no" { send "yes\r"; exp_continue }
            "password:" { send "$password\r" }
        }
        expect eof

        # 2. 공개키 파일 전송
        spawn scp -P $port $ssh_pub_key $user@$ip:~/id_rsa.pub
        expect {
            "yes/no" { send "yes\r"; exp_continue }
            "password:" { send "$password\r" }
        }
        expect eof

        # 3. authorized_keys에 추가 및 권한 설정
        spawn ssh -p $port $user@$ip "cat ~/id_rsa.pub >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys && rm ~/id_rsa.pub"
        expect {
            "yes/no" { send "yes\r"; exp_continue }
            "password:" { send "$password\r" }
        }
        expect eof

        puts "✅ $user@$ip:$port 등록 완료!"
    }
}

close $fp

puts "\n🎉 모든 서버에 공개 키 등록이 완료되었습니다."