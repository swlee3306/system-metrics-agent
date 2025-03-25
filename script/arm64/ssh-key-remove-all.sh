#!/usr/bin/expect -f

set timeout 10

# 서버 목록 파일 경로
set fp [open "../server-list/server-list.txt" r]

# SSH 공개 키 경로
set ssh_pub_key "../ssh-key/id_rsa.pub"

# 공개 키 내용 로드
set f [open $ssh_pub_key r]
set pubkey [read $f]
close $f

# 각 서버에 대해 반복 실행
while {[gets $fp line] != -1} {
    if {[regexp {(\S+)\s+(\S+)\s+(\S+)\s+(\S+)} $line -> ip user password port]} {
        # ssh-key-remove.sh 스크립트 실행하여 공개 키 제거
        puts "\n🗑 서버 ($user@$ip:$port)에서 공개 키 제거 중..."
        spawn bash ./ssh-key-remove.sh $user $ip
        expect {
            "yes/no" { send "yes\r"; exp_continue }
            "password:" { send "$password\r" }
        }
        expect eof

        puts "✅ $user@$ip:$port 공개 키 삭제 완료!"
    }
}

close $fp

puts "\n🧹 모든 서버에서 공개 키 삭제가 완료되었습니다."
