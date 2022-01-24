FROM storezhang/alpine AS maven


ARG JRE_MAJOR_VERSION=11
ARG JRE_MINOR_VERSION=0
ARG JRE_PATCH_VERSION=11
ENV JRE_VERSION=${JRE_MAJOR_VERSION}.${JRE_MINOR_VERSION}.${JRE_PATCH_VERSION}
ENV OPENJ9_VERSION 0.26.0
ENV JRE_DOWNLOAD_URL=https://ghproxy.com/https://github.com/AdoptOpenJDK/openjdk${JRE_MAJOR_VERSION}-binaries/releases/download/jdk-${JRE_VERSION}+9_openj9-${OPENJ9_VERSION}/OpenJDK${JRE_MAJOR_VERSION}U-jre_x64_linux_openj9_${JRE_VERSION}_9_openj9-${OPENJ9_VERSION}.tar.gz


ARG MAVEN_MAJOR_VERSION=3
ARG MAVEN_MINOR_VERSION=8
ARG MAVEN_PATCH_VERSION=4
ENV MAVEN_VERSION=${MAVEN_MAJOR_VERSION}.${MAVEN_MINOR_VERSION}.${MAVEN_PATCH_VERSION}
ARG MAVEN_HOME=/opt/apache/maven${MAVEN_VERSION}
ENV MAVEN_BIN_NAME=maven-${MAVEN_VERSION}.tar.gz
ENV MAVEN_DOWNLOAD_URL=https://dlcdn.apache.org/maven/maven-${MAVEN_MAJOR_VERSION}/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz


# 明确指定工作目录，防止后面运行命令出现文件或者目录找不到的问题
WORKDIR /opt


RUN apk update
RUN apk add axel

# 安装AdoptOpenJDK，替代Oracle JDK
RUN axel --num-connections 16 --output jre${JRE_VERSION}.tar.gz --insecure ${JRE_DOWNLOAD_URL}
RUN tar -xzf jre${JRE_VERSION}.tar.gz
RUN mkdir -p /usr/lib/jvm/java-${JRE_MAJOR_VERSION}-adoptopenjdk-amd64
RUN mv jdk-${JRE_VERSION}+9-jre/* /usr/lib/jvm/java-${JRE_MAJOR_VERSION}-adoptopenjdk-amd64

# 安装Maven
RUN axel --num-connections 16 --output ${MAVEN_BIN_NAME} --insecure ${MAVEN_DOWNLOAD_URL}
RUN tar -xzf ${MAVEN_BIN_NAME}
RUN mkdir -p ${MAVEN_HOME}
RUN ls -al






# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


ARG MAVEN_VERSION
ARG MAVEN_HOME


# 复制文件
COPY --from=maven /usr/lib/jvm /usr/lib/jvm
COPY maven /bin


RUN set -ex \
    \
    \
    \
    && apk update \
    \
    # 安装FastGithub依赖库 \
    && apk --no-cache add libgcc libstdc++ gcompat icu \
    \
    # 安装Git工具
    && apk --no-cache add git openssh-client \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/git \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/git


# 配置环境变量
# 设置Java安装目录
ENV JAVA_HOME /usr/lib/jvm/java-11-adoptopenjdk-amd64
ENV JAVA_OPTS ""

ENV PATH=${MAVEN_HOME}/bin:$PATH
