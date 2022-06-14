FROM alpine:latest

RUN mkdir "/usr/app"
COPY "./tg-bot" "/usr/app/bot"
RUN ["chmod", "+x", "./usr/app/bot"]

WORKDIR "/usr/app"
CMD ["./bot"]