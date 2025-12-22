{ pkgs ? import <nixpkgs> { } }:


pkgs.mkShell {
  name = "go-dev-shell";


  packages = with pkgs; [
    go
    gopls
    gotools
    delve
  ];


  # Environment variables
  GOPRIVATE = "";


  shellHook = ''
    echo "Go development shell ready"
    go version
  '';
}

