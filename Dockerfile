FROM alpine:latest as certs
RUN apk -U add ca-certificates

FROM scratch
ENV CERTS=/etc/ssl/certs/ca-certificates.crt
COPY --from=certs $CERTS $CERTS
ADD jams-manager /
EXPOSE 8080
CMD ["/jams-manager"]