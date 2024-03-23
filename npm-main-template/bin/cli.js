#!/usr/bin/env node

// Lookup table for all platforms and binary distribution packages
const BINARY_DISTRIBUTION_PACKAGES = require("./arch-packages.json");

// Windows binaries end with .exe so we need to special case them.
const binaryName =
  process.platform === "win32" ? "gen-elm-wrappers.exe" : "gen-elm-wrappers";

// Determine package name for this platform
const platformSpecificPackageName =
  BINARY_DISTRIBUTION_PACKAGES[`${process.platform}-${process.arch}`];

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`${platformSpecificPackageName}/bin/${binaryName}`);
  } catch (e) {
    return require("path").join(__dirname, "..", binaryName);
  }
}

require("child_process").execFileSync(getBinaryPath(), process.argv.slice(2), {
  stdio: "inherit",
});
