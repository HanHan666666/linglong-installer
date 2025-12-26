# Linglong Installer Distro Scripts

This installer now supports distro-specific install scripts that are selected
at runtime based on `/etc/os-release` (ID and VERSION_ID). Each supported
release has its own `ID_VERSION.sh` script so adding a new distro is just
adding one file.

## How it works

- `core.DetectEnv` reads `/etc/os-release` and fills `env.distro` and
  `env.distroVersion`.
- The custom task `distroScript` selects the script, parses `# META:` lines,
  and writes values into the install context.
- The confirm screen renders those values for the user, then the progress
  step runs the selected script with root privileges.

## Scripts layout

- `scripts/common.sh`: shared helpers (root check, repo helpers, logging).
- `scripts/distros/*.sh`: per-distro scripts named `ID_VERSION.sh`.
- `install-linyaps.sh`: standalone dispatcher for manual use (optional).

Resolution order for scripts:

1. `scripts/distros/<ID>_<VERSION>.sh` under current working directory
2. `scripts/distros/<ID>_<VERSION>.sh` next to the executable
3. Embedded scripts (compiled into the binary)

## META header format

Each distro script should start with header lines like this:

```
#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (Debian 13)
# META: repo_url=https://ci.deepin.com/repo/obs/linglong:/CI:/release/Debian_13/
# META: command=apt update
# META: command=apt install -y linglong-bin linglong-installer
# META: next_steps=Add Linglong repo and install required packages.
```

Supported META keys:

- `repo_name`: name shown in the confirmation screen
- `repo_url`: repo URL shown in the confirmation screen
- `command`: repeatable; each entry is listed in the confirmation screen
- `next_steps`: a short sentence shown in the confirmation screen
- `distro_name` (optional): overrides the detected display name

The META data is used only for display; the actual commands are executed by
script logic, not by these strings.

## Context keys written by distroScript

- `distro.id` / `distro.version` / `distro.name`
- `distro.supported` / `distro.error`
- `distro.script` / `distro.script_name`
- `distro.repo_name` / `distro.repo_url`
- `distro.next_steps` / `distro.commands`

These are referenced by the confirmation screen in `installer.yaml`.

## Add a new distro

1. Create a new script in `scripts/distros/` named `ID_VERSION.sh`.
2. Include `# META:` headers with repo and command info.
3. Implement the install logic using helpers from `scripts/common.sh`.
4. Test by running the installer and verifying the confirmation screen.

Example filename for `/etc/os-release` with `ID=debian` and `VERSION_ID=13`:

- `scripts/distros/debian_13.sh`

## Run the installer

```
go run ./ --config installer.yaml
```
