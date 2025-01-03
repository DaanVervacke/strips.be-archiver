<div align="center">
    <img height="95" src="./.github/strips-be-archiver-logo.png" alt="strips.be-archiver-logo" style="margin-bottom: -10px;">
    <h1 style="margin-top: 0;">strips.be-archiver</h1>
    <sup><em>command-line interface (CLI) tool designed for interacting with the strips.be app</em></sup>
</div>

## Features

- 🚀 Written in Golang
- 🖥️ Cross-platform
- 🔑 Log-in with your personal account
- 🔍 Search for albums and series along with their corresponding UUIDs
- 🗃️ Archive your favorite comics as `.cbz` files
- ℹ️ Embed [ComicInfo metadata](https://anansi-project.github.io/docs/comicinfo/intro)

## Build

```bash
make build
```

## Usage

```bash
strips.be-archiver --help
```

## Config

> [!IMPORTANT]
> Users must independently seek out the internal endpoints and secrets required by the app.
> An empty example configuration file (`example.config.yaml`) is provided in the root of this project.
> To function properly, the tool requires a complete and accurate configuration file.
> Use the `--config` flag to pass the configuration to the tool, as shown below.

```bash
strips.be-archiver --config config.yaml
```

## Content

As strips.be is a paid service requiring a monthly or yearly subscription, only users with a valid subscription can
archive all series and albums.
Users without a subscription are limited to archiving comics from the "free to read" section.

## Contributing

Pull requests are welcome. For major changes, please open an issue first.

## Legal Notice

By using `strips.be-archiver`, you agree to comply with
the [strips.be terms of service](https://strips.be/algemene-gebruiks-en-verkoopvoorwaarden-strips-be/).

This tool does not bypass any digital rights management (DRM) protections and is specifically designed to avoid leaking
any internal endpoints or secrets used by the strips.be app. It is intended solely for archiving publicly available
content that users have permission to download and preserve.

`strips.be-archiver` is a passion project created for the purpose of preserving comics. The goal is
to provide a way for individuals to archive their own personal collections, especially in cases where comics may no
longer be publicly available or accessible. The tool does not promote, encourage, or facilitate the distribution of
copyrighted material.

We are not responsible for any legal issues, copyright violations, blacklisting or other consequences resulting from the
use of this
tool. The use of this tool is entirely at the user's own risk.

Users are solely responsible for ensuring that their use of the tool complies with all applicable laws, terms of
service, and copyright regulations. The repository owner and contributors cannot be held accountable for any actions or
legal consequences resulting from the use of the tool, including but not limited to the unauthorized distribution,
downloading, or archiving of content that is not publicly available or that violates copyright law.

This project is provided "as-is" and is intended for personal use only. We strongly encourage users to respect the
intellectual property rights of creators and ensure they have the proper rights or permissions for any content they
archive.

If you are a copyright holder and believe that this tool is being used in violation of your rights, please contact the
repository owner. We are committed to resolving any concerns, including the potential shutdown of the repository if
necessary.