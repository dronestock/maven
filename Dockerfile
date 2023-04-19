ARG JAVA_HOME=/opt/openjdk
ARG MAVEN_HOME=/usr/share/maven/


FROM dockerproxy.com/library/openjdk:18-alpine AS lts

# 方便环境变量的设置
RUN mkdir --parents /opt/openjdk/lts; mv /usr/lib/jvm/java-*/* /opt/openjdk/lts


FROM dockerproxy.com/library/openjdk:19-alpine AS latest

# 方便环境变量的设置
RUN mkdir --parents /opt/openjdk/latest; mv /opt/openjdk-*/* /opt/openjdk/latest


FROM dockerproxy.com/library/maven:3.9.1 AS maven


FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.17.2 AS builder

ARG JAVA_HOME
ARG MAVEN_HOME


# 复制文件
COPY --from=lts ${JAVA_HOME} /docker/${JAVA_HOME}
COPY --from=latest ${JAVA_HOME} /docker/${JAVA_HOME}
COPY --from=latest /etc/ssl/certs/java/cacerts /docker/etc/ssl/certs/java/cacerts
COPY --from=maven ${MAVEN_HOME} /docker/${MAVEN_HOME}
COPY docker /docker
COPY maven /docker/usr/local/bin




# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.17.2


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
    # 解决找不到库的问题
    && LD_PATH=/etc/ld-musl-x86_64.path \
    && echo "/lib" >> ${LD_PATH} \
    && echo "/usr/lib" >> ${LD_PATH} \
    && echo "/usr/local/lib" >> ${LD_PATH} \
    && echo "${JAVA_HOME}/lib/default" >> ${LD_PATH} \
    && echo "${JAVA_HOME}/lib/server" >> ${LD_PATH} \
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


ARG JAVA_HOME
ARG MAVEN_HOME
# 配置环境变量，配置Java主目录和Maven主目录以及Java和Maven的快捷方式
ENV JAVA_OPTS ""
ENV JAVA /var/lib/java
ENV MAVEN_LOCAL_REPOSITORY ${JAVA}/maven

ENV PATH=${JAVA_HOME}/bin:${MAVEN_HOME}/bin:$PATH
ENV JAVA_LTS ${JAVA_HOME}/lts
ENV JAVA_LATEST ${JAVA_HOME}/latest
