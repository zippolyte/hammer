# This file was generated by GoReleaser. DO NOT EDIT.
class Hammer < Formula
  desc ""
  homepage ""
  version "0.8.1"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal/hammer/releases/download/v0.8.1/hammer_darwin_amd64.tar.gz"
    sha256 "6b418f15787842cc9744fabe7fdb07b94e008c822d2a1ea9c94c3b0c918b7824"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/pivotal/hammer/releases/download/v0.8.1/hammer_linux_amd64.tar.gz"
      sha256 "15a165b04b9135fa3240813f06232a79b9af4bb618272ddda853756074d1626c"
    end
  end

  def install
    bin.install "hammer"
  end

  test do
    system "#{bin}/hammer version"
  end
end
