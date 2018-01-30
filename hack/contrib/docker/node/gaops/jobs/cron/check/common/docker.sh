#!/bin/bash

OS_VERSION=${1:-centos7}

function log.info() {
  echo "       $*"
}
 
function log.error() {
  echo " !     $*"
  echo ""
}
 
function log.stdout() {
    echo "$*" >&2
}

function prepare() {
    log.info "checking docker service..."
}
#区分系统版本,centos与debian系列
function is_install() {
    pkgname=$1
    log.info "checking $pkgname install..."
    if [[ $OS_VERSION =~ "7" ]];then
        rpm -qi $pkgname > /dev/null
        if [ $? -eq 0 ];then
            log.info "service $pkgname is installed"
        return 0
        else
            log.error "service $pkgname is not install"
        return 1
        fi
    else
        apt search $pkgname | grep installed >/dev/null
        if [ $? -eq 0 ];then
            log.info "service $pkgname is installed"
        return 0
        else
            apt search $pkgname | grep upgradable >/dev/null
            if [ $? -eq 0 ];then
                log.info "service $pkgname is installed"
            return 0
            else
                log.error "service $pkgname is not install"
            return 1
            fi
        fi
    fi
}

function is_enable() {
    UNIT=$1
    log.info "checking $UNIT enabled..."
        check_enable=$(systemctl is-enabled $UNIT)
        if [ $check_enable = "enabled" ];then
            log.info "$UNIT is enable"
        return 0
        else
            log.error "$UNIT is not enable"
            systemctl enable $UNIT
        return 0
        fi
}

#循环三次定时检测，防止restart
function is_active() {
    UNIT=$1
    log.info "checking $UNIT active..."
    for (( i=1; i <= 3; i++ ))
    do
    sleep 3
    check_active=$(systemctl is-active $UNIT)
        if [ $check_active = "active" ];then 
            log.info "$UNIT is running"
        else
            log.error "$UNIT is not running"
        return 1
        fi
    done
    return 0
}
#使用docker命令查看可用，并设置超时范围
function is_normal() {
    log.info "checking docker's health..."
    timeout 5 docker ps > /dev/null
    if [ $? -eq 0 ];then
    log.info "service docker is healthy"
    return 0
    else
    log.error "service docker is unhealthy"
    return 1
    fi
}

function run() {
    is_install gr-docker-engine
    if [ "$?" -eq 0 ];then
        is_enable docker
        if [ "$?" -eq 0 ];then
            is_active docker 
            if [ "$?" -eq 0 ];then
                is_normal
                if [ "$?" -eq 0 ];then
                log.stdout '{
                "status":[
                {
                    "name":"check-service-docker",
                    "condition_type":"SERVICE_DOCKER_NORMAL",
                    "condition_status":"True"
                }
                ],
                "exec_status":"Success",
                "type":"check"
                }'

                elif [ "$?" -eq 1 ];then
                log.stdout '{
                "status":[
                {
                    "name":"check-docker-normal",
                    "condition_type":"DOCKER_IS_ABNORMAL", 
                    "condition_status":"False" 
                }
                ],
                "exec_status":"Success", 
                "type":"check"
                }'
                fi
            elif [ "$?" -eq 1 ];then
            log.stdout '{
            "status":[
            {
                "name":"check-docker-active",
                "condition_type":"DOCKER_ISNOT_ACTIVE", 
                "condition_status":"False" 
            }
            ],
            "exec_status":"Success", 
            "type":"check"
            }'
            fi
        elif [ "$?" -eq 1 ];then
        log.stdout '{
        "status":[
        {
            "name":"check-docker-enable",
            "condition_type":"DOCKER_ISNOT_ENABLE", 
            "condition_status":"False" 
        }
        ],
        "exec_status":"Success", 
        "type":"check"
        }'
        systemctl enable docker.service
        fi
    elif [ "$?" -eq 1 ];then
    log.stdout '{
    "status":[
    {
        "name":"check-docker-installed",
        "condition_type":"DOCKER_ISNOT_INSTALL", 
        "condition_status":"False" 
    }
    ],
    "exec_status":"Success", 
    "type":"check"
    }'
    fi 
}

case $1 in
    * )
        prepare
        run
    ;;
esac