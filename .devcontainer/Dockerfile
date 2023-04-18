FROM golang:latest

# Configure to reduce warnings and limitations as instruction from official VSCode Remote-Containers.
# See https://code.visualstudio.com/docs/remote/containers-advanced#_reducing-dockerfile-build-warnings.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1

# Verify git, process tools, lsb-release (common in install instructions for CLIs) installed.
RUN apt-get -y install git iproute2 procps lsb-release

# Install Go tools.
RUN apt-get update \
    go install -v golang.org/x/tools/cmd/goimports@latest \
    go install -v golang.org/x/tools/gopls@latest \
    go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949 \
    go install -v github.com/go-delve/delve/cmd/dlv@latest

# Revert workaround at top layer.
ENV DEBIAN_FRONTEND=dialog

# Expose service ports.
EXPOSE 8000