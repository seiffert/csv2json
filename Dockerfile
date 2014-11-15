FROM busybox
MAINTAINER Paul Seiffert <paul.seiffert@gmail.com>

ADD .build/bin/csv2json /csv2json

ENTRYPOINT ["/csv2json"]
CMD ["--"]