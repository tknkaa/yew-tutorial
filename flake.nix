{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        gemini = pkgs.writeShellScriptBin "gemini" ''
          bunx @google/gemini-cli "$@"
        '';
      in
      {
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.bun
            gemini
          ]
          ;
          shellHook = ''
           source .env 
          '';
        };
      }
    );
}