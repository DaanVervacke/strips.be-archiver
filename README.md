# strips.be-archiver

strips.be-archiver is a Golang library/CLI for interacting with the strips.be app. It allows users to properly archive
comics.

Comics are saved as .cbz files and include a ComicInfo.xml file for metadata, which follows
the [Anasi Project Standards](https://anansi-project.github.io/).

## Build

```bash
make build
```

## Usage

```bash
strips.be-archiver --help
```

## Example config

```yaml
config:
  api:
    baseUrl:
    albumPath:
    seriesPath:
    accountPath:
    profilePath:
    tradePath:
    refreshPath:

    basicHeaders:
      Accept-Encoding:
      AppVersion:
      Host:
      User-Agent:

    tradeHeaders:
      x-device-os:
      x-device-os-version:
      x-device-type:

    playbookHeaders:
      Accept-Encoding:
      AppVersion:
      Host:
      User-Agent:

  auth:
    baseUrl:
    otpPath:
    otpRedirectTo:
    verifyPath:

    headers:
      apikey:
      Accept-Encoding:
      Authorization:
      Content-Type:
      Host:
      User-Agent:
      x-client-info:

    account:
      accessToken:
      refreshToken:
      deviceId:
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first.

## Legal Notice

By using strips.be-archiver, you agree to comply with
the [strips.be terms of service](https://strips.be/algemene-gebruiks-en-verkoopvoorwaarden-strips-be/).

This tool does not bypass any digital rights management (DRM) protections and is specifically designed to avoid leaking
any internal endpoints or secrets used by the strips.be app. It is intended solely for archiving publicly available
content that users have permission to download and preserve.

strips.be-archiver is a passion project created for the purpose of preserving comics. The goal is
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