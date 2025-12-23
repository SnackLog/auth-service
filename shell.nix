{ pkgs ? import <nixpkgs> { } }:


pkgs.mkShell {
  name = "go-dev-shell";


  packages = with pkgs; [
    go
    gopls
    gotools
    delve
    go-licenses
  ];


  # Environment variables
  GOPRIVATE = "";
  GONOPROXY="github.com/SnackLog/*";

  shellHook = ''
    echo "Go development shell ready"
    go version
  '';
}

