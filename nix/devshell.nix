{ pkgs }:
pkgs.mkShell {
  # Add build dependencies
  packages = with pkgs; [
    elixir
    erlang
  ];

  # Add environment variables
  env = { };

  # Load custom bash code
  shellHook = ''

  '';
}
