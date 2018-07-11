FROM scratch

COPY bin/pslytics-api pslytics-api

CMD ["/pslytics-api"]