# sept/ʂɛpt/
[![Go Report Card](https://goreportcard.com/badge/github.com/piheta/sept)](https://goreportcard.com/report/github.com/piheta/sept) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=piheta_sept&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=piheta_sept) [![tests](https://img.shields.io/github/actions/workflow/status/piheta/sept/go.yml?logo=github&label=tests)](https://img.shields.io/github/actions/workflow/status/piheta/sept/go.yml?logo=github&label=tests) [![Status](https://badgen.net/badge/works/almost/red)](https://badgen.net/badge/works/almost/red) [![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) 
## decentralized, asynchronous p2p chat

sept is a video/chat application similar to discord with no central server. The clients connect directly to eachother making the connection faster than similar applications. Other limitations such as file upload limit are also non existent. 
sept is inspired by various projects, such as hamachi and the old skype p2p infrastructure.

![network system diagram](./docs/images/sept.png)

## features
- **private & encrypted:** all messages are sent over an encrypted channel and stored on your device. Sept can't see anything
- **no upload limits:** share terrabytes of files for free. No need to upload them
- **fast:** faster than other chat and communication programs because of its p2p nature
- **better video quality:** your streams are directly being sent to your peers, no server will throttle your performance

## build
first, deploy a [sept login server](https://github.com/piheta/sept-login-server)
 ```sh
  git clone https://github.com/piheta/sept.git
  wails dev
  ```

## dev roadmap
- [x] 👍 Local per user db
- [x] 👍 P2P Message exchange
- [ ] ⌛ Video
- [ ] ⌛ Backend-Frontend connection
- [ ] ⌛ User search
- [ ] ⌛ DB encryption
- [ ] ⌛ VXLAN implementation

## license
sept is released under the [GPL v3 License](LICENSE).

## c++ poc
check out the legacy multicast [implementation of sept](https://github.com/piheta/sept/tree/legacy)
