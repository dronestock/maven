ARG MAVEN_HOME=/usr/share/maven/


FROM dockerproxy.com/library/maven:3.9.6 AS maven
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.0 AS builder

ARG MAVEN_HOME

# 复制文件
COPY --from=maven ${MAVEN_HOME} /docker/${MAVEN_HOME}
COPY docker /docker
COPY maven /docker/usr/local/bin



# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/java:0.0.3


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Maven插件，支持测试、打包、发布等常规功能"


# 复制文件
COPY --from=builder /docker /


RUN set -ex \
    \
    \
    \
    # 安装依赖库
    && apk update \
    && apk --no-cache add libstdc++ gcompat gnupg \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/maven \
    && chmod +x /usr/local/bin/gsk \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /usr/local/bin/maven


ARG MAVEN_HOME
ENV MAVEN_LOCAL_REPOSITORY ${JAVA_LIB}/maven
ENV PATH=${MAVEN_HOME}/bin:$PATH
