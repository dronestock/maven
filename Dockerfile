FROM openjdk:19-alpine AS jdk

# 方便环境变量的设置
RUN mv /opt/openjdk-* /opt/openjdk



FROM maven:3.8.6 AS maven




# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成Maven插件，支持测试、打包、发布等常规功能"


ENV JAVA_HOME /opt/openjdk
ENV MAVEN_HOME /usr/share/maven/


# 复制文件
COPY --from=jdk ${JAVA_HOME} ${JAVA_HOME}
COPY --from=maven ${MAVEN_HOME} ${MAVEN_HOME}
COPY docker /
COPY maven /bin


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
    && chmod +x /bin/maven \
    && chmod +x /usr/bin/gsk \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/maven


# 配置环境变量，配置Java主目录和Maven主目录以及Java和Maven的快捷方式
ENV JAVA_OPTS ""
ENV JAVA /var/lib/java
ENV MAVEN_LOCAL_REPOSITORY ${JAVA}/maven

ENV PATH=${JAVA_HOME}/bin:${MAVEN_HOME}/bin:$PATH
