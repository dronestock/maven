FROM storezhang/alpine AS maven


ENV JRE_VERSION 11.0.11
ENV JRE_MAJOR_VERSION 11
ENV OPENJ9_VERSION 0.26.0
ENV MAVEN_VERSIONI 3.8.4


RUN apk update
RUN apk add axel

# 安装AdoptOpenJDK，替代Oracle JDK
RUN axel --num-connections 6 --output jre${JRE_VERSION}.tar.gz --insecure "https://download.fastgit.org/AdoptOpenJDK/openjdk${JRE_MAJOR_VERSION}-binaries/releases/download/jdk-${JRE_VERSION}+9_openj9-${OPENJ9_VERSION}/OpenJDK${JRE_MAJOR_VERSION}U-jre_x64_linux_openj9_${JRE_VERSION}_9_openj9-${OPENJ9_VERSION}.tar.gz"
RUN tar -xzf jre${JRE_VERSION}.tar.gz
RUN mkdir -p /usr/lib/jvm/java-${JRE_MAJOR_VERSION}-adoptopenjdk-amd64
RUN mv jdk-${JRE_VERSION}+9-jre/* /usr/lib/jvm/java-${JRE_MAJOR_VERSION}-adoptopenjdk-amd64
# 安装Maven






# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


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
