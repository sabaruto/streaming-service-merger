{
  description = "Dev shells flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystemPassThrough (system: {

      devShells."${system}" = {
        default =
          let
            pkgs = import nixpkgs { inherit system; };
          in
          with pkgs;
          mkShell {
            packages = [
              zulu24
              jdt-language-server
            ];

            shellHook = ''
              export JDTLS_HOME=${jdt-language-server}
            '';
          };
      };
    });
}
