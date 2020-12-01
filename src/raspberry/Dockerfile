FROM gcc:9

WORKDIR /rasp
COPY . .

RUN ./scripts/install-deps.sh
CMD tail -f /dev/null