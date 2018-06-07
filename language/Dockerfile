FROM            openjdk
MAINTAINER      BDNA Engineering <engineering-docker@bdna.com>

WORKDIR         /usr/local/lib
RUN             curl -O https://www.antlr.org/download/antlr-4.7.1-complete.jar

WORKDIR         /usr/src/
ENTRYPOINT      ["java",  "-jar", "/usr/local/lib/antlr-4.7.1-complete.jar"]
