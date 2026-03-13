{ pkgs }:
pkgs.mkShell {
  # Add build dependencies
  packages = with pkgs; [
    (rust-bin.stable.latest.default.override {
      targets = [ "wasm32-unknown-unknown" ];
    })
    trunk
  ];

  # Add environment variables
  env = { };

  # Load custom bash code
  shellHook = ''

  '';
}
