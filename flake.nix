{
  description = "Solutions for the Hackattic programming challenge site.";
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";

  outputs = { self, nixpkgs, gomod2nix }:
  let
    supportedSystems = [ "x86_64-linux" "aarch64-darwin" ];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system:
    import nixpkgs {
      inherit system;
      overlays = [ gomod2nix.overlays.default ];
    });
  in
  {
    devShells = forAllSystems (system:
    let
      pkgs = nixpkgsFor.${system};
    in
    {
      default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
          gotools
          go-tools
        ];
      };
    });
  };
}
