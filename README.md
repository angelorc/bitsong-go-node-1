<p align="center" background="black"><img src="bitsong-logo.png" width="398"></p>

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/BitSongOfficial/bitsong-go-node/blob/master/LICENSE)

BitSong is a music blockchain network that lets aritsts and fans issue and manage their own coins and trade them at a fair market price with absolute and instant liquidity.

_NOTE: This is alpha software. Please contact us if you intend to run it in production._

## Installation

1. Download BitSong

    Get [latest binary build](https://github.com/BitSongOfficial/bitsong-go-node/releases) suitable for your architecture and unpack it to desired folder.

2. Run BitSong Node

```bash
./bitsongd
```

3. Use GUI

    Open http://localhost:3000/ in local browser to see node’s GUI.

## Resources
- [Official Website](https://bitsong.io)

### Community
- [Telegram Channel (English)](https://t.me/BitSongOfficial)
- [Telegram Group (English)](https://t.me/bitsong_ico)
- [Facebook](https://www.facebook.com/BitSongOfficial)
- [Twitter](https://twitter.com/BitSongOfficial)
- [Medium](https://medium.com/@BitSongOfficial)
- [Reddit](https://www.reddit.com/r/bitsong/)
- [BitcoinTalk ANN](https://bitcointalk.org/index.php?topic=2850943)
- [Linkedin](https://www.linkedin.com/company/bitsong)

## License

MIT License

## Versioning

### SemVer

BitSong uses [SemVer](http://semver.org/) to determine when and how the version changes.
According to SemVer, anything in the public API can change at any time before version 1.0.0

To provide some stability to BitSong users in these 0.X.X days, the MINOR version is used
to signal breaking changes across a subset of the total public API. This subset includes all
interfaces exposed to other processes, but does not include the in-process Go APIs.

### Upgrades

In an effort to avoid accumulating technical debt prior to 1.0.0,
we do not guarantee that breaking changes (ie. bumps in the MINOR version)
will work with existing blockchain. In these cases you will
have to start a new blockchain, or write something custom to get the old
data into the new chain.

However, any bump in the PATCH version should be compatible with existing histories
(if not please open an [issue](https://github.com/BitSongOfficial/bitsong-go-node/issues)).