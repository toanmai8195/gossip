# Sử dụng Debian Bookworm (Hỗ trợ tốt ARM64)
FROM debian:bookworm AS builder

# Cài đặt dependencies (thêm gcc, g++)
RUN apt-get update && apt-get install -y \
    curl \
    openjdk-17-jdk \
    unzip \
    zip \
    bash \
    git \
    gcc \
    g++ \
    && rm -rf /var/lib/apt/lists/*

# Cài đặt Bazelisk (phiên bản ổn định của Bazel)
RUN curl -fsSL https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-arm64 -o /usr/local/bin/bazel && \
    chmod +x /usr/local/bin/bazel

# Tạo user non-root để chạy Bazel
RUN useradd -m -s /bin/bash bazeluser

# Chuyển sang user bazeluser
USER bazeluser
WORKDIR /home/bazeluser/app

# Sao chép mã nguồn vào container
COPY --chown=bazeluser . .

# Thiết lập biến môi trường cho compiler
ENV CC=/usr/bin/gcc CXX=/usr/bin/g++

# Build ứng dụng với Bazel
RUN bazel build //:hello_world_deploy.jar   

# Chạy ứng dụng khi container khởi động
CMD ["java", "-jar", "bazel-bin/hello_world_deploy.jar"]