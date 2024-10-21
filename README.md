# Sept /ʂɛpt/
[![Go Report Card](https://goreportcard.com/badge/github.com/piheta/sept)](https://goreportcard.com/report/github.com/piheta/sept) 
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=piheta_sept&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=piheta_sept) 
[![Tests](https://img.shields.io/github/actions/workflow/status/piheta/sept/go.yml?logo=github&label=tests)](https://img.shields.io/github/actions/workflow/status/piheta/sept/go.yml?logo=github&label=tests) 
[![Status](https://badgen.net/badge/works/almost/red)](https://badgen.net/badge/works/almost/red) 
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Decentralized, Asynchronous P2P Chat

**Sept** is a video and chat application similar to Discord but operates without a central server. Clients connect directly to each other, resulting in faster connections than similar applications. Other limitations, such as file upload limits, do not exist.

Inspired by various projects and technologies, such as [VXLAN](https://www.rfc-editor.org/rfc/rfc7348) and the old [Skype P2P](https://arxiv.org/pdf/cs/0412017) infrastructure.

![Sept UI](./docs/images/sept.gif)

## Features
- **Private & Encrypted**: All messages are sent over an encrypted channel and stored locally on your device. Sept cannot see your messages.
- **No Upload Limits**: Share terabytes of files for free—no need to upload them.
- **Fast**: Faster than other chat and communication programs due to its P2P nature.
- **Better Video Quality**: Streams are sent directly to your peers without server throttling.

## Build
First, deploy a [Sept Login Server](https://github.com/piheta/sept-login-server):
```bash
git clone https://github.com/piheta/sept.git && cd sept
go install github.com/wailsapp/wails/v2/cmd/wails@latest
wails dev
```

## Dev Roadmap
- [x] 👍 Local per user db
- [x] 👍 P2P Message exchange
- [x] 👍 User search
- [ ] ⌛ Video
- [ ] ⌛ File uploads
- [ ] ⌛ DB encryption
- [ ] ⌛ VXLAN implementation

## License
**Sept** is released under the [GPL v3 License](LICENSE).

### Dependency Licenses
This project uses dependencies that are licensed under the following licenses:
- **MIT:** [pion/webrtc](https://github.com/pion/webrtc), [wails](https://github.com/wailsapp/wails)

## C++ POC
check out the legacy multicast [implementation of sept](https://github.com/piheta/sept/tree/legacy)
