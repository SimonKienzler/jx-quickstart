FROM scratch
EXPOSE 8080
ENTRYPOINT ["/jx-quickstart"]
COPY ./bin/ /