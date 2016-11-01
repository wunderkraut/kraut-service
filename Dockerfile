FROM golang
MAINTAINER james.nesbitt@wunder.io

# Create app user with app group. Specifically assign uid and gid 1000
RUN /usr/sbin/useradd --home /app --uid 1000 app
RUN mkdir /app && \
    mkdir /app/bin && \
    mkdir /app/build && \
    /bin/chown -R app:app /app

ADD ./ /app/build
ENV KRAUT_BUILD_PATH=/app/bin
RUN cd /app/build && \
    make all && \

ENTRYPOINT ["/app/bin/kraut-service"]
