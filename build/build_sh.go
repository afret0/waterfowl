package build

func BuildSh_tem() string {
	t := `
	#!/bin/bash
	set -e
	
	appid=$1
	port=$2
	env=$3
	#编译可执行文件
	GOOS=linux GOARCH=amd64 go build -o 'bin/server' ./main.go
	
	docker build -f Dockerfile --build-arg HttpPort=${port} -t registry.cn-hangzhou.aliyuncs.com/kiwi0325/${appid}:${env} .
	docker push registry.cn-hangzhou.aliyuncs.com/kiwi0325/${appid}:${env}
	docker rmi registry.cn-hangzhou.aliyuncs.com/kiwi0325/${appid}:${env}
	
	if [ "$env" = "test"  ]; then
	  echo "测试环境"
	else
	  echo "线上环境"
	fi
	
	`
	return t
}
