# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Tvm < Formula
  desc "tvm - Terraform Version Manager"
  homepage "https://www.thorsten-hans.com"
  version "0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/ThorstenHans/tvm/releases/download/v0.0.1/tvm_darwin_arm64_0.0.1.tar.gz"
      sha256 "5b137b322326e6ecb77cde808cf970c2e46f1584e942db3fdb1811a668c3dbd8"

      def install
        bin.install "tvm"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ThorstenHans/tvm/releases/download/v0.0.1/tvm_darwin_amd64_0.0.1.tar.gz"
      sha256 "c8fa2c025daeaea8557e4dc7537cbfa44d8b357cd455a78f2a8310484d796273"

      def install
        bin.install "tvm"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/ThorstenHans/tvm/releases/download/v0.0.1/tvm_linux_arm64_0.0.1.tar.gz"
      sha256 "3463f6cbb019c15b95b8c922db32a4c5c22af399b72d4972f1f86e6a176f11bc"

      def install
        bin.install "tvm"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ThorstenHans/tvm/releases/download/v0.0.1/tvm_linux_amd64_0.0.1.tar.gz"
      sha256 "1ea37c56ff8b41465bbc37b9083b44242d6ffabdf9808c34df14a0bfc52c78fb"

      def install
        bin.install "tvm"
      end
    end
  end

  conflicts_with "terraform"
end
