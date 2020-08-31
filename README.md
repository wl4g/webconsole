# XCloud WebConsole
A lightweight enterprise Web Console terminal, supporting SSH2/(RDP/RFB/Telnet) protocol, etc

### Features
- It is designed as a native JavaScript class library, which can be easily integrated with React/Vue/AngularJS and other frameworks
- It can run on Android / IOS and any other terminal that can render HTML. It can almost completely replace the shell client based on the installation program
- Fully support lrzsz command set (implemented based on zmodem)
- Enhanced support for mobile terminal copy, paste, fast forward, backward and other key combination command, user-friendly operation habits

### TODO
- [√] Completely unify the daily output of each component, such as `gin` framework.
- Enhance the administrator functions of webconsole service, such as its own health/metrics/indicator(CPU/Mem/Network/Connections...) And more detailed indicators
- In order to realize the remote image UI control protocol compatible with windows RDP(Remote Desktop Protocol) based on Web
- In order to realize the remote image UI control protocol compatible with windows RFB(Remote FrameBuffer) based on Web
- In order to realize remote command control protocol compatible with Telnet protocol based on Web