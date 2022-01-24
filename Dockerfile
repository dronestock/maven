FROM storezhang/alpine AS jdk

# 明确指定工作目录，防止后面运行命令出现文件或者目录找不到的问题
WORKDIR /opt


RUN apk update
RUN apk add axel

# 安装AdoptOpenJDK，替代Oracle JDK
ARG JDK_MAJOR_VERSION=17
ARG JDK_MINOR_VERSION=0
ARG JDK_PATCH_VERSION=1+12
ARG JAVA_HOME=/usr/lib/jvm/java-${JDK_MAJOR_VERSION}-adoptopenjdk-amd64
ENV JDK_VERSION ${JDK_MAJOR_VERSION}.${JDK_MINOR_VERSION}.${JDK_PATCH_VERSION}
ENV OPENJ9_VERSION 0.29.1
ENV JDK_BIN_FILENAME jre${JDK_VERSION}.tar.gz
ENV JDK_DOWNLOAD_URL https://ghproxy.com/https://github.com/ibmruntimes/semeru${JDK_MAJOR_VERSION}-binaries/releases/download/jdk-${JDK_VERSION}_openj9-${OPENJ9_VERSION}/ibm-semeru-open-jdk_x64_linux_${JDK_VERSION/+/_}_openj9-${OPENJ9_VERSION}.tar.gz

RUN axel --num-connections 16 --output ${JDK_BIN_FILENAME} --insecure ${JDK_DOWNLOAD_URL}
RUN tar -xzf ${JDK_BIN_FILENAME}
RUN mkdir -p /usr/lib/jvm/java-${JDK_MAJOR_VERSION}-adoptopenjdk-amd64
RUN mv jdk-${JDK_VERSION}+9-jre/* ${JAVA_HOME}/



FROM storezhang/alpine AS maven

# 明确指定工作目录，防止后面运行命令出现文件或者目录找不到的问题
WORKDIR /opt


RUN apk update
RUN apk add axel


# 安装Maven
ARG MAVEN_MAJOR_VERSION=3
ARG MAVEN_MINOR_VERSION=8
ARG MAVEN_PATCH_VERSION=4
ENV MAVEN_VERSION=${MAVEN_MAJOR_VERSION}.${MAVEN_MINOR_VERSION}.${MAVEN_PATCH_VERSION}
ARG MAVEN_HOME=/opt/apache/maven${MAVEN_VERSION}
ENV MAVEN_FULL_NAME=apache-maven-${MAVEN_VERSION}
ENV MAVEN_BIN_FILENAME=${MAVEN_FULL_NAME}.tar.gz
ENV MAVEN_DOWNLOAD_URL=https://dlcdn.apache.org/maven/maven-${MAVEN_MAJOR_VERSION}/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz

RUN axel --num-connections 16 --output ${MAVEN_BIN_FILENAME} --insecure ${MAVEN_DOWNLOAD_URL}
RUN tar -xzf ${MAVEN_BIN_FILENAME}
RUN mkdir -p ${MAVEN_HOME}
RUN mv ${MAVEN_FULL_NAME}/* ${MAVEN_HOME}/





# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


ARG JAVA_HOME
ARG MAVEN_HOME


# 复制文件
COPY --from=jdk ${JAVA_HOME} ${JAVA_HOME}
COPY --from=maven ${MAVEN_HOME} ${MAVEN_HOME}
COPY maven /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/maven \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/maven


# 配置环境变量，配置Java主目录和Maven主目录以及Java和Maven的快捷方式
ENV JAVA_HOME ${JAVA_HOME}
ENV JAVA_OPTS ""
ENV MAVEN_HEOM ${MAVEN_HOME}

ENV PATH=${JAVA_HOME}/bin:${MAVEN_HOME}/bin:$PATH
