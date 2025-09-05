FROM debian:13-slim AS mise

RUN apt-get update  \
  && apt-get -y --no-install-recommends install  \
  sudo curl git ca-certificates build-essential \
  && rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ENV MISE_DATA_DIR="/mise"
ENV MISE_CONFIG_DIR="/mise"
ENV MISE_CACHE_DIR="/mise/cache"
ENV MISE_INSTALL_PATH="/usr/local/bin/mise"
ENV MISE_EXPERIMENTAL="true"
ENV PATH="/mise/shims:$PATH"

RUN curl https://mise.run | sh

FROM mise AS build
WORKDIR /app

COPY mise.toml .
RUN mise trust && mise install

COPY . .
RUN cd frontend && bun install --frozen-lockfile

ENV NODE_ENV=production
RUN task build

FROM debian:13-slim AS production
ENV NODE_ENV=production
COPY --from=build /app/frontend/dist ./static
COPY --from=build /app/backend/bin/server .

RUN useradd -m appuser
USER appuser

CMD ["./server"]
