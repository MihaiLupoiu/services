FROM alpine

COPY users_unix /users
ENTRYPOINT [ "/users" ]
CMD ["-config", "/config.json"]